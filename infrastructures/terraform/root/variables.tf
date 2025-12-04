variable "aws_region" {
  type    = string
  default = "ap-southeast-1"
}

variable "aws_profile" {
  type    = string
  default = "default"
}

variable "project_name" {
  type = string
}

variable "environment" {
  type = string
}

variable "vpc_cidr" {
  type = string
}

variable "availability_zones" {
  type = list(string)
}

variable "public_subnet_cidrs" {
  type = list(string)
}

variable "private_subnet_cidrs" {
  type = list(string)
}

variable "tags" {
  type    = map(string)
  default = {}
}

variable "rds_master_password" {
  description = "RDS master password"
  type        = string
  sensitive   = true
}

variable "rds_identifier" {
  type = string
}

variable "rds_engine_version" {
  type = string
}

variable "rds_instance_class" {
  type = string
}

variable "rds_allocated_storage" {
  type = number
}

variable "rds_max_allocated_storage" {
  type = number
}

variable "rds_database_name" {
  type = string
}

variable "rds_master_username" {
  type      = string
  sensitive = true
}

variable "rds_backup_retention_days" {
  type = number
}

variable "rds_multi_az" {
  type = bool
}

variable "ec2_instance_type" {
  type = string
}

variable "ec2_instance_name" {
  type = string
}

variable "ec2_key_pair_name" {
  type      = string
  default   = ""
  sensitive = true
}

variable "alb_target_group_port" {
  type = number
}

variable "alb_listener_port" {
  type = number
}