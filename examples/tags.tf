/*
data "ochk_billing_tag" "cc1" {
  display_name = var.billing_tag_cc
}

data "ochk_system_tag" "os1" {
  display_name = var.system_tag_os
}

resource "ochk_billing_tag" "res-bt-cc2" {
    display_name = "${var.test-data-prefix}-billing-tag-cc2"
}

resource "ochk_system_tag" "res-st-os2" {
    display_name = "${var.test-data-prefix}-system-tag-os2"
}
*/