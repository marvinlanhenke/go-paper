output "ecs_security_group_id" {
  description = "The connection endpoint for the RDS instance."
  value       = aws_security_group.ecs_service_sg.id
}
