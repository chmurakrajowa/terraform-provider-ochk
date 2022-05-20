---
page_title: "NAT Source"
---

# NAT Data Source

Data Source for creating NAT

## Example Usage


## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name for NAT.
* `virtual_network_id` - (Required) Exact virtual network id.
* `vrf_id` - (Optional) Exact vrf id.
* `nat_type` - (Optional) Exact nat type.
* `action` - (Optional) Exact action.
* `priority` - (Optional) Exact priority.
* `service_id` - (Optional) Exact service id.
* `source_network` - (Optional) Exact source network.
* `translated_network` - (Optional) Exact translated network.
* `destination_network` - (Optional) Exact destination network.



## Attribute Reference

The following attributes are exported in addition to above arguments:
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.

