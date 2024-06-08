variable "allocated_storage" {
    type = number
    description = "Memory allocated in GB for this database"
}

variable "engine" {
    type = string
    description = "Choice of database engine"
}

variable "username" {
    type = string
    description = "User name of main database user"
}

variable "password" {
    type = string
    description = "Password of main database user"
}

variable "name" {
    type = string
    description = "Name of main database"
}

variable "instance_class" {
    type = string
    description = "Instance to provision for database"
}

variable "subnet_ids" {
    type = list(string)
    description = "List of subnets this database should be added to"
}
