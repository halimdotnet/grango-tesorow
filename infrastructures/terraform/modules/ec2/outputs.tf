output "instance_id" {
  value = aws_instance.app.id
}

output "private_ip" {
  value = aws_instance.app.private_ip
}

output "public_ip" {
  value = aws_instance.app.public_ip
}

output "availability_zone" {
  value = aws_instance.app.availability_zone
}