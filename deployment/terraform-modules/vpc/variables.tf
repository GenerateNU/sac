variable "vpc_name" {
    description = "The name to use for the VPC"
    type = string
    default = "SAC VPC"
}

variable "public_subnet_cidr" {
    type = string
    description = "Public Subnet CIDR value"
    default = "10.0.1.0/24"
}

variable "private_subnet_cidrs" {
    type = list(string)
    description = "Private Subnet CIDR values"
    default = ["10.0.2.0/24", "10.0.3.0/24"]
}

variable "private_azs" {
    type = list(string)
    description = "Availability Zones"
    default = [ "us-east-1a", "us-east-1b" ]
}