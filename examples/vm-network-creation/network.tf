resource "random_string" "random" {
  length = 5
  special = false
}

resource "ochk_router" "default" {
  display_name = "${var.tenant}-router-${random_string.random}"
}

resource "ochk_virtual_network" "default" {
  display_name = "${var.tenant}-vnet-${random_string.random}"
  subtenants = [
    data.ochk_subtenant.subtenant.id
  ]
}

resource "ochk_virtual_network" "vnet2" {
  display_name = "${var.tenant}-vnet-${random_string.random}"
  ipam_enabled = true
  subtenants = [
    data.ochk_subtenant.subtenant.id
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
