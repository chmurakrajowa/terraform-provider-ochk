---
page_title: "Firewall E-W Rule Data Source"
---

# Firewall E-W Rule Data Source

Data Source for getting Firewall E-W Rule by VPC id and display name.

## Example Usage

```hcl
data "ochk_vrf" "vrf" {
  display_name = "vrf_name"
}

data "ochk_vpc" "vpc" {
  display_name = "vpc_name"
  vrf_id = data.ochk_vrf.vrf.id
}

data "ochk_firewall_ew_rule" "{{ .DataSourceName}}" {
  vpc_id = data.ochk_vpc.vpc.id
  display_name = "ew_rule_name"
}
```

## Argument Reference

The following arguments are supported:

* `vpc_id` - (Required) VPC id.
* `display_name` - (Required) Rule name.
* `project_id` - (Required) Project id to which Rule is assigned.

## Attribute Reference

The following attributes are exported in addition to above arguments:

* `action` -  Action name, values: `ALLOW`, `DENY`, `DROP`.
* `direction` - Direction name, values: `IN`, `OUT`, `IN_OUT`.
* `disabled` - Boolean value for disabled or enabled.
* `ip_protocol` - IP protocol type, values: `IPV4`, `IPV6`, `IPV4_IPV6`.
* `services` - List of services id.
* `custom_services` - List of custom services id.
* `source` - Source rule.
* `destination` - Destination rule.
* `priority` - Priority value for rule.
* `created_by` - Who created this resource.
* `created_at` - When this resource was created.
* `modified_by` - Who last modified this resource.
* `modified_at` - When last modification occurred.



    
 
