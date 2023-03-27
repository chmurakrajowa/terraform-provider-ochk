---
page_title: "KMS encryption keys Data Source"
---

# KMS encryption keys Data Source

Data Source for getting KMS encryption keys list.

## Example Usage

```hcl
data "ochk_kms_keys" "keys" {
}
```

## Argument Reference

No additional arguments are required.

## Attribute Reference

The following attributes are exported:
* `kms_keys` - KMS encryption keys list. Each entry has the following values:
  * **kms_key_id** - KMS encryption key id.
  * **display_name** - Display name of KMS encryption key.
  * **key_usage** - List of key usages, what actions the key will be applied to: e.g. [`ENCRYPT`, `DECRYPT`].
  * **state** - State of the key, e.g. `Active`.

   
 
