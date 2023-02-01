# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "A" {
    default = "A default"
}

variable "B" {
    default =  "B default"
    sensitive = true
}

resource "null_resource" "A" {}
