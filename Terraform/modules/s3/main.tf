
terraform {
  required_providers {
    aws = {
      source                = "hashicorp/aws"
      version               = "~> 4.0"
      configuration_aliases = [aws]
      # AWS provider Configuration will be passed from Root main.tf , earlier we used to keep an empty block and now deprecated.
    }
  }
}

resource "aws_s3_bucket" "s3_bucket" {
  bucket = var.bucket_name
  tags   = var.tags
}

# Enable versioning
resource "aws_s3_bucket_versioning" "s3_versioning" {
  bucket = aws_s3_bucket.s3_bucket.id

  versioning_configuration {
    status = "Enabled"
  }
}

# Enable encryption
resource "aws_s3_bucket_server_side_encryption_configuration" "s3_encryption" {
  bucket = aws_s3_bucket.s3_bucket.id

  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}

# Get the logging bucket dynamically (if logging is enabled)
#data "aws_s3_bucket" "logging_bucket" {
#  count  = var.enable_logging ? 1 : 0
#  bucket = var.logging_bucket_name
#}

# Enable access logging (optional)
#resource "aws_s3_bucket_logging" "s3_logging" {
#  count = var.enable_logging ? 1 : 0

#  bucket        = aws_s3_bucket.s3_bucket.id
#  target_bucket = data.aws_s3_bucket.logging_bucket[0].id
#  target_prefix = "${var.bucket_name}/logs/"
#}

#resource "aws_s3_bucket_policy" "logging_bucket_policy" {
#  count = var.enable_bucketpolicy ? 1 : 0
#  #bucket = data.aws_s3_bucket.logging_bucket[0].id
#  bucket = aws_s3_bucket.s3_bucket.id

#  policy = jsonencode({
#    Version = "2012-10-17"
#    Statement = [
#      {
#        Effect = "Allow"
#        Principal = {
#          Service = "logging.s3.amazonaws.com"
#        }
#        Action   = "s3:PutObject"
#        Resource = ["arn:aws:s3:::prakash-kops", "arn:aws:s3:::prakash-kops/*"]
#        Condition = {
#          ArnLike = {
#            "aws:SourceArn" = aws_s3_bucket.s3_bucket.arn
#          }
#          StringEquals = {
#            "aws:SourceAccount" = var.aws_account_id
#          }
#        }
#      }
#    ]
#  })
#}
