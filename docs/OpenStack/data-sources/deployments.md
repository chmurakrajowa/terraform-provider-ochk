---
page_title: "Deployments Data Source"
---

# Deployments Data Source

Data Source for getting deployment instances list.

## Example Usage

```hcl
data "ochk_deployments" "{{ .DataSourceName}}" {
  deployment_type = "TEMPLATE"
}
```

## Argument Reference

The following arguments are supported:

* `"deployment_type` - (Optional) Deployment type. You can get deployment types list from [Deployment types](deployment_types.md) Data Source.

## Attribute Reference

The following attributes are exported:
* `deployments` - Deployments list. Each entry has the following values:
  * **display_name** - Name of deployment. For `ISO` na `OVF` deployment name is image file name.
  * **deployment_type** - Type of deployment. You can get list of deployment types from [Deployment Types](deployment_types.md) Data Source.
  * **initial_size_mb** - Size of first partition in deployment in megabytes.
  * **deployment_id** - Id of deployment.
    
 
