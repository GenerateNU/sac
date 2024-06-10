source "amazon-ebs" "echotest" {
  ami_name      = "echotest-image-${local.suffix}"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami = "ami-04b70fa74e45c3917"
  ssh_username =  "ubuntu"
}

build {
  name    = "echotest-image-${local.suffix}"
  sources = [
    "source.amazon-ebs.echotest"
  ]

  provisioner "shell" {
    script = "${path.root}/docker-install.sh"
  }
}
