# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "foo" {
  type = true
  description = true
}

output "foo" {
  description = true
}

module "foo" {
  source = true
  version = true
}

provider "foo" {
  version = true
}

resource "foo" "foo" {
  provider = true
}

terraform {
  required_version = true
  required_providers {
    yep = true
  }
}
