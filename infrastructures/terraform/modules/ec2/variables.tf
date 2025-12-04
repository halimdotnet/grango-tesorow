variable "project_name" {
  type = string
}

variable "instance_type" {
  type = string
}

variable "ami_id" {
  type = string
}

variable "private_subnet_id" {
  type = string
}

variable "app_security_group_id" {
  type = string
}

variable "instance_name" {
  type = string
}

variable "key_pair_name" {
  type      = string
  default   = ""
  sensitive = true
}