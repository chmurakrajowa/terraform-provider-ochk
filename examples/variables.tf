variable "test-data-prefix" {}
variable "host" {}
variable "tenant" {}
variable "username" {}
variable "password" {}
variable "debug_log_file" {
  default = ""
}
variable "bg_manager_user" {}
variable "subtenant_network_name" {}
variable "subtenant_for_vm_name" {}
variable "initial_password_for_vm" {}

variable "vrf_router" {}

variable "iso_image" {}
variable "ovf_image" {}
variable "backup_plan" {}
variable "backup_list" {}

variable "billing_tag_cc" {}
variable "system_tag_os" {}