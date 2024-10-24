//data "ochk_security_groups" "security_groups" {
//}

//data "ochk_security_group" "security_group01" {
//        display_name = "${var.test-data-prefix}-sg8-src8"
//          project_id = "3cda830c-b37f-46dc-be54-f649d31bec66"
//}


resource "ochk_security_group" "project-test" {
        display_name = "${var.test-data-prefix}-sg12"
        project_id = "9a23ad90-3d9f-48d8-b4e2-788db66d90f7"
     //  project_id = "0f00844f-9da4-40f5-860e-98f0e24711a1"

        members {
            id = "7accc7a7-43a2-472a-81c6-6fdb9a35bc2a" // inny id vm
            type = "VIRTUAL_MACHINE"
        }
  //      members {
  //            id = "8f39fa0d-0158-4b4c-8dc9-20be31159220" // id vm
    //          type = "VIRTUAL_MACHINE"
      //  }
}

/*
resource "ochk_security_group" "project-test02" {
        display_name = "${var.test-data-prefix}-sg8-src9"
        project_id = "9a23ad90-3d9f-48d8-b4e2-788db66d90f7"

        members {
            id = "69e02224-dc9a-4963-893a-4309886f43f4" // inny id  vm
            type = "VIRTUAL_MACHINE"
        }
  //      members {
  //            id = "8f39fa0d-0158-4b4c-8dc9-20be31159220" // id vm
    //          type = "VIRTUAL_MACHINE"
      //  }
}

*/