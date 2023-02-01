terraform {
  required_providers {
    foo = {
      source = "abc/foo"
      version = "~> 2.0.0"
    }
    bat = {
      source = "abc/bat"
      version = "1.0.0"
    }
  }
}

terraform {
  required_providers {
    foo = {
      version = "~> 2.1.0"
    }
    bat = {
      source  = "baz/bat"
    }
  }
}
