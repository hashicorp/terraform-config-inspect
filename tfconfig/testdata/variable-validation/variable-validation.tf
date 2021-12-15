variable "A" {
    default = "A default"
}

variable "B" {
    default = "B default"
    validation {
      condition = true
      error_message = "B error message"
    }
}

resource "null_resource" "A" {}
