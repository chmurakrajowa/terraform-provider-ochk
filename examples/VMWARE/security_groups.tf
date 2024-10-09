//data "ochk_security_groups" "security_groups" {
//}


//  OPENSTACK
/*
resource "ochk_security_group" "project-test" {
        display_name = "${var.test-data-prefix}-sg8-src8"
        project_id = "3cda830c-b37f-46dc-be54-f649d31bec66"
        members {
            //id = "8f39fa0d-0158-4b4c-8dc9-20be31159220" // id vm
            id = "dcbf922a-a2fc-401f-ac1f-5159c15d4b8b" // inny id  vm
            //id = "9f842807-edef-45ad-ab14-3dc0511d0f3f" // id ip ochk_ip_collection
            type = "VIRTUAL_MACHINE"
            //type = "IPSET"
        }
#         members {
#               id = "8f39fa0d-0158-4b4c-8dc9-20be31159220" // id vm
#               //id = "dcbf922a-a2fc-401f-ac1f-5159c15d4b8b" // inny id  vm
#               //id = "9f842807-edef-45ad-ab14-3dc0511d0f3f" // id ip ochk_ip_collection
#               type = "VIRTUAL_MACHINE"
#               //type = "IPSET"
#         }
}

*/
// VMAWARE

#
# resource "ochk_security_group" "project-test" {
#       display_name = "${var.test-data-prefix}-sg8-src-devel-3"
#       project_id = "8FE2C804-9AF1-49D8-B60F-75E0FEEE2922"
#
#       members {
#             //id = "8f39fa0d-0158-4b4c-8dc9-20be31159220" // id vm
#             id = "71c620fd-41ab-4208-8441-02dedfc3c18f" // inny id  vm
#             //id = "9f842807-edef-45ad-ab14-3dc0511d0f3f" // id ip ochk_ip_collection
#             //type = "VIRTUAL_MACHINE"
#             type = "VIRTUAL_MACHINE"
#
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

