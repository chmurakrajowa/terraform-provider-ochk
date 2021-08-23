
data "ochk_network" "subtenant-network" {
  name = var.subtenant_network_name
}

resource "ochk_virtual_network" "network_for_vm" {
  display_name = "${var.test-data-prefix}-vnet3"
  subtenants = [
    data.ochk_subtenant.subtenant_for_vm.id,
    ochk_subtenant.subtenant-1.id
  ]
  ipam_enabled = true
  router = data.ochk_router.subtenant-vpc1234.id
  dns_search_suffix = "test.lcl,prod.lcl"
  dns_suffix = "test.lcl"
  primary_dns_address = "192.168.1.6"
  secondary_dns_address = "192.168.1.2"
  primary_wins_address = "192.168.1.3"
  secondary_wins_address = "192.168.1.3"
  subnet_network_cidr = "192.168.200.0/24"
}

/*
resource "ochk_virtual_network" "vnet2" {
  display_name = "${var.test-data-prefix}-vnet2"
  subtenants = [
    data.ochk_subtenant.subtenant_for_vm.id
  ]
}

resource "ochk_virtual_network" "vnet1" {
  display_name = "${var.test-data-prefix}-vnet1"
  subtenants = [
    data.ochk_subtenant.subtenant_for_vm.id
  ]
}
*/
