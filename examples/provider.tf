terraform {
  required_providers {
    ochk = {
      source = "chmurakrajowa/ochk"
    }
  }
  backend "local" {}
}

provider "ochk" {
  host = var.host
  platform = var.platform
  username = var.username
  password = var.password
  debug_log_file = var.debug_log_file
}
