variable "vpc_cidr" {
  description = "the CIDR block for the VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "azs" {
  description = "list of available AZs to create subnets in"
  type        = list(string)
  default     = ["eu-central-1a", "eu-central-1b"]
}

variable "public_subnet_cidrs" {
  description = "list of CIDR blocks for public subnets"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24"]
}

variable "private_subnet_cidrs" {
  description = "list of CIDR blocks for public subnets"
  type        = list(string)
  default     = ["10.0.3.0/24", "10.0.4.0/24"]
}

variable "enable_dns_support" {
  description = "enable DNS support in the vpc"
  type        = bool
  default     = true
}

variable "enable_dns_hostnames" {
  description = "enable DNS hostnames in the vpc"
  type        = bool
  default     = true
}
