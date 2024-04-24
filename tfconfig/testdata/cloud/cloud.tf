terraform {
  cloud {
    organization = "hashicorp"
    workspaces {
      name = "example"
    }
  }
}
