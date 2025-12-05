variable "project_name" {
  type = string
}

variable "github_repo_url" {
  type = string
  description = "GitHub repository URL (e.g., https://github.com/username/repo)"
}

variable "github_branch" {
  type = string
  default = "main"
}

variable "app_port" {
  type = number
  default = 8081
}

variable "cpu" {
  type = string
  default = "0.25 vCPU"
}

variable "memory" {
  type = string
  default = "512"
}

variable "environment_variables" {
  type = map(string)
  default = {}
}

variable "private_subnet_ids" {
  type = list(string)
}

variable "app_security_group_id" {
  type = string
}