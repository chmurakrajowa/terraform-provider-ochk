---
page_title: "Subtenant Resource"
---

# Subtenant Resource

Resource for managing projects.

## Example Usage

```hcl
resource "ochk_project" "{{ .ResourceName}}" {
	display_name = "project_name"
    account_id = "63182099-5396-48f5-9474-c8031e08f25a"
    vrf_id = "62182099-5396-48g5-9432-c8031e08f23a"
	description = "short description"
	memory_reserved_size_mb = 30000
	storage_reserved_size_gb = 200
	vcpu_reserved_quantity = 100
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Display name for the project. Updates to this attribute forces recreate.
* `description` - (Optional) Description.
* `vrf_id` - (Required) VRF id.
* `account_id` - (Required) Billing account.
* `limits_enabled` - (Optional) Limits for mamtory, sorage and vcpus enabled (True) or disabled (False). Default: True.
* `memory_reserved_size_mb` - (Required if `limits_enabled` is True) Memory reservation size in megabytes.
* `storage_reserved_size_gb` - (Required if `limits_enabled` is True) Storage reservation size in gigabytes.
* `vcpu_reserved_quantity` - (Required if `limits_enabled` is True) VCPU reservation quantity.
  
## Attribute Reference

No additional attributes are exported.

Attention! The TF user has permissions to the project created at the TF level. If you want to use this project in GUI level, you must add appropriate permissions to other users in GUI.

 
