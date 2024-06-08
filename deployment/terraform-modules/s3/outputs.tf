output "bucket_name" {
  description = "Name of the deployed S3 bucket"
  value = aws_s3_bucket.this.bucket
}

output "iam_access_key" {
    description = "Access key ID for the S3 bucket's IAM user"
    value = aws_iam_access_key.this_user_key.id
}

output "iam_secret_key" {
    description = "Secret access key for the S3 bucket's IAM user"
    value = aws_iam_access_key.this_user_key.secret
}
