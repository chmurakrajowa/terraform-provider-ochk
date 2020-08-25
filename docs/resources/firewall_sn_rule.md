---
page_title: "Firewall SN Rule Resource"
---

# Firewall SN Rule Resource

Resource for managing firewall SN Rules. 

## Example Usage

```hcl
data "ochk_router" "T1" {
    display_name = "T1"
}

data "ochk_gateway_policy" "T1" {
  display_name = ochk_router.T1.display_name
}

data "ochk_service" "ssh" {
  display_name = "SSH"
}

resource "ochk_firewall_sn_rule" "fw-sn-1" {
  display_name = "fw-sn-drop-ssh"
  gateway_policy_id = data.ochk_gateway_policy.T1.id
  scope = [data.ochk_router.T1.id]

  services = [data.ochk_service.ssh.id]

  source = [ochk_security_group.source.id]
  destination = [ochk_security_group.destination.id]

  action = "ALLOW"
  ip_protocol = "IPV4"
}
```

## Argument Reference

The following arguments are supported:

* `gateway_policy_id` - (Required) Identifier of gateway policy referencing T1 router. Use `ochk_gateway_policy` for finding gateway policy id. Updates to this attribute forces recreate.
* `scope` - (Required) Scope of this rule. Currently, this needs to be set to id of T1 router referenced in `gateway_policy_id`. Use `ochk_router` with display name the same as in `gateway_policy_id` for finding router identifier for scope. Updates to this attribute forces recreate. 
* `display_name` - (Required) Display name.
* `action` - (Optional) Action to make when this rule is satisfied, allowed values: `ALLOW`, `DROP`, `REJECT`. Default value: `ALLOW`.
* `direction` - (Optional) Direction, allowed values: `IN`, `IN_OUT`, `OUT`. Default value: `IN_OUT`.
* `disabled` - (Optional) Sets this rule to be disabled. Default: false
* `ip_protocol` - (Optional) IP protocol, allowed values: `IPV4`, `IPV6`, `IPV4_IPV6`. Default value: `IPV4_IPV6`.
* `services` - (Optional) Identifier of service. Use `ochk_service` data source for finding service id. 
* `source` - (Optional) Identifier of source security group. Use ochk_security_group data source for finding security group id. One of source or destination identifiers are required. 
* `destination` - (Optional) Identifier of destination security group. Use ochk_security_group data source for finding security group id. One of source or destination identifiers are required. 
* `position` - (Optional) Ordering of this rule with respect to other rules. Updates to this attribute forces recreate. 
  * **revise_operation**: rule placement on the list, allowed values: `BEFORE`, `AFTER`, `TOP`, `BOTTOM`.
  * **rule_id**: (Optional) identifier of other Firewall SN Rule. Required, when `revise_operation` set to `BEFORE` or `AFTER`.   
  
## Attribute Reference

The following attributes are exported in addition to above arguments: 
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred.     

 
