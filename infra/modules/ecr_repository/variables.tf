variable "repository_name" {
  description = "name of the ECR repository"
  type        = string
}

variable "image_tag_mutability" {
  description = "the image tag mutability setting"
  type        = string
  default     = "MUTABLE"
}

variable "image_scanning_configuration" {
  description = "image sanning configuration"
  type = object({
    scan_on_push = bool
  })
  default = {
    scan_on_push = true
  }
}
