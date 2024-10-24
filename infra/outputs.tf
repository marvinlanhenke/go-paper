output "ecr_repository_url" {
  description = "the URI of the ECR repository"
  value       = module.ecr.repository_url
}

output "alb_dns" {
  description = "the DNS of the ALB"
  value       = module.alb.alb_dns
}
