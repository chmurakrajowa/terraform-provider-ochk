data "ochk_virtual_network" "auto-nat" {
    display_name = "bypasstest"
}

data "ochk_auto_nat" "auto-nat" {
  display_name = "bypasstest"
}

/*
resource "ochk_auto_nat" "first-auto-nat" {
    display_name = "${var.test-data-prefix}-nat3"
    virtual_network_id = data.ochk_virtual_network.auto-nat.id
}
*/
