
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
