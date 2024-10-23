variable "repository_name" {
  description = "name of the ECR repository"
  type        = string
}

variable "force_delete" {
  description = "whether to force delete the repository even if it contains objects"
  type        = bool
  default     = true
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
