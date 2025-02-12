aws_profile    = "awscli"
aws_region     = "us-east-1"
aws_account_id = 107440719127

s3_buckets = {
  bucket1 = {
    name                = "prakashbabuposam-bucket-107440719127-1"
    enable_bucketpolicy = true
    tags = {
      Environment = "Dev"
      Project     = "TerraformDemo"
    }

  }

  bucket2 = {
    name                = "prakashbabuposam-bucket-107440719127-2"
    enable_bucketpolicy = false
    tags = {
      Environment = "Prod"
      Project     = "TerraformDemo"
    }
  }

  bucket3 = {
    name                = "prakashbabuposam-bucket-107440719127-3"
    enable_bucketpolicy = true
    tags = {
      Environment = "QA"
      Project     = "TerraformDemo"
    }
  }
}


aws_providers = {
  awscli = {
    profile = "awscli"
    region  = "us-east-1"
  }

  jenkins = {
    profile = "jenkins"
    region  = "us-east-1"
  }
}

selected_profile = "awscli"
