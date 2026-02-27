# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

output "unique_name_of_output" {
  description = "Description of the purpose of this output"
  type        = string
  value       = component.component_name.some_value
  sensitive   = false
  ephemeral   = false
}
