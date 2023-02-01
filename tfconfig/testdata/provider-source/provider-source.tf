# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

terraform {
  required_providers {
    foo = "2.0.0"
    bat = {
      version = "1.0.0"
    }
  }
}

terraform {
  required_providers {
    bat = {
      source  = "baz/bat"
    }
  }
}
