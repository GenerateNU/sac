source "amazon-ebs" "opensearch" {
  ami_name      = "opensearch-image-${local.suffix}"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami = "ami-04b70fa74e45c3917"
  ssh_username =  "ubuntu"
}

build {
  name    = "opensearch-image-${local.suffix}"
  sources = [
    "source.amazon-ebs.opensearch"
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
