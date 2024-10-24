output "repository_url" {
  description = "the URI of the ECR repository"
  value       = aws_ecr_repository.this.repository_url
}

