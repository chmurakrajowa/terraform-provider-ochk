---
page_title: "KMS encryption key Data Source"
---

# KMS encryption key Data Source

Data Source for getting KMS encryption keys display name and version.

## Example Usage

```hcl
data "ochk_kms_key" "key" {
  display_name = "vm_encryption_key"
  version = 1
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of KMS encryption key.
* `version` - (Required) Version of the key.

## Attribute Reference

The following attributes are exported in addition to above arguments:
* `key_usage` - List of key usages, what actions the key will be applied to: e.g. [`ENCRYPT`, `DECRYPT`].
* `algorithm` - Encryption algorithm.
* `size` - The maximum number of combinations required to break the encryption algorithm.
* `activation_date` - When the key was activated.
* `created_at` - When the key was created.
* `default_iv` - Default IV.
* `object_type` - Type of key.
* `revocation_reason` - Reason of revocation if the key was revoked.
* `sha1_fingerprint` - SHA-1 unique identifier of the certificate.
* `sha256_fingerprint` - SHA-256 unique identifier of the certificate.
* `state` - State of the key, e.g. `Active`.

   
 
