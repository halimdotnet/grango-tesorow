variable "environment" {
  description = "Environment name"
  type        = string
}

variable "project_name" {
  description = "Project name"
  type        = string
  default     = "tesorow"
}

# ECS Task Execution Role
# Used by ECS agent to pull images and write logs
resource "aws_iam_role" "ecs_task_execution" {
  name = "${var.project_name}-ecs-task-execution-${var.environment}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })

  tags = {
    Name        = "${var.project_name}-ecs-task-execution-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

# Attach AWS managed policy for ECS task execution
resource "aws_iam_role_policy_attachment" "ecs_task_execution_policy" {
  role       = aws_iam_role.ecs_task_execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# Additional policy for Secrets Manager access (for database credentials)
resource "aws_iam_role_policy" "ecs_task_execution_secrets" {
  name = "${var.project_name}-ecs-secrets-${var.environment}"
  role = aws_iam_role.ecs_task_execution.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "secretsmanager:GetSecretValue",
          "secretsmanager:DescribeSecret"
        ]
        Resource = "arn:aws:secretsmanager:*:*:secret:${var.project_name}/${var.environment}/*"
      },
      {
        Effect = "Allow"
        Action = [
          "kms:Decrypt",
          "kms:DescribeKey"
        ]
        Resource = "*"
      }
    ]
  })
}

# ECS Task Role
# Used by your application to access AWS services
resource "aws_iam_role" "ecs_task" {
  name = "${var.project_name}-ecs-task-${var.environment}"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })

  tags = {
    Name        = "${var.project_name}-ecs-task-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

# Policy for application to access S3 (document storage)
resource "aws_iam_role_policy" "ecs_task_s3" {
  name = "${var.project_name}-ecs-s3-${var.environment}"
  role = aws_iam_role.ecs_task.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:GetObject",
          "s3:PutObject",
          "s3:DeleteObject",
          "s3:ListBucket"
        ]
        Resource = [
          "arn:aws:s3:::${var.project_name}-documents-${var.environment}",
          "arn:aws:s3:::${var.project_name}-documents-${var.environment}/*"
        ]
      }
    ]
  })
}

# Policy for application to access Secrets Manager (runtime secrets)
resource "aws_iam_role_policy" "ecs_task_secrets" {
  name = "${var.project_name}-ecs-runtime-secrets-${var.environment}"
  role = aws_iam_role.ecs_task.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "secretsmanager:GetSecretValue",
          "secretsmanager:DescribeSecret"
        ]
        Resource = "arn:aws:secretsmanager:*:*:secret:${var.project_name}/${var.environment}/*"
      }
    ]
  })
}

# Policy for application to write logs to CloudWatch
resource "aws_iam_role_policy" "ecs_task_logs" {
  name = "${var.project_name}-ecs-logs-${var.environment}"
  role = aws_iam_role.ecs_task.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents",
          "logs:DescribeLogStreams"
        ]
        Resource = "arn:aws:logs:*:*:log-group:/ecs/${var.project_name}/${var.environment}/*"
      }
    ]
  })
}

# Policy for application to access SES (for email notifications) - Optional
# Uncomment if application needs to send email
# resource "aws_iam_role_policy" "ecs_task_ses" {
#   name = "${var.project_name}-ecs-ses-${var.environment}"
#   role = aws_iam_role.ecs_task.id
#
#   policy = jsonencode({
#     Version = "2012-10-17"
#     Statement = [
#       {
#         Effect = "Allow"
#         Action = [
#           "ses:SendEmail",
#           "ses:SendRawEmail"
#         ]
#         Resource = "*"
#       }
#     ]
#   })
# }

# Policy for application to access SNS (for push notifications) - Optional
# Uncomment if application needs to publish to SNS
# resource "aws_iam_role_policy" "ecs_task_sns" {
#   name = "${var.project_name}-ecs-sns-${var.environment}"
#   role = aws_iam_role.ecs_task.id
#
#   policy = jsonencode({
#     Version = "2012-10-17"
#     Statement = [
#       {
#         Effect = "Allow"
#         Action = [
#           "sns:Publish"
#         ]
#         Resource = "arn:aws:sns:*:*:${var.project_name}-${var.environment}-*"
#       }
#     ]
#   })
# }

# RDS Enhanced Monitoring Role - Optional for production
# Uncomment to enable detailed RDS monitoring
# resource "aws_iam_role" "rds_monitoring" {
#   name = "${var.project_name}-rds-monitoring-${var.environment}"
#
#   assume_role_policy = jsonencode({
#     Version = "2012-10-17"
#     Statement = [
#       {
#         Effect = "Allow"
#         Principal = {
#           Service = "monitoring.rds.amazonaws.com"
#         }
#         Action = "sts:AssumeRole"
#       }
#     ]
#   })
#
#   tags = {
#     Name        = "${var.project_name}-rds-monitoring-${var.environment}"
#     Environment = var.environment
#     Project     = var.project_name
#   }
# }
#
# resource "aws_iam_role_policy_attachment" "rds_monitoring_policy" {
#   role       = aws_iam_role.rds_monitoring.name
#   policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonRDSEnhancedMonitoringRole"
# }