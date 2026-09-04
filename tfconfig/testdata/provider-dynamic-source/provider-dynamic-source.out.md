
# Module `testdata/provider-dynamic-source`

Provider Requirements:
* **both (`var.some_source`):** `var.some_version`
* **only_source (`"app.terraform.io/${var.some_source}"`):** (any version)
* **only_version (`bar/baz`):** `var.some_version`

## Input Variables
* `some_source` (required)
* `some_version` (required)

