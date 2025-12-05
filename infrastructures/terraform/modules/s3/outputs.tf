output "documents_bucket_id" {
  description = "Documents bucket ID"
  value       = aws_s3_bucket.documents.id
}

output "documents_bucket_arn" {
  description = "Documents bucket ARN"
  value       = aws_s3_bucket.documents.arn
}

output "documents_bucket_name" {
  description = "Documents bucket name"
  value       = aws_s3_bucket.documents.bucket
}

output "documents_bucket_region" {
  description = "Documents bucket region"
  value       = aws_s3_bucket.documents.region
}

output "logs_bucket_id" {
  description = "Logs bucket ID"
  value       = aws_s3_bucket.logs.id
}

output "logs_bucket_arn" {
  description = "Logs bucket ARN"
  value       = aws_s3_bucket.logs.arn
}

output "logs_bucket_name" {
  description = "Logs bucket name"
  value       = aws_s3_bucket.logs.bucket
}