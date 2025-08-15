component "compute" {
  source  = "hashicorp/vault-starter/aws"
  version = "~> 4.0"

  inputs = {
    public_ssh_key_url = var.public_ssh_key_url
  }

  providers = {
    aws  = provider.aws.this
    http = provider.http.this
  }
}

component "repositories" {
  source  = "app.terraform.io/hashicorp/nomad-starter/aws"
  version = "~> 1.0"

  inputs = {
  }

  providers = {
    github = provider.aws.this
  }
}
