variable "bucket_name" {
  description = "the name of the S3 bucket"
  type        = string
  default     = "ml-sa-s3-static-site"
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
