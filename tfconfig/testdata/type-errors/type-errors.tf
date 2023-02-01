# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "foo" {
  type = {"what":"the"}
  description = ["blah"]
}

output "foo" {
  description = ["blah"]
  sensitive   = "blah"
}

module "foo" {
  source = ["blah"]
  version = ["blah"]
}

provider "foo" {
  version = ["blah"]
}

resource "foo" "foo" {
  provider = ["nope"]
}

terraform {
  required_version = ["1.0.0"]
  required_providers {
    nope = ["definitely not"]
  }
}
