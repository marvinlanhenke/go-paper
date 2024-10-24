variable "region" {
  description = "AWS region"
  type        = string
  default     = "eu-central-1"
}

variable "environment" {
  description = "The current deployment environment"
  type        = string
  default     = "production"
}

variable "db_username" {
  description = "The DB username"
  type        = string
}

variable "db_password" {
  description = "The DB password"
  type        = string
}

variable "db_name" {
  description = "The default database name"
  type        = string
  default     = "gopaper"
}
