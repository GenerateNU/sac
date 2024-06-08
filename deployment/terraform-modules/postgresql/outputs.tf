output "port" {
    value = aws_db_instance.this.port
}

output "host" {
    value = aws_db_instance.this.address
}