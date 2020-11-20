resource "ochk_security_group" "ipcoll" {
  display_name = "tf-${var.demo-id}-ipcoll-ranges"

  members {
    id = ochk_ip_collection.ranges.id
    type = "IPCOLLECTION"
  }

  lifecycle {
    ignore_changes = [members]
  }
}

resource "ochk_security_group" "vm" {
  display_name = "tf-${var.demo-id}-vm"

  members {
    id = data.ochk_virtual_machine.vm1.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_security_group" "vm_ipc" {
  display_name = "tf-${var.demo-id}-vm-ipcoll-local24"

  members {
    id = data.ochk_virtual_machine.vm2.id
    type = "VIRTUAL_MACHINE"
  }

  members {
    id = ochk_ip_collection.local24.id
    type = "IPCOLLECTION"
  }
}

resource "ochk_security_group" "vm-with-ip-collections" {
  display_name = "tf-${var.demo-id}-vm-ipcoll-dns-servers"

  members {
    id = data.ochk_virtual_machine.vm2.id
    type = "VIRTUAL_MACHINE"
  }

  members {
    id = ochk_ip_collection.dns-servers.id
    type = "IPCOLLECTION"
  }
}


