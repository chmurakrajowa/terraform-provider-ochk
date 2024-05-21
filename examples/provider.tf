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
  api_key = var.api_key
  debug_log_file = var.debug_log_file
}
