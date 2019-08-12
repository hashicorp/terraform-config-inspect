terraform {
  required_version = ">= 0.11.0"

  required_providers {
    aws = ">= 2.7.0"
  }

  backend "s3" {
    bucket = "mybucket"
    key    = "path/to/my/key"
    region = "us-east-1"
  }

  ignored = 1
}

provider "aws" {
  region = "us-east-1"
}

variable "foo" {
  default = "bar"
}

output "foo" {
  value = "${var.foo}"
}
