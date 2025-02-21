---
page_title: "Deployment Params Types Data Source"
---

# Deployment Params Types Data Source

Data Source for getting supported deployment params types.

## Example Usage

```hcl
data "ochk_deployment_params_types" "{{ .DataSourceName}}" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `deployment_param_types` - Deployment params types list. Each entry has the following values:
    * **deployment_param_type** - Deployment params type. For example: `NET_IP_ADDR_01`, `NET_DNS_PRIMARY`, `LOGIN_USERNAME`, `LOGIN_PASSWORD`.

