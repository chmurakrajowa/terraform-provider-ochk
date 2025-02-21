# data "ochk_billing_accounts" "accts" {
# }
#
# data "ochk_billing_account" "acct2" {
#     display_name = "Rachunek Platformowy"
# }

#
# resource "ochk_billing_account" "acct4" {
#   display_name = "${var.test-data-prefix}-acct-07j"
#  account_description = "opis"
#   projects {
#     project_id = "566cf9f6-6633-443f-9f3b-24b41c0689f8"
#   }
# }
/*
resource "ochk_billing_account" "acct2" {
  display_name = "${var.test-data-prefix}-acct-06j"

  projects {
    project_id = "397ff706-51a3-4cd6-9615-57378f9c9e70"
  }
}*/

