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

variable "public_subnet_ids" {
  description = "Public subnet IDs for ECS tasks (no ALB)"
  type        = list(string)
}

variable "app_security_group_id" {
  description = "Application security group ID"
  type        = string
}

variable "db_security_group_id" {
  description = "Database security group ID"
  type        = string
}

variable "ecr_repository_url" {
  description = "ECR repository URL"
  type        = string
}

variable "task_execution_role_arn" {
  description = "ECS task execution role ARN"
  type        = string
}

variable "task_role_arn" {
  description = "ECS task role ARN"
  type        = string
}

variable "db_secret_arn" {
  description = "Database credentials secret ARN"
  type        = string
}

variable "cpu" {
  description = "Fargate task CPU units"
  type        = string
  default     = "256" # 0.25 vCPU
}

variable "memory" {
  description = "Fargate task memory in MB"
  type        = string
  default     = "512" # 512 MB
}

variable "desired_count" {
  description = "Desired number of tasks"
  type        = number
  default     = 1
}

# ECS Cluster
resource "aws_ecs_cluster" "main" {
  name = "${var.project_name}-cluster-${var.environment}"

  setting {
    name  = "containerInsights"
    value = "enabled"
  }

  tags = {
    Name        = "${var.project_name}-cluster-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

# CloudWatch Log Group
resource "aws_cloudwatch_log_group" "app" {
  name              = "/ecs/${var.project_name}/${var.environment}"
  retention_in_days = 7

  tags = {
    Name        = "${var.project_name}-logs-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

# ECS Task Definition
resource "aws_ecs_task_definition" "app" {
  family                   = "${var.project_name}-task-${var.environment}"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.cpu
  memory                   = var.memory
  execution_role_arn       = var.task_execution_role_arn
  task_role_arn            = var.task_role_arn

  container_definitions = jsonencode([
    {
      name      = "app"
      image     = "${var.ecr_repository_url}:latest"
      essential = true

      portMappings = [
        {
          containerPort = 8081
          hostPort      = 8081
          protocol      = "tcp"
        }
      ]

      environment = [
        {
          name  = "APP_ENV"
          value = var.environment
        }
      ]

      secrets = [
        {
          name      = "DB_PASSWORD"
          valueFrom = "${var.db_secret_arn}:password::"
        }
      ]

      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-group"         = aws_cloudwatch_log_group.app.name
          "awslogs-region"        = data.aws_region.current.name
          "awslogs-stream-prefix" = "ecs"
        }
      }

      healthCheck = {
        command     = ["CMD-SHELL", "curl -f http://localhost:8081/health || exit 1"]
        interval    = 30
        timeout     = 5
        retries     = 3
        startPeriod = 60
      }
    }
  ])

  tags = {
    Name        = "${var.project_name}-task-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

# Security Group for ECS Tasks (allow internet access on port 8081)
resource "aws_security_group" "ecs_task" {
  name        = "${var.project_name}-ecs-task-sg-${var.environment}"
  description = "Security group for ECS tasks - allow public access"
  vpc_id      = var.vpc_id

  ingress {
    description = "Allow HTTP from internet"
    from_port   = 8081
    to_port     = 8081
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description = "Allow all outbound"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name        = "${var.project_name}-ecs-task-sg-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

# Security Group Rule - Allow ECS tasks to access database
resource "aws_security_group_rule" "ecs_to_db" {
  type                     = "ingress"
  from_port                = 5432
  to_port                  = 5432
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.ecs_task.id
  security_group_id        = var.db_security_group_id
  description              = "Allow ECS tasks to access database"
}

# ECS Service (without ALB)
resource "aws_ecs_service" "app" {
  name            = "${var.project_name}-service-${var.environment}"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.app.arn
  desired_count   = var.desired_count
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = var.public_subnet_ids
    security_groups  = [aws_security_group.ecs_task.id]
    assign_public_ip = true # Required for tasks in public subnet without NAT
  }

  # Health check grace period
  health_check_grace_period_seconds = 60

  tags = {
    Name        = "${var.project_name}-service-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }

  # Ignore changes to desired_count (for auto-scaling later)
  lifecycle {
    ignore_changes = [desired_count]
  }
}

# CloudWatch Alarm - No healthy tasks
resource "aws_cloudwatch_metric_alarm" "ecs_no_tasks" {
  alarm_name          = "${var.project_name}-ecs-no-tasks-${var.environment}"
  alarm_description   = "Alert when no ECS tasks are running"
  comparison_operator = "LessThanThreshold"
  evaluation_periods  = "1"
  metric_name         = "RunningTaskCount"
  namespace           = "ECS/ContainerInsights"
  period              = "60"
  statistic           = "Average"
  threshold           = "1"

  dimensions = {
    ServiceName = aws_ecs_service.app.name
    ClusterName = aws_ecs_cluster.main.name
  }

  tags = {
    Name        = "${var.project_name}-ecs-no-tasks-alarm-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

# CloudWatch Alarm - High CPU
resource "aws_cloudwatch_metric_alarm" "ecs_cpu_high" {
  alarm_name          = "${var.project_name}-ecs-cpu-high-${var.environment}"
  alarm_description   = "Alert when ECS CPU is too high"
  comparison_operator = "GreaterThanThreshold"
  evaluation_periods  = "2"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/ECS"
  period              = "300"
  statistic           = "Average"
  threshold           = "80"

  dimensions = {
    ServiceName = aws_ecs_service.app.name
    ClusterName = aws_ecs_cluster.main.name
  }

  tags = {
    Name        = "${var.project_name}-ecs-cpu-alarm-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

# Data source for current region
data "aws_region" "current" {}