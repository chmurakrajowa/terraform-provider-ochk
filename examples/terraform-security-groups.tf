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

resource "ochk_security_group" "vm" {
  display_name = "tf-${var.demo-id}-vm"

  members {
    id = "fa6457e3-aea9-4ef1-8450-de1ce676b6b9"
    type = "VIRTUAL_MACHINE"
  }

  lifecycle {
    ignore_changes = [members]
  }
}

resource "ochk_security_group" "vm_ipset" {
  display_name = "tf-${var.demo-id}-vm-ipset"

  members {
    id = "41a4a906-b635-4524-9283-fd4997ee7b62"
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

