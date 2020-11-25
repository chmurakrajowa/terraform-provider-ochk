data "ochk_security_policy" "default" {
  display_name = var.tenant
}

data "ochk_user" "default" {
  name = var.test_user
}
