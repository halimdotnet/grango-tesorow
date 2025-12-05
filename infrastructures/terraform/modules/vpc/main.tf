variable "environment" {
  description = "Environment name"
  type        = string
}

variable "vpc_cidr" {
  description = "CIDR block for VPC"
  type        = string
}

variable "project_name" {
  description = "Project name for tagging"
  type        = string
  default     = "tesorow"
}

# VPC
resource "aws_vpc" "main" {
  cidr_block           = var.vpc_cidr
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name        = "${var.project_name}-vpc-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
    ManagedBy   = "terraform"
  }
}

# Internet Gateway
resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name        = "${var.project_name}-igw-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

# Public Subnets (2 AZs untuk high availability)
resource "aws_subnet" "public" {
  count                   = 2
  vpc_id                  = aws_vpc.main.id
  cidr_block              = cidrsubnet(var.vpc_cidr, 8, count.index)
  availability_zone       = data.aws_availability_zones.available.names[count.index]
  map_public_ip_on_launch = true

  tags = {
    Name        = "${var.project_name}-public-subnet-${var.environment}-${count.index + 1}"
    Environment = var.environment
    Project     = var.project_name
    Type        = "public"
  }
}

# Private Subnets (untuk aplikasi)
resource "aws_subnet" "private" {
  count             = 2
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(var.vpc_cidr, 8, count.index + 10)
  availability_zone = data.aws_availability_zones.available.names[count.index]

  tags = {
    Name        = "${var.project_name}-private-subnet-${var.environment}-${count.index + 1}"
    Environment = var.environment
    Project     = var.project_name
    Type        = "private"
  }
}

# Database Subnets
resource "aws_subnet" "database" {
  count             = 2
  vpc_id            = aws_vpc.main.id
  cidr_block        = cidrsubnet(var.vpc_cidr, 8, count.index + 20)
  availability_zone = data.aws_availability_zones.available.names[count.index]

  tags = {
    Name        = "${var.project_name}-db-subnet-${var.environment}-${count.index + 1}"
    Environment = var.environment
    Project     = var.project_name
    Type        = "database"
  }
}

# NAT Gateway - Cost: ~$32/month per environment - Uncomment for production
# resource "aws_eip" "nat" {
#   domain = "vpc"
#
#   tags = {
#     Name        = "${var.project_name}-nat-eip-${var.environment}"
#     Environment = var.environment
#     Project     = var.project_name
#   }
#
#   depends_on = [aws_internet_gateway.main]
# }
#
# resource "aws_nat_gateway" "main" {
#   allocation_id = aws_eip.nat.id
#   subnet_id     = aws_subnet.public[0].id
#
#   tags = {
#     Name        = "${var.project_name}-nat-${var.environment}"
#     Environment = var.environment
#     Project     = var.project_name
#   }
# }

# Route Table - Public
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main.id
  }

  tags = {
    Name        = "${var.project_name}-public-rt-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

# Route Table - Private
resource "aws_route_table" "private" {
  vpc_id = aws_vpc.main.id

  # Uncomment route below when NAT Gateway is enabled
  # route {
  #   cidr_block     = "0.0.0.0/0"
  #   nat_gateway_id = aws_nat_gateway.main.id
  # }

  tags = {
    Name        = "${var.project_name}-private-rt-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}

# Route Table Associations
resource "aws_route_table_association" "public" {
  count          = 2
  subnet_id      = aws_subnet.public[count.index].id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "private" {
  count          = 2
  subnet_id      = aws_subnet.private[count.index].id
  route_table_id = aws_route_table.private.id
}

# Data source untuk AZs
data "aws_availability_zones" "available" {
  state = "available"
}

# DB Subnet Group
resource "aws_db_subnet_group" "main" {
  name       = "${var.project_name}-db-subnet-group-${var.environment}"
  subnet_ids = aws_subnet.database[*].id

  tags = {
    Name        = "${var.project_name}-db-subnet-group-${var.environment}"
    Environment = var.environment
    Project     = var.project_name
  }
}