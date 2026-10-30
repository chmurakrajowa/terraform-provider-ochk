---
page_title: "Subtenant Resource"
---

# Subtenant Resource

Resource for managing public ip address.

## Example Usage

```hcl
resource "ochk_public_ip_address" "{{ .ResourceName}}" {
	display_name = "public_ip_address_name"
	description = "short description"
    public_address_ip = "203.0.113.93"
    public_address_ip_id = "f09ba8de-35a6-46e5-8173-18740589a5c3"
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Display name for the public ip address. Updates to this attribute forces recreate.
* `description` - (Optional) Description.
* `public_address_ip` - (Required) Public ip allocation. Very important!. This parameter must be explicitly specified in the resource.
* `public_address_ip_id` - (Required) Public ip allocation id. Very important!. This parameter must be explicitly specified in the resource.
  
## Attribute Reference

No additional attributes are exported.
