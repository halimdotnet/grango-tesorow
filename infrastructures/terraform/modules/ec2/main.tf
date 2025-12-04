resource "aws_instance" "app" {
  ami                    = var.ami_id
  instance_type          = var.instance_type
  subnet_id              = var.private_subnet_id
  vpc_security_group_ids = [var.app_security_group_id]

  associate_public_ip_address = false

  root_block_device {
    volume_type           = "gp3"
    volume_size           = 20
    delete_on_termination = true
    encrypted             = true
  }

  monitoring = true

  tags = {
    Name = var.instance_name
  }

  lifecycle {
    create_before_destroy = true
  }
}