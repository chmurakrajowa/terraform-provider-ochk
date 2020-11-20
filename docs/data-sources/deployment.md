---
page_title: "Deployment Data Source"
---

# Deployment Data Source

Data Source for getting deployment instances by display name. 

## Example Usage

```hcl
data "ochk_deployment" "centos" {
  display_name = "CentOS 7"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of deployment.

## Attribute Reference

The following attributes are exported:
 * `display_name` - Name of deployment. 
    
 
