# Copyright IBM Corp. 2018, 2025
# SPDX-License-Identifier: MPL-2.0

# Test file for provider blocks with different label configurations

required_providers {
  aws = {
    source  = "hashicorp/aws"
    version = "~> 5.0"
  }

  random = {
    source  = "hashicorp/random"
    version = "~> 3.0"
  }

  null = {
    source = "hashicorp/null"
  }
}

# Terraform Stacks provider with two labels (type + name)
provider "aws" "main" {
  config {
    region = "us-east-1"
  }
}

# Another two-label provider configuration
provider "aws" "secondary" {
  config {
    region = "us-west-2"
  }
}

# Two-label provider with for_each
provider "random" "default" {}

# Two-label provider with minimal config
provider "null" "test" {}

variable "test_var" {
  type        = string
  description = "Test variable for provider labels test"
  default     = "test"
}

output "test_output" {
  description = "Test output for provider labels test"
  value       = var.test_var
}

component "test_component" {
  source = "./test-module"
}
