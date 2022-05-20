---
page_title: "Security Group Resource"
---

# Security Group Resource

Resource for managing security groups to provide a modular way to define and compose firewall rules. Security groups can contains virtual machines, ip collections, logical ports which can be assigned as the source or target in the firewall rule to control incoming and outgoing network traffic.

## Example Usage

```hcl
resource "ochk_security_group" "vm_ipcollection" {
  display_name = "tf-sg-vm-ipcollection"

  members {
    id = "41a4a906-b635-4524-9283-fd4997ee7b62"
    type = "VIRTUAL_MACHINE"
  }

  members {
    id = "0098bcc2-1b57-462f-a329-d3a935f37f45"
    type = "IPCOLLECTION"
  }
}
```

It is not recommended to directly use resource identifiers. To avoid that prefer to use data sources:
```hcl
data "ochk_virtual_machine" "vm" {
  display_name = "devel0000001157"
}

data "ochk_ip_set" "google" {
  display_name = "googledns"
}

resource "ochk_security_group" "google_dns_sg" {
  display_name = "tf-sg-google-dns"

  members {
    id = data.ochk_virtual_machine.vm.id
    type = "VIRTUAL_MACHINE"
  }

  members {
    id = data.ochk_ip_collection.google.id
    type = "IPCOLLECTION"
  }
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Security group name.
* `members` - (Required) Members that are assigned to the same security group. 
  Each entry must have the following values:

  * **id**: resource identifier depending on the security group type selection. Use data sources `ochk_ip_collection`, `ochk_virtual_machine`, `ochk_logical_port` for getting id by name.
  * **display_name** (Optional) display name of member of security group
  * **type**: type of security group member. Allowed values: `IPCOLLECTION`, `VIRTUAL_MACHINE`, `LOGICAL_PORT`.

## Attribute Reference

The following attributes are exported in addition to above arguments:
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.  