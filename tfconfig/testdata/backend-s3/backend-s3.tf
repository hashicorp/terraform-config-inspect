terraform {
  backend "s3" {
    bucket = "terraform-state-bucket"
    key    = "network/terraform.tfstate"
    region = "us-west-2"
  }
}
