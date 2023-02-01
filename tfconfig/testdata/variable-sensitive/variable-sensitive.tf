variable "A" {
    default = "A default"
}

variable "B" {
    default =  "B default"
    sensitive = true
}

resource "null_resource" "A" {}
