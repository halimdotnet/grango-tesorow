terraform {
  required_version = ">= 1.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.5"
    }
  }
}

provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      Environment = "staging"
      Project     = "tesorow"
      ManagedBy   = "terraform"
    }
  }
}

module "vpc" {
  source = "../../modules/vpc"

  environment  = "staging"
  vpc_cidr     = "10.1.0.0/16"
  project_name = "tesorow"
}

module "security_groups" {
  source = "../../modules/security-groups"

  environment  = "staging"
  vpc_id       = module.vpc.vpc_id
  project_name = "tesorow"
}

module "iam" {
  source = "../../modules/iam"

  environment  = "staging"
  project_name = "tesorow"
}

module "rds" {
  source = "../../modules/rds"

  environment            = "staging"
  project_name           = "tesorow"
  vpc_id                 = module.vpc.vpc_id
  database_subnet_ids    = module.vpc.database_subnet_ids
  db_security_group_id   = module.security_groups.db_security_group_id
  db_subnet_group_name   = module.vpc.db_subnet_group_name


  instance_class         = "db.t3.micro"
  allocated_storage      = 20
  backup_retention_period = 7
}

module "s3" {
  source = "../../modules/s3"

  environment  = "staging"
  project_name = "tesorow"
}

# module "alb" {
#   source = "../../modules/alb"
#
#   environment           = "staging"
#   project_name          = "tesorow"
#   vpc_id                = module.vpc.vpc_id
#   public_subnet_ids     = module.vpc.public_subnet_ids
#   alb_security_group_id = module.security_groups.alb_security_group_id
# }