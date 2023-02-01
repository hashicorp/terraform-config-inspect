# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "A" {
    default = "A default"
}

variable "B" {
    description = "The B variable"
}

variable "C" {
    description = "The C variable"
}

output "A" {
    value = "${var.A}"
}

output "B" {
    description = "I am B"
    value = "${var.A}"
    sensitive = false
}

output "C" {
    description = "C is sensitive"
    value = "${var.C}"
    sensitive = true
}

resource "null_resource" "A" {}
resource "null_resource" "B" {}
resource "null_resource" "C" {}
