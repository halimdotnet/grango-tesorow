output "vpc_id" {
  description = "VPC ID"
  value       = module.vpc.vpc_id
}

output "vpc_cidr" {
  description = "VPC CIDR"
  value       = module.vpc.vpc_cidr
}

output "public_subnet_ids" {
  description = "Public subnet IDs"
  value       = module.vpc.public_subnet_ids
}

output "private_subnet_ids" {
  description = "Private subnet IDs"
  value       = module.vpc.private_subnet_ids
}

output "database_subnet_ids" {
  description = "Database subnet IDs"
  value       = module.vpc.database_subnet_ids
}

output "alb_security_group_id" {
  description = "ALB security group ID"
  value       = module.security_groups.alb_security_group_id
}

output "app_security_group_id" {
  description = "App security group ID"
  value       = module.security_groups.app_security_group_id
}

output "db_security_group_id" {
  description = "Database security group ID"
  value       = module.security_groups.db_security_group_id
}

output "ecs_task_execution_role_arn" {
  description = "ECS task execution role ARN"
  value       = module.iam.ecs_task_execution_role_arn
}

output "ecs_task_role_arn" {
  description = "ECS task role ARN"
  value       = module.iam.ecs_task_role_arn
}

output "db_endpoint" {
  description = "RDS instance endpoint"
  value       = module.rds.db_endpoint
}

output "db_address" {
  description = "RDS instance address"
  value       = module.rds.db_address
}

output "db_name" {
  description = "Database name"
  value       = module.rds.db_name
}

output "db_secret_name" {
  description = "Database credentials secret name"
  value       = module.rds.db_secret_name
}

output "documents_bucket_name" {
  description = "Documents S3 bucket name"
  value       = module.s3.documents_bucket_name
}

output "documents_bucket_arn" {
  description = "Documents S3 bucket ARN"
  value       = module.s3.documents_bucket_arn
}

output "logs_bucket_name" {
  description = "Logs S3 bucket name"
  value       = module.s3.logs_bucket_name
}