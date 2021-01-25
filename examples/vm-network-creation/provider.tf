terraform {
  required_version = "~> 0.14.0"

  required_providers {
    ochk = {
      source = "chmurakrajowa/ochk"
    }
  }
}

provider "ochk" {
  host = var.host
  tenant = var.tenant
  username = var.username
  password = var.password
  debug_log_file = var.debug_log_file
}
