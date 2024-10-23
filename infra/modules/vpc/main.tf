resource "aws_vpc" "this" {
  cidr_block           = var.vpc_cidr
  enable_dns_support   = var.enable_dns_support
  enable_dns_hostnames = var.enable_dns_hostnames

  tags = {
    Name        = "vpc-${var.vpc_cidr}"
    Environment = "Production"
  }
}

resource "aws_internet_gateway" "this" {
  vpc_id = aws_vpc.this.id

  tags = {
    Name        = "igw-${aws_vpc.this.id}"
    Environment = "Production"
  }
}

resource "aws_subnet" "public-sn" {
  count                   = length(var.azs)
  vpc_id                  = aws_vpc.this.id
  cidr_block              = var.public_subnet_cidrs[count.index]
  availability_zone       = var.azs[count.index]
  map_public_ip_on_launch = true

  tags = {
    Name        = "public-sn-${count.index + 1}"
    Environment = "Production"
  }
}

resource "aws_subnet" "private-sn" {
  count             = length(var.azs)
  vpc_id            = aws_vpc.this.id
  cidr_block        = var.private_subnet_cidrs[count.index]
  availability_zone = var.azs[count.index]

  tags = {
    Name        = "private-sn-${count.index + 1}"
    Environment = "Production"
  }
}

resource "aws_route_table" "public-rt" {
  vpc_id = aws_vpc.this.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.this.id
  }

  tags = {
    Name        = "public-route-table"
    Environment = "Production"
  }
}

resource "aws_route_table_association" "public-rt-assoc" {
  count          = length(aws_subnet.public-sn)
  subnet_id      = aws_subnet.public-sn[count.index].id
  route_table_id = aws_route_table.public-rt.id
}

resource "aws_route_table" "private-rt" {
  vpc_id = aws_vpc.this.id

  tags = {
    Name        = "private-route-table"
    Environment = "Production"
  }
}

resource "aws_route_table_association" "private-rt-assoc" {
  count          = length(aws_subnet.private-sn)
  subnet_id      = aws_subnet.private-sn[count.index].id
  route_table_id = aws_route_table.private-rt.id
}
