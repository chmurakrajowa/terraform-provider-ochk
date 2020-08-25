---
page_title: "Security Group Resource"
---

# Security Group Resource

Resource for managing security groups. 

## Example Usage

```hcl
resource "ochk_security_group" "vm_ipset" {
  display_name = "tf-sg-vm-ipset"

  members {
    id = "41a4a906-b635-4524-9283-fd4997ee7b62"
    type = "VIRTUAL_MACHINE"
  }

  members {
    id = "0098bcc2-1b57-462f-a329-d3a935f37f45"
    type = "IPSET"
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
    id = data.ochk_ip_set.google.id
    type = "IPSET"
  }
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Display name.
* `members` - (Required) Members of security group. 
  Each entry must have the following values:
  
  * **type**: type of security group member, allowed values: IPSET, VIRTUAL_MACHINE, LOGICAL_PORT.
  * **id**: resource identifier
  
## Attribute Reference

No additional attributes are exported. 
