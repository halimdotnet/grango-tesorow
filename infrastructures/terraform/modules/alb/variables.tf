variable "project_name" {
  type = string
}

variable "public_subnet_ids" {
  type = list(string)
}

variable "alb_security_group_id" {
  type = string
}

variable "app_security_group_id" {
  type = string
}

variable "app_instance_ids" {
  type    = list(string)
  default = []
}

variable "alb_target_group_port" {
  type = number
}

variable "alb_listener_port" {
  type = number
}