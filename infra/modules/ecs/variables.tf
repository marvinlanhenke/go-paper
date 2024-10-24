variable "region" {
  description = "AWS region"
  type        = string
}

variable "environment" {
  description = "deployment environment"
  type        = string
}

variable "vpc_id" {
  description = "The VPC id"
  type        = string
}

variable "private_subnets" {
  description = "A list of private subnet IDs where ECS tasks will be deployed"
  type        = list(string)
}

variable "alb_target_group_arn" {
  description = "The ARN of the ALB target group"
  type        = string
}

variable "alb_sg_id" {
  description = "The ID of the ALB security group"
  type        = string
}

variable "container_image" {
  description = "The docker image for the ECS task"
  type        = string
}

variable "environment_variables" {
  description = "Environment variables for the container"
  type        = map(string)
  default     = {}
}

variable "container_port" {
  description = "The port on which the container listens"
  type        = number
  default     = 8080
}

variable "desired_count" {
  description = "Desired number of ECS tasks"
  type        = number
  default     = 2
}

variable "cpu" {
  description = "CPU units for the task"
  type        = string
  default     = "256"
}

variable "memory" {
  description = "Memory (MiB) for the task"
  type        = string
  default     = "512"
}
