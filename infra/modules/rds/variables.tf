variable "vpc_id" {
  description = "The VPC id"
  type        = string
}

variable "environment" {
  description = "The deployment environment"
  type        = string
}

variable "private_subnets" {
  description = "List of private subnet IDs"
  type        = list(string)
}

variable "security_group_ids" {
  description = "List of security group ids"
  type        = list(string)
}

variable "db_name" {
  description = "The name of the default database to create."
  type        = string
}

variable "db_username" {
  description = "The database admin username."
  type        = string
}

variable "db_password" {
  description = "The database admin password."
  type        = string
  sensitive   = true
}

variable "db_instance_class" {
  description = "The instance type for the RDS instance"
  type        = string
  default     = "db.t3.medium"
}

variable "db_engine_version" {
  description = "The version of the PostgreSQL engine"
  type        = string
  default     = "16.3"
}

variable "allocated_storage" {
  description = "The allocated storage in gigabytes."
  type        = number
  default     = 20
}

variable "max_allocated_storage" {
  description = "The maximum allocated storage in gigabytes for storage autoscaling."
  type        = number
  default     = 100
}

variable "multi_az" {
  description = "Specifies if the RDS instance is a Multi-AZ deployment."
  type        = bool
  default     = false
}

variable "storage_type" {
  description = "The storage type for the RDS instance."
  type        = string
  default     = "gp2"
}

variable "engine" {
  description = "The database engine to use."
  type        = string
  default     = "postgres"
}

variable "backup_retention_period" {
  description = "The number of days to retain backups."
  type        = number
  default     = 7
}
