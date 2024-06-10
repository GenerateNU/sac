resource "aws_launch_template" "this_template" {
  image_id = "ami-00beae93a2d981137"

  instance_type = "t2.micro"

  user_data = base64encode(templatefile("${path.module}/init.sh", {}))

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_autoscaling_group" "this_asg" {
    min_size = 1
    max_size = 1
    
    vpc_zone_identifier = [ var.subnet_id ]

    tag {
        key = "Name"
        value = "SAC Backend ASG"
        propagate_at_launch = true
    }

    launch_template {
      id = aws_launch_template.this_template.id
      version = "$Latest"
    }
}