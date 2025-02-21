terraform {
  required_providers {
    ochk = {
      source = "chmurakrajowa/ochk"
      version = "3.0"
    }
  }
  backend "local" {
    path="examples/VMWARE/test.tfstate"
    }
}

provider "ochk" {
  host = var.host
  platform = var.platform
  api_key = var.api_key
  debug_log_file = var.debug_log_file
  platform_type = var.platform_type
}
