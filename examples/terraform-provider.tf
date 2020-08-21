terraform {
  required_providers {
    ochk = {
      source = "registry.terraform.io/chmurakrajowa/ochk"
    }
  }
}

provider "ochk" {
  host = var.host
  tenant = var.tenant
  username = var.username
  debug_log_file = var.debug_log_file
}