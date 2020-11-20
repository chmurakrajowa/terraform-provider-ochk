data "ochk_security_policy" "default" {
  display_name = "devel"
}

data "ochk_virtual_machine" "vm1" {
  display_name = "devel0000000098"
}

data "ochk_virtual_machine" "vm2" {
  display_name = "devel0000001110"
}