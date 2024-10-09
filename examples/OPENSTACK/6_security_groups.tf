//data "ochk_security_groups" "security_groups" {
//}

//data "ochk_security_group" "security_group01" {
//        display_name = "${var.test-data-prefix}-sg8-src8"
//          project_id = "3cda830c-b37f-46dc-be54-f649d31bec66"
//}

/*
resource "ochk_security_group" "project-test" {
        display_name = "${var.test-data-prefix}-sg8-src8"
        project_id = "3cda830c-b37f-46dc-be54-f649d31bec66"

        members {
            //id = "8f39fa0d-0158-4b4c-8dc9-20be31159220" // id vm
            id = "dcbf922a-a2fc-401f-ac1f-5159c15d4b8b" // inny id  vm
            type = "VIRTUAL_MACHINE"
        }
        members {
              id = "8f39fa0d-0158-4b4c-8dc9-20be31159220" // id vm
              //id = "dcbf922a-a2fc-401f-ac1f-5159c15d4b8b" // inny id  vm
              type = "VIRTUAL_MACHINE"
        }
}
*/
