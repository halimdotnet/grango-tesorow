variable "environment" {
  description = "Environment name"
  type        = string
}

variable "project_name" {
  description = "Project name"
  type        = string
  default     = "tesorow"
}

variable "vpc_id" {
  description = "VPC ID"
  type        = string
}

variable "database_subnet_ids" {
  description = "Database subnet IDs"
  type        = list(string)
}

variable "db_security_group_id" {
  description = "Database security group ID"
  type        = string
}

variable "allocated_storage" {
  description = "Allocated storage in GB"
  type        = number
  default     = 20
}

variable "instance_class" {
  description = "RDS instance class"
  type        = string
  default     = "db.t3.micro"
}

variable "engine_version" {
  description = "PostgreSQL engine version"
  type        = string
  default     = "16"
}

variable "backup_retention_period" {
  description = "Backup retention period in days"
  type        = number
  default     = 7
}

variable "db_subnet_group_name" {
  description = "Database subnet group name (from VPC module)"
  type        = string
}

# Random password for master database user
resource "random_password" "db_master_password" {
  length  = 32
  special = true
  # Exclude characters that might cause issues in connection strings
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

# Store database credentials in Secrets Manager
resource "aws_secretsmanager_secret" "db_credentials" {
  name        = "${var.project_name}/${var.environment}/db-credentials"
  description = "Database credentials for ${var.project_name} ${var.environment}"

  tags = {
    Name        = "${var.project_name}-db-credentials-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

resource "aws_secretsmanager_secret_version" "db_credentials" {
  secret_id = aws_secretsmanager_secret.db_credentials.id
  secret_string = jsonencode({
    username = "tesorow_admin"
    password = random_password.db_master_password.result
    engine   = "postgres"
    host     = aws_db_instance.main.address
    port     = aws_db_instance.main.port
    dbname   = aws_db_instance.main.db_name
  })
}

# DB Parameter Group for PostgreSQL
resource "aws_db_parameter_group" "main" {
  name   = "${var.project_name}-postgres-${var.environment}"
  family = "postgres16"

  # Enable query logging for audit
  parameter {
    name  = "log_statement"
    value = "all"
  }

  parameter {
    name  = "log_min_duration_statement"
    value = "1000" # Log queries taking more than 1 second
  }

  # SSL enforcement
  parameter {
    name  = "rds.force_ssl"
    value = "1"
  }

  tags = {
    Name        = "${var.project_name}-postgres-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

# RDS PostgreSQL Instance
resource "aws_db_instance" "main" {
  identifier = "${var.project_name}-db-${var.environment}"

  # Engine configuration
  engine               = "postgres"
  engine_version       = var.engine_version
  instance_class       = var.instance_class
  allocated_storage    = var.allocated_storage
  storage_type         = "gp3"
  storage_encrypted    = true

  # Database configuration
  db_name  = "tesorow_${var.environment}"
  username = "tesorow_admin"
  password = random_password.db_master_password.result
  port     = 5432

  # Network configuration
  db_subnet_group_name   = var.db_subnet_group_name
  vpc_security_group_ids = [var.db_security_group_id]
  publicly_accessible    = false

  # Parameter group
  parameter_group_name = aws_db_parameter_group.main.name

  # Backup configuration
  backup_retention_period = var.backup_retention_period
  backup_window          = "03:00-04:00" # UTC
  maintenance_window     = "mon:04:00-mon:05:00" # UTC

  # Enable automated backups
  skip_final_snapshot       = var.environment == "staging" ? true : false
  final_snapshot_identifier = var.environment == "staging" ? null : "${var.project_name}-db-final-snapshot-${var.environment}-${formatdate("YYYY-MM-DD-hhmm", timestamp())}"

  # Copy tags to snapshots
  copy_tags_to_snapshot = true

  # Monitoring - Basic (free tier)
  enabled_cloudwatch_logs_exports = ["postgresql", "upgrade"]
  monitoring_interval             = 0 # 0 = disabled, 60 = enabled (cost extra)

  # Performance Insights - Cost extra, comment out for free tier
  # performance_insights_enabled    = true
  # performance_insights_retention_period = 7

  # Multi-AZ - Cost ~2x, comment out for free tier
  # multi_az = var.environment == "production" ? true : false

  # Deletion protection - Enable for production
  deletion_protection = var.environment == "production" ? true : false

  # Auto minor version upgrade
  auto_minor_version_upgrade = true

  tags = {
    Name        = "${var.project_name}-db-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }

  # Prevent accidental deletion during terraform destroy
  lifecycle {
    prevent_destroy = false # Set to true for production
  }
}

# CloudWatch Alarm for high CPU
resource "aws_cloudwatch_metric_alarm" "database_cpu" {
  alarm_name          = "${var.project_name}-db-cpu-${var.environment}"
  alarm_description   = "Database CPU utilization is too high"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/RDS"
  period              = "300"
  statistic           = "Average"
  threshold           = "80"

  dimensions = {
    DBInstanceIdentifier = aws_db_instance.main.id
  }

  tags = {
    Name        = "${var.project_name}-db-cpu-alarm-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

# CloudWatch Alarm for low free storage
resource "aws_cloudwatch_metric_alarm" "database_storage" {
  alarm_name          = "${var.project_name}-db-storage-${var.environment}"
  alarm_description   = "Database free storage is too low"
  comparison_operator = "LessThanThreshold"
  evaluation_periods  = "1"
  metric_name         = "FreeStorageSpace"
  namespace           = "AWS/RDS"
  period              = "300"
  statistic           = "Average"
  threshold           = "2000000000" # 2GB in bytes

  dimensions = {
    DBInstanceIdentifier = aws_db_instance.main.id
  }

  tags = {
    Name        = "${var.project_name}-db-storage-alarm-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}