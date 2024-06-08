packer {
  required_plugins {
    amazon = {
      version = ">= 1.2.8"
      source  = "github.com/hashicorp/amazon"
    }
  }
}

variable 
source "amazon-ebs" "webserver" {
  ami_name      = "webserver-image-${formatdate( \"YYYY-MM-DD-hhmmss\", timestamp())}"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami = "ami-04b70fa74e45c3917"
  ssh_username =  "ubuntu"
}

source "amazon-ebs" "redis" {
  ami_name      = "redis-image-${timestamp()}"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami = "ami-04b70fa74e45c3917"
  ssh_username =  "ubuntu"
}

source "amazon-ebs" "opensearch" {
  ami_name      = "opensearch-image-${timestamp()}"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami = "ami-04b70fa74e45c3917"
  ssh_username =  "ubuntu"
}

build {
  name    = "webserver-image-${timestamp()}"
  sources = [
    "source.amazon-ebs.webserver"
  ]

  provisioner "shell" {
    script = "${path.root}/webserver-init.sh"
  }
}

build {
  name    = "redis-image-${timestamp()}"
  sources = [
    "source.amazon-ebs.webserver"
  ]

  provisioner "shell" {
    script = "${path.root}/docker-install.sh"
  }

  provisioner "shell" {
    inline = [
        "cd ~",
        "git clone -b prod-shenanigans https://github.com/GenerateNU/sac",
        "mv sac/deployment/docker-composes/redis.yml ~/docker-compose.yml"
    ]
  }
}

build {
  name    = "opensearch-image-${timestamp()}"
  sources = [
    "source.amazon-ebs.webserver"
  ]

  provisioner "shell" {
    script = "${path.root}/docker-install.sh"
  }

  provisioner "shell" {
    inline = [
        "cd ~",
        "git clone -b prod-shenanigans https://github.com/GenerateNU/sac",
        "mv sac/deployment/docker-composes/opensearch.yml ~/docker-compose.yml"
    ]
  }
}