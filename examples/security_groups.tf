
/*
resource "ochk_security_group" "subtenant-test" {
  display_name = "${var.test-data-prefix}-sg8-src"

  members {
    id = data.ochk_virtual_machine.test.id
    type = "VIRTUAL_MACHINE"
  }
}
*/

resource "ochk_security_group" "subtenant-sg1-src" {
  display_name = "${var.test-data-prefix}-sg1-src"

  members {
    id = ochk_ip_collection.default.id
    type = "IPCOLLECTION"
  }
}

resource "ochk_security_group" "subtenant-sg1-dst" {
  display_name = "${var.test-data-prefix}-sg1-dst"

   members {
    id = ochk_ip_collection.default.id
    type = "IPCOLLECTION"
  }
}


