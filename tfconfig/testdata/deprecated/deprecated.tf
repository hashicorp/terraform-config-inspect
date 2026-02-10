# Test file with deprecated variables and outputs

variable "old_api_key" {
  type        = string
  description = "API key for the old service"
  deprecated  = "Use var.new_api_key instead, this will be removed in v2.0"
}

variable "legacy_endpoint" {
  description = "Legacy endpoint URL"
  default     = "https://old.example.com"
  deprecated  = "This endpoint is deprecated and will be removed"
}

variable "current_setting" {
  description = "A current setting that is not deprecated"
  type        = string
  default     = "default_value"
}

output "old_result" {
  description = "The old result output"
  value       = var.old_api_key
  deprecated  = "Use output.new_result instead"
}

output "legacy_data" {
  description = "Legacy data output"
  value       = var.legacy_endpoint
  sensitive   = true
  deprecated  = "This output is deprecated, use modern_data"
}

output "current_output" {
  description = "A current output that is not deprecated"
  value       = var.current_setting
}
