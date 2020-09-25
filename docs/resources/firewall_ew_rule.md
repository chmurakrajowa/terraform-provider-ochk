---
page_title: "Firewall EW Rule Resource"
---

# Firewall EW Rule Resource

Resource for managing firewall EW Rules. 

## Example Usage

```hcl
data "ochk_security_policy" "default" {
  display_name = "devel"
}

data "ochk_service" "http" {
  display_name = "HTTP"
}

resource "ochk_firewall_ew_rule" "fw-ew2" {
  display_name = "tf-fw-ew-http"
  security_policy_id = data.ochk_security_policy.default.id
  services = [data.ochk_service.http.id]

  source = [ochk_security_group.source.id]
  destination = [ochk_security_group.destination.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  position {
    rule_id = ochk_firewall_ew_rule.other-ew-rule.id
    revise_operation = "AFTER"
  }
}
```

## Argument Reference

The following arguments are supported:

* `security_policy_id` - (Required) Identifier of security policy.
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
  * **rule_id**: (Optional) identifier of other Firewall EW Rule. Required, when `revise_operation` set to `BEFORE` or `AFTER`.    
  
## Attribute Reference

The following attributes are exported in addition to above arguments: 
 * `created_by` - Who created this resource.
 * `created_at` - When this resource was created.
 * `modified_by` - Who last modified this resource. 
 * `modified_at` - When last modification occurred.
 *

 
