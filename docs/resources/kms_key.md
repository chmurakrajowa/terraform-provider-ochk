---
page_title: "KMS encryption key resource"
---

# KMS encryption key Resource

Resource for managing KMS encryption keys. Currently, `ochk_kms_key` resource supports only creating, importing and deleting keys.

Created/uploaded key always has version set to 0.

It is not possible to update a key or create new key version. Any updates will recreate key resource.
  
## Example Usage

```hcl
resource "ochk_kms_key" "rsa" {
  display_name = "rsa-key"
  key_usage = ["ENCRYPT", "DECRYPT"]
  algorithm = "RSA"
  size = 2048
}
```

Creates a new RSA key with version = 0. 

```hcl
resource "ochk_kms_key" "aes-uploaded" {
  display_name = "aes-uploaded"
  key_usage = ["ENCRYPT", "DECRYPT"]
  algorithm = "AES"
  size = 256
  material = "Er8BJNFpGPIfeBDZuw[...]ywGjvc/1e3k7F/6DVg=="
}
```

Upload a new AES key with version = 0.

```hcl
resource "ochk_kms_key" "rsa-uploaded" {
  display_name = "rsa-uploaded"
  key_usage = ["ENCRYPT", "DECRYPT"]
  algorithm = "RSA"
  size = 2048
  material = <<EOT
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA8vXmbi9IAPnkIQTwTijtX3tdNRokMJ7WA2pwMOswEihck5n3
f+tbtBTGwWuat2UiPUgWfDtNiqHeK8CTWZuGFyITxF8n2iiQM9CTcR8AJUkp08Fn
[...]
0yG2J8e9jUhdd70nTWsAQikQaaszrZLdJ/Hk1zm0VIU+ohBaWmg02m4=
-----END RSA PRIVATE KEY-----
EOT
}

resource "ochk_kms_key" "aes-enc-uploaded" {
  display_name = "aes-encrypted"
  key_usage = ["ENCRYPT", "DECRYPT"]
  algorithm = "AES"
  size = 256
  material = "IgQEzbAU/oCTkXyrRZ6zKEna3aNQv+[...]dhrXbX/yA6JIw=="
  private_key_id_to_unwrap = ochk_kms_key.rsa-uploaded.id
}

```

Upload a new RSA key encrypted with other symmetric key.

## Argument Reference

The following arguments are supported:

* `display_name` - (Required) Exact display name of KMS encryption key.
* `key_usage` - (Required) List of key usages, what actions the key will be applied to: e.g. [`ENCRYPT`, `DECRYPT`].
* `algorithm` - (Required) Encryption algorithm, values: `AES`, `RSA`.
* `size` - (Required) Size of the key in bits.
* `private_key_id_to_unwrap` - (Optional) Id of the RSA private key that we used to secure the cryptographic material. Use `ochk_kms_key` data source to get key id.
* `material` - (Optional) Uploaded cryptographic material, the format depends on the selected key type. AES keys are encoded in Base64, RSA keys are in PEM format.
* `version` - (Required) Version of the key.

## Attribute Reference

The following attributes are exported in addition to above arguments:
* `activation_date` - When the key was activated.
* `created_at` - When the key was created.
* `default_iv` - Default IV.
* `object_type` - Type of key.
* `revocation_reason` - Reason of revocation if the key was revoked.
* `sha1_fingerprint` - SHA-1 unique identifier of the certificate.
* `sha256_fingerprint` - SHA-256 unique identifier of the certificate.
* `state` - State of the key, e.g. `Active`.

   
 
