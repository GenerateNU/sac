output "vpc_id" {
    value = aws_vpc.main.id
    description = "ID of the VPC"
}

output "public_subnet_id" {
    value = aws_subnet.public_subnet.id
    description = "ID of the VPC's public subnet"
}

output "private_subnet_ids" {
    value = aws_subnet.private_subnets[*].id
    description = "IDs of the VPC's private subnets"
}