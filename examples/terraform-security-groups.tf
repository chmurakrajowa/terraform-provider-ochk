resource "ochk_security_group" "ipset" {
  display_name = "tf-${var.demo-id}-ipset"

  members {
    id = "48c33c78-532b-489f-8750-a167a501c6d2"
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
    id = data.ochk_virtual_machine.vm1.id
    type = "VIRTUAL_MACHINE"
  }

  lifecycle {
    ignore_changes = [members]
  }
}

resource "ochk_security_group" "vm_ipset" {
  display_name = "tf-${var.demo-id}-vm-ipset"

  members {
    id = data.ochk_virtual_machine.vm2.id
    type = "VIRTUAL_MACHINE"
  }

  members {
    id = "0098bcc2-1b57-462f-a329-d3a935f37f45"
    type = "IPSET"
  }

  lifecycle {
    ignore_changes = [members]
  }
}

