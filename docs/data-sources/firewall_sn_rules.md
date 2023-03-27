---
page_title: "Firewall S-N Rules Data Source"
---

# Firewall S-N Rules Data Source

Data Source for getting Firewall S-N Rules list by VPC id.

## Example Usage

```hcl
data "ochk_vrf" "vrf" {
  display_name = "T0"
}

data "ochk_vpc" "vpc" {
  display_name = "VPC1"
  vrf_id = data.ochk_vrf.vrf.id
}

data "ochk_firewall_sn_rules" "firewall_sn_rules" {
  vpc_id = data.ochk_vpc.vpc.id
}
```

## Argument Reference

The following arguments are supported:

* `vpc_id` - (Required) VPC id.

## Attribute Reference

The following attributes are exported:
* `firewall_sn_rules` - Firewall S-N rules. Each entry has the following values:
    * **display_name** - Rule name.
    * **firewall_en_rule_id** - Rule id.
    * **project_id** - Project id to which Rule is assigned.



    
 
