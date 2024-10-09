variable "test-data-prefix" {}
variable "host" {}
variable "platform" {}
variable "api_key" {}
variable "platform_type" {}
variable "debug_log_file" {
  default = ""
}
variable "project_network_name" {}
variable "project_for_vm_name" {}
variable "initial_password_for_vm" {}

variable "vrf_router" {}

variable "iso_image" {}
variable "ovf_image" {}
variable "backup_plan" {}
variable "backup_list" {}

variable "billing_tag_cc" {}
variable "system_tag_os" {}