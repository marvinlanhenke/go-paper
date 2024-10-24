variable "environment" {
  description = "The current deployment environment"
  type        = string
}

variable "subnets" {
  description = "A list of public subnet IDs where the ALB will be deployed"
  type        = list(string)
}

variable "vpc_id" {
  description = "The VPC id"
  type        = string
}
