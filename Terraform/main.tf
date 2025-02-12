provider "aws" {
  alias   = "selected"
  profile = var.aws_providers[var.selected_profile].profile
  region  = var.aws_providers[var.selected_profile].region
}


module "s3_buckets" {
  for_each = var.s3_buckets

  source              = "./modules/s3"
  bucket_name         = each.value.name
  enable_bucketpolicy = each.value.enable_bucketpolicy
  tags                = each.value.tags
  aws_account_id      = var.aws_account_id
  providers = {
    aws = aws.selected
  }
}

# Output the created bucket names
output "created_s3_buckets" {
  description = "List of created S3 bucket names , S3 Bucket ARNs , Logging S3 Buckets"
  value = { for k, v in module.s3_buckets : k => {
    "S3 Bucket_name" = v.bucket_name
    "S3 Bucket_ARN"  = v.bucket_arn
    #"logging_status" = v.logging_status != null ? v.logging_status : "No Bucket Policy enabled" 
    }
  }
}

