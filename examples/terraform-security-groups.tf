resource "ochk_security_group" "ipset" {
  display_name = "tf-${var.demo-id}-ipset"

  members {
    id = data.ochk_ip_set.ipset1.id
    type = "IPSET"
  }

  lifecycle {
    ignore_changes = [members]
  }
}

data "ochk_virtual_machine" "vm1" {
  display_name = "devel0000000343"
}

data "ochk_virtual_machine" "vm2" {
  display_name = "devel0000000350"
}

resource "ochk_security_group" "vm" {
  display_name = "tf-${var.demo-id}-vm"

  members {
    id = "fa6457e3-aea9-4ef1-8450-de1ce676b6b9"
//    id = data.ochk_virtual_machine.vm1.id
    type = "VIRTUAL_MACHINE"
  }

  lifecycle {
    ignore_changes = [members]
  }
}

resource "ochk_security_group" "vm_ipset" {
  display_name = "tf-${var.demo-id}-vm-ipset"

  members {
    id = "c1b66861-0d52-4b22-950e-15a99a6d546c"
//    id = data.ochk_virtual_machine.vm2.id
    type = "VIRTUAL_MACHINE"
  }

  members {
    id = data.ochk_ip_set.ipset2.id
    type = "IPSET"
  }

  lifecycle {
    ignore_changes = [members]
  }
}

