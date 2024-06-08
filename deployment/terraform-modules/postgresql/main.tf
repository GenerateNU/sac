resource "aws_db_instance" "this" {
    allocated_storage = var.allocated_storage
    engine = var.engine
    instance_class = var.instance_class
    username = var.username
    password = var.password
    db_name = var.name
    multi_az = false
    
    skip_final_snapshot = true

    tags = {
        Name = "${var.name} RDS instance"
    }

    db_subnet_group_name = aws_db_subnet_group.this_subnet_group.name
}

resource "aws_db_subnet_group" "this_subnet_group" {
    name = "${var.name}-database-subnet-group"
    subnet_ids = var.subnet_ids

    tags = {
        Name = "${var.name} DB Subnet Group"
    }
}