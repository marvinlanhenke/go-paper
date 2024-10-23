variable "vpc_id" {
  description = "the VPC id"
  type        = string
}

variable "environment" {
  description = "deployment environment"
  type        = string
  default     = "production"
}

variable "environment_variables" {
  description = "environment variables for the container"
  type        = map(string)
  default     = {}
}

variable "private_subnets" {
  description = "list of private subnet IDs where ECS tasks will be deployed"
  type        = list(string)
}

variable "alb_target_group_arn" {
  description = "the ARN of the ALB target group"
  type        = string
}

variable "alb_sg_id" {
  description = "the ID of the ALB security group"
  type        = string
}

variable "container_image" {
  description = "the docker image for the ECS task"
  type        = string
}

variable "container_port" {
  description = "the port on which the container listens"
  type        = number
  default     = 8080
}

variable "desired_count" {
  description = "desired number of ECS tasks"
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
