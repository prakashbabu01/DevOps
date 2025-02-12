

module "my_ec2_instance" {
  source = "terraform-aws-modules/ec2-instance/aws"

  name = "best-practice-ec2"

  instance_type          = "t3.micro"
  ami                    = "ami-085ad6ae776d8f09c"                # Example AMI, replace with latest Amazon Linux 2 AMI
  key_name               = aws_key_pair.my-test-key-pair.key_name # Replace with your SSH key pair
  vpc_security_group_ids = [module.my-security_group.security_group_id]
  subnet_id              = module.vpc.public_subnets[0]

  root_block_device = [{
    volume_size = 30
    volume_type = "gp3"
    encrypted   = true
  }]

  tags = {
    Name        = "BestPracticeEC2"
    Environment = "Production"
    Owner       = "DevOps"
  }
}

resource "aws_key_pair" "my-test-key-pair" {
  key_name   = "my-test-key-pair"        # Name for your key pair
  public_key = file("~/.ssh/id_rsa.pub") # Path to your public key file
}

module "my-security_group" {
  source = "terraform-aws-modules/security-group/aws"

  name        = "ec2-security-group"
  description = "Security group for EC2 instance"
  vpc_id      = module.vpc.vpc_id

  ingress_with_cidr_blocks = [
    {
      from_port   = 22
      to_port     = 22
      protocol    = "tcp"
      cidr_blocks = "0.0.0.0/0" # Restrict this in production
    }
  ]

  egress_with_cidr_blocks = [{
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = "0.0.0.0/0"
  }]
}

module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = "best-practice-vpc"
  cidr = "10.0.0.0/16"

  enable_dns_support   = true
  enable_dns_hostnames = true

  azs             = ["us-east-1a", "us-east-1b"]
  public_subnets  = ["10.0.1.0/24", "10.0.2.0/24"]
  private_subnets = ["10.0.3.0/24", "10.0.4.0/24"]

  enable_nat_gateway = true
  enable_vpn_gateway = false
}
