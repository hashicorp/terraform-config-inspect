# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

module "foo" {
  source = "foo/bar/baz"
  version = "1.0.2"

  unused = 2
}

module "bar" {
  source = "./child"

  unused = 1
}
