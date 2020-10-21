provider "foo" {
  alias = "blue"
}

provider "foo" {
  alias = "red"
}

provider "bar" {
}

provider "bar" {
  alias = "yellow"
}

provider "baz" {
}

provider "empty" {
  alias = ""
}
