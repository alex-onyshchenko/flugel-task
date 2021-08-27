terraform {
  required_version = ">= 0.12.26"
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

locals {
  bucket_prefix = "bucket-"
}

resource "aws_s3_bucket" "bucket" {
  bucket_prefix = local.bucket_prefix
}

resource "aws_s3_bucket_object" "bucket_objects" {
  count = 2

  bucket  = aws_s3_bucket.bucket.id
  key     = "test${count.index + 1}.txt"
  content = timestamp()
}

output "bucket_id" {
  value = aws_s3_bucket.bucket.id
}

output "test1_content" {
  value = aws_s3_bucket_object.bucket_objects[0].content
}

output "test2_content" {
  value = aws_s3_bucket_object.bucket_objects[1].content
}
