output "vpc_id" {
  description = "The ID of the VPC."
  value       = aws_vpc.this.id
}

output "public_subnets" {
  description = "List of public subnet IDs."
  value       = aws_subnet.public-sn[*].id
}

output "private_subnets" {
  description = "List of private subnet IDs."
  value       = aws_subnet.private-sn[*].id
}
