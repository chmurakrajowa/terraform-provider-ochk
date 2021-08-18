/*
resource "ochk_router" "subtenant-vpc" {
  display_name = "${var.test-data-prefix}-vpc"
  parent_router_id = data.ochk_router.subtenant-vrf.id
}
*/

data "ochk_router" "subtenant-vpc1234" {
  display_name = "1234"
  parent_router_id = data.ochk_router.subtenant-vrf.id
}