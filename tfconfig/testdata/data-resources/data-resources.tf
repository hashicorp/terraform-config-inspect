# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "external" "foo" {
}

data "external" "bar" {
  provider = notexternal
}
