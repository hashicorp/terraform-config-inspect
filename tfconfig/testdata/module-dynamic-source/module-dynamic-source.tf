variable "child_source" {
  type    = string
  const   = true
  default = "app.terraform.io/example-org/child/aws"
}

locals {
  sibling_source  = "app.terraform.io/example-org/sibling/aws"
  sibling_version = "2.1.0"
}

module "child" {
  source  = var.child_source
  version = "1.0.0"
}

module "sibling" {
  source  = local.sibling_source
  version = local.sibling_version
}

module "static" {
  source  = "app.terraform.io/example-org/static/aws"
  version = "3.0.0"
}
