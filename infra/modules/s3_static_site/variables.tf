variable "bucket_name" {
  description = "the name of the S3 bucket"
  type        = string
  default     = "ml-sa-s3-static-site"
}

variable "force_destroy" {
  description = "whether to force destroy the bucket even if it contains objects"
  type        = bool
  default     = true
}

variable "website_index_document" {
  description = "the index document for the website"
  type        = string
  default     = "index.html"
}

variable "website_error_document" {
  description = "the error document for the website"
  type        = string
  default     = "error.html"
}
