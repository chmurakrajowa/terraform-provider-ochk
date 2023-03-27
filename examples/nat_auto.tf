/*
data "ochk_auto_nat" "auto-nat" {
  display_name = "UITestNAT01"
}
*/

data "ochk_auto_nats" "auto_nats" {

}

/*
resource "ochk_auto_nat" "first-auto-nat" {
    display_name = "${var.test-data-prefix}-nat3"
    virtual_network_id = data.ochk_virtual_network.auto-nat.id
}
*/
