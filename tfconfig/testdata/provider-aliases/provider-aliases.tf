# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

terraform {
  required_providers {
    bleep = {
      configuration_aliases = [ bleep.bloop ]
    }
  }
}

provider "foo" {
  alias = "blue"
}

provider "foo" {
  alias = "red"
}

provider "bar" {
}

provider "bar" {
  alias = "yellow"
}

provider "baz" {
}

provider "empty" {
  alias = ""
}
