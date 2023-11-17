//data "ochk_security_groups" "security_groups" {

//}

/*
resource "ochk_security_group" "project-test" {
  display_name = "${var.test-data-prefix}-sg8-src"
  project_id = data.ochk_project.project_for_vm.id

  members {
    id = data.ochk_virtual_machine.test.id
    type = "VIRTUAL_MACHINE"
  }
}
*/

/*
resource "ochk_security_group" "project-sg1-src" {
  display_name = "${var.test-data-prefix}-sg1-src"
  project_id = data.ochk_project.project_for_vm.id

  members {
    id = ochk_ip_collection.default.id
    type = "IPCOLLECTION"
  }
}

resource "ochk_security_group" "project-sg1-dst" {
  display_name = "${var.test-data-prefix}-sg1-dst"
  project_id = data.ochk_project.project_for_vm.id

   members {
    id = ochk_ip_collection.default.id
    type = "IPCOLLECTION"
  }
}
*/

