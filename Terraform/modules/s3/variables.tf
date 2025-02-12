variable "bucket_name" {
  description = "The name of the S3 bucket"
  type        = string
}

variable "tags" {
  description = "Tags to assign to the bucket"
  type        = map(string)
  default     = {}
}

variable "enable_bucketpolicy" {
  description = "create bucket policy"
  type        = bool
  default     = false
}



variable "aws_account_id" {
  description = "AWS Account ID to allow log writes"
  type        = string
}

