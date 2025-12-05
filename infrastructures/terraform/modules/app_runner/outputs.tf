output "service_arn" {
  value = aws_apprunner_service.main.arn
}

output "service_url" {
  value = aws_apprunner_service.main.service_url
}

output "service_status" {
  value = aws_apprunner_service.main.status
}

output "connection_arn" {
  value = aws_apprunner_connection.github.arn
}