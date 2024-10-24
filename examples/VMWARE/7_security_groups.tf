//data "ochk_security_groups" "security_groups" {
//}

#
# resource "ochk_security_group" "project-test" {
#       display_name = "${var.test-data-prefix}-sg8-src-devel-3"
#       project_id = "8FE2C804-9AF1-49D8-B60F-75E0FEEE2922"
#
#       members {
#             id = "71c620fd-41ab-4208-8441-02dedfc3c18f" // inny id  vm
#             type = "VIRTUAL_MACHINE"
#       }
#       members {
#                id = "d29c20be-d384-4573-9487-a8a9682c415a" // id vm
#                type = "VIRTUAL_MACHINE"
#         }
#
# }


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

