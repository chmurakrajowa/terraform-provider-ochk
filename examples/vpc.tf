
resource "ochk_router" "subtenant-vpc" {
  display_name = "${var.test-data-prefix}-vpc"
  parent_router_id = data.ochk_router.subtenant-vrf.id
}


/*
data "ochk_router" "subtenant-vpc1234" {
  display_name = "r18-vpc"
  parent_router_id = data.ochk_router.subtenant-vrf.id
}
*/

/*
resource "ochk_router" "subtenant-vpc" {
  display_name = "vpc-testy"
  parent_router_id = data.ochk_router.subtenant-vrf.id
}
*/