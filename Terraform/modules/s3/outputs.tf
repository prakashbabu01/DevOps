output "bucket_name" {
  description = "The name of the created S3 bucket"
  value       = aws_s3_bucket.s3_bucket.id
}

output "bucket_arn" {
  description = "The ARN of the created S3 bucket"
  value       = aws_s3_bucket.s3_bucket.arn
}



#output "versioning_status" {
#  description = "The versioning status of the bucket"
#  value       = aws_s3_bucket_versioning.s3_versioning.versioning_configuration[0].status
#}

#output "logging_status" {
#  description = "Indicates if bucket policy is creted"
#  value       = var.enable_bucketpolicy ? aws_s3_bucket_policy.logging_bucket_policy[0].id : 0
#}
