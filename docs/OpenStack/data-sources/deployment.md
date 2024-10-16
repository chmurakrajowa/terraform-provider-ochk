---
page_title: "Deployment Data Source"
---

# Deployment Data Source

Data Source for getting deployment instances by display name. 

## Example Usage

```hcl
data "ochk_deployment" "{{ .DataSourceName}}" {
  display_name = "Debian 12"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of deployment.

## Attribute Reference

The following attributes are exported:
 * `display_name` - Name of deployment. For ISO and OVF deployment name is image file name.
 * `deployment_type` - Type of deployment. You can get list of deployment types from [Deployment Types](deployment_types.md) Data Source.
 * `initial_size_mb` - Size of first partition in deployment in megabytes.
    
 
