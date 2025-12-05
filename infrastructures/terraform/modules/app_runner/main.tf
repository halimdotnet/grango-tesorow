resource "aws_apprunner_service" "main" {
  service_name = "${var.project_name}-app-runner"

  source_configuration {
    authentication_configuration {
      connection_arn = aws_apprunner_connection.github.arn
    }

    code_repository {
      repository_url = var.github_repo_url
      source_code_version {
        type   = "BRANCH"
        value  = var.github_branch
      }
      code_configuration {
        configuration_source = "API"
        code_configuration_values {
          runtime                = "GO_1"
          build_command          = "go build -o grango-tesorow cmd/rest/main.go"
          start_command          = "./grango-tesorow"
          port                   = var.app_port
          runtime_environment_variables = var.environment_variables
        }
      }
    }
  }

  instance_configuration {
    instance_role_arn = aws_iam_role.apprunner_role.arn
    cpu               = var.cpu
    memory            = var.memory
  }

  network_configuration {
    egress_configuration {
      egress_type       = "VPC"
      vpc_connector_arn = aws_apprunner_vpc_connector.main.arn
    }
  }

  tags = {
    Name = "${var.project_name}-app-runner"
  }

  depends_on = [aws_apprunner_connection.github]
}

resource "aws_apprunner_connection" "github" {
  provider_type = "GITHUB"
  connection_name = "${var.project_name}-github-connection"
}

resource "aws_apprunner_vpc_connector" "main" {
  vpc_connector_name = "${var.project_name}-vpc-connector"
  subnets            = var.private_subnet_ids
  security_groups    = [var.app_security_group_id]

  tags = {
    Name = "${var.project_name}-vpc-connector"
  }
}

resource "aws_iam_role" "apprunner_role" {
  name = "${var.project_name}-apprunner-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "tasks.apprunner.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "apprunner_policy" {
  name = "${var.project_name}-apprunner-policy"
  role = aws_iam_role.apprunner_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Resource = "*"
      }
    ]
  })
}