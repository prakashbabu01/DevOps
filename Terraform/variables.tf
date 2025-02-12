variable "aws_profile" {
  description = "AWS CLI profile name"
  type        = string
}

variable "aws_region" {
  description = "AWS region to deploy the resources"
  type        = string
}



variable "aws_account_id" {
  description = "AWS Account ID to allow log writes"
  type        = string
}


variable "s3_buckets" {
  description = "A map of objects defining multiple S3 buckets"
  type = map(object({
    name                = string
    enable_bucketpolicy = bool
    tags                = map(string)
  }))
}

variable "aws_providers" {
  description = "A map of AWS providers with profile and region"
  type = map(object({
    profile = string
    region  = string
  }))
}

variable "selected_profile" {
  description = "Select a profile from aws providers with this value"
  type        = string
}
