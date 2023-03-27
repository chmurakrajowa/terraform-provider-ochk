data "ochk_virtual_networks" "virtual_networks" {

}

resource "ochk_virtual_network" "network_for_vm" {
  display_name = "${var.test-data-prefix}-vnet3"
  ipam_enabled = true
  vpc_id = ochk_vpc.project-vpc.id
  project_id = data.ochk_project.project_for_vm.id
  subnet_network_cidr = "192.168.51.0/24"
}

/*
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
  subnet_network_cidr = "192.168.51.0/24"
}
*/
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
