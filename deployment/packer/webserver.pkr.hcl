source "amazon-ebs" "webserver" {
  ami_name      = "webserver-image-${local.suffix}"
  instance_type = "t2.micro"
  region        = "us-east-1"
  source_ami = "ami-04b70fa74e45c3917"
  ssh_username =  "ubuntu"
}

build {
  name    = "webserver-image-${local.suffix}"
  sources = [
    "source.amazon-ebs.webserver"
  ]

  provisioner "shell" {
    script = "${path.root}/webserver-init.sh"
  }
}