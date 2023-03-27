---
page_title: "Firewall SN Rule Resource"
---

# Firewall SN Rule Resource

Resource for managing firewall SN Rules to control incoming and outgoing network traffic between virtual machines. The rules can allow traffic (ALLOW rule), cut a connection (DROP rule) or drop the connection (REJECT rule).

## Example Usage

```hcl
data "ochk_vrf" "vrf" {
  display_name = "T0"
}

data "ochk_project" "project" {
  display_name = "example_project"
}

data "ochk_vpc" "vpc" {
  display_name = "VPC1"
  vrf_id = data.ochk_vrf.vrf.id
}

data "ochk_service" "ssh" {
  display_name = "SSH"
}

data "ochk_custom_service" "web-servers" {
  display_name = "web-servers"
  project_id = data.ochk.project.project.id
}

data "ochk_security_group" "source" {
  display_name = "security-group-name-src"
}

data "ochk_security_group" "destination" {
  display_name = "security-group-name-dst"
}

resource "ochk_firewall_sn_rule" "fw-sn-1" {
  display_name = "fw-sn-drop-ssh"
  vpc_id = data.ochk_vpc.vpc.id
  project_id = data.ochk.project.project.id
  scope = [data.ochk_router.T1.id]

  services = [data.ochk_service.ssh.id]
  custom_services = [data.ochk_custom_service.web-servers.id]

  source = [data.ochk_security_group.source.id]
  destination = [data.ochk_security_group.destination.id]

  action = "ALLOW"
  ip_protocol = "IPV4"

  priority = 1000
}
```

## Argument Reference

The following arguments are supported:

* `vpc_id` - (Required) Identifier of vpc id.
* `display_name` - (Required) The Firewall SN Rule name.
* `project_id` - (Required) Project id to which Rule is assigned.
* `action` - (Optional) Action to control the traffic between the source and the target. It is possible to open the traffic between the source and target with the ALLOW rule, cut the traffic between the source and target with the DROP rule, and reject the connection between the source and target with the REJECT rule. Allowed values: `ALLOW`, `DROP`, `REJECT`. Default value: `ALLOW`.
* `direction` - (Optional) The traffic direction that the firewall rule applies to. Allowed values: `IN`, `IN_OUT`, `OUT`. Default value: `IN_OUT`.
* `disabled` - (Optional) Sets this rule to be disabled. Default: false
* `ip_protocol` - (Optional) The transport protocol used for the connection. Allowed values: `IPV4`, `IPV6`, `IPV4_IPV6`. Default value: `IPV4_IPV6`.
* `services` - (Optional) Identifier of the type of traffic to which a firewall rule applies. Use `ochk_service` data source for finding service id.
* `custom_services` - (Optional) Identifier of the type of traffic to which a firewall rule applies. Use `ochk_custom_service` data source for finding custom service id. 
* `source` - (Optional) Identifier of source. The source in a rule can be a previously created security group. Use `ochk_security_group` data source for finding security group id. One of source or destination identifiers are required. 
* `destination` - (Optional) Identifier of destination that will to be used as match criteria for outgoing traffic. The destination in a rule can be a previously created security group. Use ochk_security_group data source for finding security group id. One of source or destination identifiers are required. 
* `priority` - (Required) Priority of the firewall rule. Rules with lower priority are matched first.

## Attribute Reference

The following attributes are exported in addition to above arguments: 
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred.     

 
