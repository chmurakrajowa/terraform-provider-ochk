resource "ochk_ip_collection" "default" {
  display_name = "${var.test-data-prefix}-ipc-default"
  ip_addresses = [
    "1.1.1.1",
    "1.0.0.1",
    "8.8.8.8"
  ]
}

data "ochk_network" "subtenant-network" {
  name = var.subtenant_network_name
}

resource "ochk_router" "default" {
  display_name = "${var.test-data-prefix}-router"
}

resource "ochk_virtual_network" "default" {
  display_name = "${var.test-data-prefix}-vnet3"
  subtenants = [
    data.ochk_subtenant.subtenant_for_vm.id
  ]
}

resource "ochk_virtual_network" "vnet2" {
  display_name = "${var.test-data-prefix}-vnet4"
  ipam_enabled = true
  subtenants = [
    data.ochk_subtenant.subtenant_for_vm.id
  ]

  router = ochk_router.default.id
  dns_search_suffix = "test.lcl,prod.lcl"
  dns_suffix = "test.lcl"
  primary_dns_address = "192.168.1.6"
  secondary_dns_address = "192.168.1.2"
  primary_wins_address = "192.168.1.3"
  secondary_wins_address = "192.168.1.3"
  subnet_network_cidr = "192.168.200.0/24"
}
