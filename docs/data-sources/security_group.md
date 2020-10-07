---
page_title: "Security Group Data Source"
---

# Security Group Data Source

Data Source for reading security groups by display name. 

## Example Usage

```hcl
data "ochk_security_group" "sg" {
  display_name = "security-group-name"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of security group.

## Attribute Reference

The following attributes are exported:
 * `display_name` - Security group name. 
 * `members` - Members that are assigned to the security group.
   Each entry has following values:
    * **type**: type of security group member, values: IPCOLLECTION, IPSET, VIRTUAL_MACHINE, LOGICAL_PORT
    * **id**: resource identifier depending on the security group type selection
 
