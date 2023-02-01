# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "A" {
  default = "A default"
}

variable "B" {
  description = "The B variable"
}

output "A" {
  value = "${var.A}"
}

output "B" {
  description = "I am B"
  value       = "${var.A}"
}

resource "null_resource" "A" {}
resource "null_resource" "B" {}

module "foo" {
  source  = "foo/bar/baz"
  version = "1.0.2"

  unused = 2
}
