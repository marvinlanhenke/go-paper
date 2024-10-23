output "alb_arn" {
  description = "the ARN of the ALB"
  value       = aws_lb.this.arn
}

output "alb_dns" {
  description = "the DNS of the ALB"
  value       = aws_lb.this.dns_name
}

output "alb_sg_id" {
  description = "the ID of the ALB security group"
  value       = aws_security_group.alb-sg.id
}

output "alb_target_group_arn" {
  description = "the ARN of the ALB target group"
  value       = aws_lb_target_group.this.arn
}
