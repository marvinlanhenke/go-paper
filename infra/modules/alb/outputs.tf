output "alb_arn" {
  description = "the ARN of the ALB"
  value       = aws_lb.this.arn
}

output "alb_dns" {
  description = "the DNS of the ALB"
  value       = aws_lb.this.dns_name
}
