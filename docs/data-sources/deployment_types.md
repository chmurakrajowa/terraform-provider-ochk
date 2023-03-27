---
page_title: "Deployment Types Data Source"
---

# Deployment Types Data Source

Data Source for getting supported deployment types.

## Example Usage

```hcl
data "ochk_deployment_types" "deployment_types" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `deployment_types` - Deployment types list. Each entry has the following values:
    * **deployment_type** - Deployment type, for example: `ISO`, `OVF`, `TEMPLATE`.
