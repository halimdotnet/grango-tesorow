output "vpc_id" {
  value = module.vpc.vpc_id
}

output "public_subnet_ids" {
  value = module.vpc.public_subnet_ids
}

output "private_subnet_ids" {
  value = module.vpc.private_subnet_ids
}

output "alb_sg_id" {
  value = module.security_groups.alb_sg_id
}

output "app_sg_id" {
  value = module.security_groups.app_sg_id
}

output "rds_sg_id" {
  value = module.security_groups.rds_sg_id
}

output "rds_endpoint" {
  value     = module.rds.db_endpoint
  sensitive = true
}

output "rds_database_name" {
  value = module.rds.db_name
}

output "rds_master_username" {
  value     = module.rds.db_username
  sensitive = true
}

# output "app_runner_service_url" {
#   value = module.app_runner.service_url
# }
#
# output "app_runner_service_arn" {
#   value = module.app_runner.service_arn
# }
#
# output "github_connection_arn" {
#   value = module.app_runner.connection_arn
# }

# NOTE: EC2 outputs are commented out due to module being disabled.

# output "ec2_instance_id" {
#   value = module.ec2.instance_id
# }
#
# output "ec2_private_ip" {
#   value = module.ec2.private_ip
# }
#
# output "ec2_availability_zone" {
#   value = module.ec2.availability_zone
# }

# NOTE: ALB outputs are commented out due to module being disabled.

# output "alb_dns_name" {
#   value = module.alb.alb_dns_name
# }
#
# output "alb_arn" {
#   value = module.alb.alb_arn
# }
#
# output "target_group_arn" {
#   value = module.alb.target_group_arn
# }