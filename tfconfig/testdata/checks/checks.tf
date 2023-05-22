
data "external" "foo" {}

check "external" {
  assert {
    condition = data.external.foo.id == "foo"
    error_message = "invalid id"
  }
}

check "internal" {
  data "external" "bar" {}

  assert {
    condition = length(data.external.bar.id) > 0
    error_message = "id too short"
  }

  assert {
    condition = length(data.external.bar.id) < 10
    error_message = "id too long"
  }
}
