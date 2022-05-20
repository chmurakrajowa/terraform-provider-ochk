---
page_title: "Firewall E-W Rule Data Source"
---

# Firewall E-W Rule Data Source

Data Source for creating Firewall E-W Rule.

## Example Usage


## Argument Reference

The following arguments are supported:

* `router_id` - (Required) Exact router id.
* `display_name` - (Required) Exact router name.
* `action` - (Optional) Exact action name.
* `direction` - (Optional) Put direction name.
* `disabled` - (Optional) Put boolean value for disable or enable.
* `ip_protocol` - (Optional) Exact ip_protocol value
* `custom_services` - (Optional) Exact id custom services
* `source` - (Optional) Exact source rule
* `destination` - (Optional) Exact destination rule
* `priority` - (Optional) Put priority value for rule


The following attributes are exported in addition to above arguments:
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.



    
 
