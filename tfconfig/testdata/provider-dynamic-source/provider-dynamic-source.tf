terraform {
  required_providers {
    both = {
      source  = var.some_source
      version = var.some_version
    }
    only_source = {
      source = "app.terraform.io/${var.some_source}"
    }
    only_version = {
      source  = "bar/baz"
      version = var.some_version
    }
  }
}

variable "some_source" {
  const = true
  type  = string
}

variable "some_version" {
  const = true
  type  = string
}
