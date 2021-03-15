package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	KMSKeyRetryTimeout = 5 * time.Minute
)

func resourceKMSKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceKMSKeyCreate,
		ReadContext:   resourceKMSKeyRead,
		DeleteContext: resourceKMSKeyDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(KMSKeyRetryTimeout),
			Delete: schema.DefaultTimeout(KMSKeyRetryTimeout),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_usage": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"algorithm": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"private_key_id_to_unwrap": {
				Type:      schema.TypeString,
				Sensitive: true,
				Optional:  true,
				ForceNew:  true,
				// Setting private_key_id_to_unwrap is possible only on create, subsequent read gets empty string,
				// which causes config drift. This suppresses any reported differences.
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
			},
			"material": {
				Type:      schema.TypeString,
				Sensitive: true,
				ForceNew:  true,
				Optional:  true,
				// Setting material is possible only on create, subsequent read gets empty string,
				// which causes config drift. This suppresses any reported differences.
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
			},
			"activation_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_iv": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"revocation_reason": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sha1_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sha256_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceKMSKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*sdk.Client)

	if material, ok := d.GetOk("material"); ok && material.(string) != "" {
		return resourceKMSKeyImport(ctx, d, meta)
	}

	kmsKey := mapResourceDataToKeyInstance(d)

	keyInstance, err := client.KMSKeys.Create(ctx, kmsKey)
	if err != nil {
		return diag.Errorf("error while creating KMS key: %+v", err)
	}

	d.SetId(keyInstance.ID)

	return resourceKMSKeyRead(ctx, d, meta)
}

func resourceKMSKeyImport(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*sdk.Client)

	kmsKeyImport := mapResourceDataToKeyImport(d)

	keyInstance, err := client.KMSKeys.Import(ctx, kmsKeyImport)
	if err != nil {
		return diag.Errorf("error while importing KMS key: %+v", err)
	}

	d.SetId(keyInstance.ID)

	return resourceKMSKeyRead(ctx, d, meta)
}

func resourceKMSKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).KMSKeys

	kmsKey, err := proxy.Read(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("KMS key with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading KMS key: %+v", err)
	}

	if err := mapKMSKeyToResourceData(d, kmsKey); err != nil {
		return diag.Errorf("error setting KMS key: %v", err)
	}

	return nil
}

func resourceKMSKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkClient := meta.(*sdk.Client)

	err := sdkClient.KMSKeys.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("KMS key with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting KMS key: %+v", err)
	}

	return nil
}

func mapKMSKeyToResourceData(d *schema.ResourceData, kmsKey *models.KeyInstance) error {
	if err := d.Set("display_name", kmsKey.Name); err != nil {
		return fmt.Errorf("error setting display_name: %w", err)
	}
	if err := d.Set("key_usage", flattenStringSlice(kmsKey.KeyUsageList)); err != nil {
		return fmt.Errorf("error setting key_usage: %w", err)
	}
	if err := d.Set("algorithm", kmsKey.Algorithm); err != nil {
		return fmt.Errorf("error setting algorithm: %w", err)
	}
	if err := d.Set("size", kmsKey.Size); err != nil {
		return fmt.Errorf("error setting size: %w", err)
	}
	if err := d.Set("activation_date", kmsKey.ActivationDate.String()); err != nil {
		return fmt.Errorf("error setting activation_date: %w", err)
	}
	if err := d.Set("created_at", kmsKey.CreatedAt.String()); err != nil {
		return fmt.Errorf("error setting created_at: %w", err)
	}
	if err := d.Set("default_iv", kmsKey.DefaultIV); err != nil {
		return fmt.Errorf("error setting default_iv: %w", err)
	}
	if err := d.Set("object_type", kmsKey.ObjectType); err != nil {
		return fmt.Errorf("error setting object_type: %w", err)
	}
	if err := d.Set("revocation_reason", kmsKey.RevocationReason); err != nil {
		return fmt.Errorf("error setting revocation_reason: %w", err)
	}
	if err := d.Set("sha1_fingerprint", kmsKey.Sha1Fingerprint); err != nil {
		return fmt.Errorf("error setting sha1_fingerprint: %w", err)
	}
	if err := d.Set("sha256_fingerprint", kmsKey.Sha256Fingerprint); err != nil {
		return fmt.Errorf("error setting sha256_fingerprint: %w", err)
	}
	if err := d.Set("state", kmsKey.State); err != nil {
		return fmt.Errorf("error setting state: %w", err)
	}

	return nil
}

func mapResourceDataToKeyInstance(d *schema.ResourceData) *models.KeyInstance {
	keyInstance := models.KeyInstance{
		ID:           d.Id(),
		Algorithm:    d.Get("algorithm").(string),
		KeyUsageList: transformSetToStringSlice(d.Get("key_usage").(*schema.Set)),
		Name:         d.Get("display_name").(string),
		Size:         int32(d.Get("size").(int)),
	}

	return &keyInstance
}

func mapResourceDataToKeyImport(d *schema.ResourceData) *models.KeyImport {
	keyInstance := models.KeyImport{
		Algorithm:    d.Get("algorithm").(string),
		KeyName:      d.Get("display_name").(string),
		KeyUsageList: transformSetToStringSlice(d.Get("key_usage").(*schema.Set)),
		Size:         int32(d.Get("size").(int)),
	}

	if material, ok := d.GetOk("material"); ok && material.(string) != "" {
		keyInstance.Material = material.(string)
	}

	if privateKeyIDToUnwrap, ok := d.GetOk("private_key_id_to_unwrap"); ok && privateKeyIDToUnwrap.(string) != "" {
		keyInstance.Material = privateKeyIDToUnwrap.(string)
	}

	return &keyInstance
}
