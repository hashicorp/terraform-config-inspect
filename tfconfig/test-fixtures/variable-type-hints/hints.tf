variable "primitive" {
}

variable "list" {
  type = list
}

variable "map" {
  # quoted value here is a legacy/deprecated form, but supported for compatibility
  # with older configurations.
  type = "map"
}
