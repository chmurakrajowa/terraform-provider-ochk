/*data "ochk_vpc" "project-vpc" {

   vrf_id = "5df9b4f3-8f50-4a98-95d5-8ef2304f4362"
  display_name = "5665655665dfgh"
}*/


/*data "ochk_vpcs" "vpcs" {
    //vrf_id = data.ochk_vrf.project-vrf.id
}

/*data "ochk_vpcs" "vpcs" {
    //vrf_id = data.ochk_vrf.project-vrf.id
    //display_name = "${var.test-data-prefix}_vpc_test01"
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
resource "ochk_vpc" "project-vpc" {
  display_name = "vpc-testy"
  vrf_id = data.ochk_vrf.project-vrf.id
}
*/