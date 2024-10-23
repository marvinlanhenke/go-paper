resource "aws_s3_bucket" "static_site" {
  bucket = var.bucket_name
}

resource "aws_s3_bucket_public_access_block" "public_access_block" {
  bucket                  = aws_s3_bucket.static_site.id
  block_public_acls       = true
  block_public_policy     = false
  ignore_public_acls      = true
  restrict_public_buckets = false
}

resource "aws_s3_bucket_policy" "public_read" {
  bucket = aws_s3_bucket.static_site.id
  policy = data.aws_iam_policy_document.public_read.json
}

resource "aws_s3_bucket_website_configuration" "static_site_website_configuration" {
  bucket = aws_s3_bucket.static_site.id
  index_document {
    suffix = "index.html"
  }
}
