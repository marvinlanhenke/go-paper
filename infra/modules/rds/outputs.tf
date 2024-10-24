output "db_instance_endpoint" {
  description = "The connection endpoint for the RDS instance."
  value       = aws_db_instance.this.endpoint
}

output "db_instance_port" {
  description = "The port number on which the database accepts connections."
  value       = aws_db_instance.this.port
}

output "db_instance_id" {
  description = "The identifier of the RDS instance."
  value       = aws_db_instance.this.id
}

output "db_credentials_secret_arn" {
  description = "The ARN of the Secrets Manager secret containing DB credentials."
  value       = aws_secretsmanager_secret.db_credentials.arn
}
