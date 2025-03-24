---
page_title: "Firewall Rules Data Source"
---

# Firewall Rules Data Source

Data Source for getting Firewall Rules list by project id and security group id.

## Example Usage

```hcl
data "ochk_project" "proj01" {
display_name = "proj_test01"
}

data "ochk_security_group" "sg1" {
display_name = "sg_test01"
project_id = data.ochk_project.proj01.id
}

data "ochk_firewall_rules" "{{ .DataSourceName}}" {
project_id = data.ochk_project.proj01.id
security_group_id = data.ochk_security_group.sg1.id
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Required) Project id to which Rule is assigned.
* `security_group_id` - (Required) Security Group id.

## Attribute Reference

The following attributes are exported:
* `firewall_rules` - Firewall rules. Each entry has the following values:
    * **display_name** - Rule name.
    * **firewall_rule_id** - Rule id.
    * **project_id** - Project id to which Rule is assigned.
    * **security_group_id** - Security Group id.



    
 
