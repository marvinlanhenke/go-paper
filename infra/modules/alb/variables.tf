variable "environment" {
  description = "deployment environment"
  type        = string
  default     = "production"
}

variable "subnets" {
  description = "list of public subnet IDs where the ALB will be deployed"
  type        = list(string)
}

variable "vpc_id" {
  description = "the VPC id"
  type        = string
}

variable "idle_timeout" {
  description = "idle timeout for the ALB in seconds"
  type        = number
  default     = 60
}
