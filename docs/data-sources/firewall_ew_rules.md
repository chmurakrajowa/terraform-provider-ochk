---
page_title: "Firewall E-W Rules Data Source"
---

# Firewall E-W Rules Data Source

Data Source for getting Firewall E-W Rules list by VPC id.

## Example Usage

```hcl
data "ochk_vrf" "vrf" {
  display_name = "T0"
}

data "ochk_vpc" "vpc" {
  display_name = "VPC1"
  vrf_id = data.ochk_vrf.vrf.id
}

data "ochk_firewall_ew_rules" "firewall_ew_rules" {
  vpc_id = data.ochk_vpc.vpc.id
}
```

## Argument Reference

The following arguments are supported:

* `vpc_id` - (Required) VPC id.

## Attribute Reference

The following attributes are exported:
* `firewall_ew_rules` - Firewall E-W rules. Each entry has the following values:
    * **display_name** - Rule name.
    * **firewall_ew_rule_id** - Rule id.
    * **project_id** - Project id to which Rule is assigned.



    
 
