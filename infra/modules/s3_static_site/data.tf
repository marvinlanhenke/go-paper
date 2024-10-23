data "aws_iam_policy_document" "public_read" {
  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.static_site.arn}/*"]
    principals {
      type        = "AWS"
      identifiers = ["*"]
    }
  }
}
