output "s3_static_site_enpoint" {
  description = "The website endpoint of the S3 bucket"
  value       = aws_s3_bucket_website_configuration.static_site_website_configuration.website_endpoint
}
