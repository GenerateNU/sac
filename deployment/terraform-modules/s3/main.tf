resource "aws_s3_bucket" "this" {
    tags = {
        Name = "SAC - ${var.bucket_name}"
        Environment = "prod"
    }
}

resource "aws_iam_user" "this_user" {
    name = var.user_name

    tags = {
        Name = "SAC S3 bucket - Access user"
    }
}

resource "aws_iam_access_key" "this_user_key" {
    user = aws_iam_user.this_user.name
}

resource "aws_s3_bucket_policy" "this_access_policy" {
  bucket = aws_s3_bucket.this.id
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
        {
            Effect = "Allow",
            Principal = {
                AWS = "${aws_iam_user.this_user.arn}"
            }
            Action = [
                "s3:*"
            ]
            Resource = [
                "${aws_s3_bucket.this.arn}/*",
                "${aws_s3_bucket.this.arn}"
            ]
        }
    ]
  })
}

resource "aws_iam_user_policy_attachment" "iam_policy" {
    policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
    user = aws_iam_user.this_user.name
}