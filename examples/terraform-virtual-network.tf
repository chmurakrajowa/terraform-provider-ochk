resource "ochk_virtual_network" "vnet2" {
  display_name = "ls-tf-${var.demo-id}"
  ipam_enabled = true
  subtenants = [ochk_subtenant.bg-1.id]
  router = ochk_router.T1.id
  dns_search_suffix = "test.lcl,prod.lcl"
  dns_suffix = "test.lcl"
  primary_dns_address = "192.168.1.6"
  secondary_dns_address = "192.168.1.2"
  primary_wins_address = "192.168.1.3"
  secondary_wins_address = "192.168.1.3"
  subnet_network_cidr = "10.16.1.0/24"
}
