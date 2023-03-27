/*
data "ochk_router" "project-vrf" {
  display_name = var.vrf_router
}
*/

data "ochk_vrf" "project-vrf" {
  display_name = var.vrf_router
}

data "ochk_vrfs" "vrfs" {

}