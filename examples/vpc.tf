/*data "ochk_vpcs" "vpcs" {
    //vrf_id = data.ochk_vrf.project-vrf.id
}

resource "ochk_vpc" "project-vpc" {
  display_name = "${var.test-data-prefix}-vpc2"
  vrf_id = data.ochk_vrf.project-vrf.id
  project_id = data.ochk_project.project_for_vm.id
}

resource "ochk_vpc" "project-vpc-folder" {
  display_name = "${var.test-data-prefix}-vpc-folder"
  vrf_id = data.ochk_vrf.project-vrf.id
  project_id = data.ochk_project.project_for_vm.id
  folder_path = "/radek"
}

resource "ochk_vpc" "project-vpc-1" {
  display_name = "${var.test-data-prefix}-vpc3"
  vrf_id = data.ochk_vrf.project-vrf.id
  project_id = ochk_project.project1.id
}

resource "ochk_vpc" "project-vpc1" {
  display_name = "${var.test-data-prefix}-vpc4"
  vrf_id = data.ochk_vrf.project-vrf.id
  project_id = data.ochk_project.project_for_vm.id
}*/

/*
data "ochk_router" "project-vpc1234" {
  display_name = "r18-vpc"
  vrf_id = data.ochk_vrf.project-vrf.id
}
*/

/*
resource "ochk_vpc" "project-vpc" {
  display_name = "vpc-testy"
  vrf_id = data.ochk_vrf.project-vrf.id
}
*/