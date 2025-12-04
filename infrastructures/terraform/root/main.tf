module "vpc" {
  source = "../modules/vpc"

  project_name          = var.project_name
  vpc_cidr              = var.vpc_cidr
  availability_zones    = var.availability_zones
  public_subnet_cidrs   = var.public_subnet_cidrs
  private_subnet_cidrs  = var.private_subnet_cidrs
}

module "security_groups" {
  source = "../modules/security_groups"

  project_name = var.project_name
  vpc_id       = module.vpc.vpc_id
}

module "rds" {
  source = "../modules/rds"

  project_name              = var.project_name
  db_subnet_group_name      = "${var.project_name}-db-subnet-group"
  private_subnet_ids        = module.vpc.private_subnet_ids
  rds_security_group_id     = module.security_groups.rds_sg_id

  rds_identifier            = var.rds_identifier
  rds_engine_version        = var.rds_engine_version
  rds_instance_class        = var.rds_instance_class
  rds_allocated_storage     = var.rds_allocated_storage
  rds_max_allocated_storage = var.rds_max_allocated_storage
  rds_database_name         = var.rds_database_name
  rds_master_username       = var.rds_master_username
  rds_master_password       = var.rds_master_password
  rds_backup_retention_days = var.rds_backup_retention_days
  rds_multi_az              = var.rds_multi_az
}

# NOTE: ALB/NLB module is commented out due to AWS account limitations.

# module "ec2" {
#   source = "../modules/ec2"
#
#   project_name           = var.project_name
#   instance_type          = var.ec2_instance_type
#   ami_id                 = "ami-0c55b159cbfafe1f0"
#   private_subnet_id      = module.vpc.private_subnet_ids[0]
#   app_security_group_id  = module.security_groups.app_sg_id
#   instance_name          = var.ec2_instance_name
#   key_pair_name          = var.ec2_key_pair_name
# }

module "alb" {
  source = "../modules/alb"

  project_name             = var.project_name
  public_subnet_ids        = module.vpc.public_subnet_ids
  alb_security_group_id    = module.security_groups.alb_sg_id
  app_security_group_id    = module.security_groups.app_sg_id
  app_instance_ids         = []
  alb_target_group_port    = var.alb_target_group_port
  alb_listener_port        = var.alb_listener_port
}