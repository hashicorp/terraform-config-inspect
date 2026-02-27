# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

variable "primitive" {
}

variable "list" {
  type = list(string)
}

variable "string_default_empty" {
  type    = string
  default = ""
}

variable "string_default_null" {
  type    = string
  default = null
}

variable "list_default_empty" {
  type    = list(string)
  default = []
}

variable "object_default_empty" {
  type    = object({})
  default = {}
}

variable "number_default_zero" {
  type    = number
  default = 0
}

variable "bool_default_false" {
  type    = bool
  default = false
}

variable "unique_variable_name" {
  description = "Description of the purpose of this variable"
  type        = string
  default     = "Default variable value"
  sensitive   = false
}
