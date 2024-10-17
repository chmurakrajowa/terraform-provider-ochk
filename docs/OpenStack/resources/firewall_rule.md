---
page_title: "Firewall Rule Resource"
---

# Firewall Rule Resource

Resource for managing firewall rules to control incoming and outgoing network traffic.

## Example Usage

```hcl
data "ochk_project" "project" {
  display_name = "project_name"
}

data "ochk_security_group" "sg" {
  display_name = "security_group_name"
}

resource "ochk_firewall_rule" "{{ .ResourceName}}" {
    display_name = "firewall_rule_name"
    project_id = data.ochk_project.project.id
    security_group_id = data.ochk_security_group.sg.id
    description = "short description"
    ether_type = "IPv4"
    direction = "EGRESS"
    protocol = "TCP"
    port_range_min = 20008
    port_range_max = 20020
    remote_ip_prefix = ""
}

```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) Project id to which Rule is assigned.
* `security_group_id` - (Required) Security Group id.
* `display_name` - (Required) The Firewall Rule name.
* `description` - Short description
* `ether_type` - Ether: IPv4, IPv6
* `direction` - Direction name, values: `INGRESS`, `EGRESS`.
* `protocol` - Protocol type, values: `TCP`, `UPD`, `ICMP`.
* `port_range_min` - Min port value.
* `port_range_max` - Max port value.
* `remote_ip_prefix` - IP address (prefix).

## Attribute Reference

The following attributes are exported in addition to above arguments: 
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred.

 
