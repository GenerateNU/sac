provider "aws" {
    region = "us-east-1"
}

terraform {
    backend "s3" {
        bucket = "terraform-state-sac-495a"
        key = "global/s3/terraform.tfstate"
        region = "us-east-1"

        dynamodb_table = "terraform-state-locks"
        encrypt = true
    }
}

data "aws_secretsmanager_secret" "database-secrets" {
    name = "prod/DatabaseSecrets"
}

data "aws_secretsmanager_secret_version" "database-secrets-version" {
  secret_id = data.aws_secretsmanager_secret.database-secrets.id
}


module "sac_vpc" {
    source = "./terraform-modules/vpc/"
}

module "sac_s3_bucket" {
    source = "./terraform-modules/s3"

    bucket_name = "sac-s3-bucket-64cf"
    user_name = "sac-s3-bucket-64cf-user"
}

module "sac_postgresql" {
    source = "./terraform-modules/postgresql"

    allocated_storage = 10
    engine = "postgres"
    instance_class = "db.t2.micro"
    name = "sac"
    subnet_ids = module.sac_vpc.private_subnet_ids

    username = jsondecode(data.aws_secretsmanager_secret_version.database-secrets-version.secret_string)["SAC_DB_USERNAME"]
    password = jsondecode(data.aws_secretsmanager_secret_version.database-secrets-version.secret_string)["SAC_DB_PASSWORD"]
}

module "sac_goserver" {
  source = "./terraform-modules/go-server"

  subnet_id = module.sac_vpc.public_subnet_id
}