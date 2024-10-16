---
page_title: "Firewall Rule Data Source"
---

# Firewall Rule Data Source

Data Source for getting Firewall Rule by project id, security group id and display name.

## Example Usage

```hcl
data "ochk_project" "proj01" {
display_name = "project_name"
}

data "ochk_security_group" "sg1" {
  display_name = "security_group_name"
  project_id = data.ochk_project.proj01.id
}

data "ochk_firewall_rule" "{{ .DataSourceName}}" {
  project_id = data.ochk_project.proj01.id
  security_group_id = data.ochk_security_group.sg1.id
  display_name = "firewall_rule_name"
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) Project id to which Rule is assigned.
* `security_group_id` - (Required) Security Group id.
* `display_name` - (Required) Rule name.


## Attribute Reference

The following attributes are exported in addition to above arguments:

* `description` - Short description
* `ether_type` - Ether: IPv4, IPv6
* `direction` - Direction name, values: `INGRESS`, `EGRESS`.
* `protocol` - Protocol type, values: `TCP`, `UPD`, `ICMP`.
* `port_range_min` - Min port value.
* `port_range_max` - Max port value.
* `remote_ip_prefix` - IP address (prefix).
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.



    
 
