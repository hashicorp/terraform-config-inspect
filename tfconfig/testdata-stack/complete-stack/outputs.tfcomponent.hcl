# Copyright IBM Corp. 2018, 2026
# SPDX-License-Identifier: MPL-2.0

output "lambda_urls" {
  type = list(string)
  description = "URLs to invoke lambda functions"
  value = [ for x in component.lambda: x.invoke_arn ]
}
