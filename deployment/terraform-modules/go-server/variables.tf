variable "subnet_id" {
  type = string
  description = "ID of subnet to place EC2 in."
}

variable "envvars" {
  type = map(string)
  default = {}
}

