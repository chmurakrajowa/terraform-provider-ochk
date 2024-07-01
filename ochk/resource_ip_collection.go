package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	IPCollectionRetryTimeout = 1 * time.Minute
)

func resourceIPCollection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIPCollectionCreate,
		ReadContext:   resourceIPCollectionRead,
		UpdateContext: resourceIPCollectionUpdate,
		DeleteContext: resourceIPCollectionDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(IPCollectionRetryTimeout),
			Update: schema.DefaultTimeout(IPCollectionRetryTimeout),
			Delete: schema.DefaultTimeout(IPCollectionRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: resourceIPCollectionImportState,
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_addresses": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIPCollectionImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	d.SetId(strings.ToLower(d.Id()))
	return []*schema.ResourceData{d}, nil
}

func resourceIPCollectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).IPCollections

	ipCollection := mapResourceDataToIPCollection(d)

	created, err := proxy.Create(ctx, ipCollection)
	if err != nil {
		return diag.Errorf("error while creating ip collection: %+v", err)
	}

	d.SetId(created.ID)

	return resourceIPCollectionRead(ctx, d, meta)
}

func resourceIPCollectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).IPCollections

	ipCollection, err := proxy.Read(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("ip collection with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading ip collection: %+v", err)
	}

	if err := d.Set("ip_addresses", flattenStringSlice(ipCollection.IPCollectionAddresses)); err != nil {
		return diag.Errorf("error setting ip_addresses: %+v", err)
	}

	if err := d.Set("display_name", ipCollection.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("project_id", ipCollection.ProjectID); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
	}

	if err := d.Set("created_by", ipCollection.CreatedBy); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", ipCollection.CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", ipCollection.ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", ipCollection.ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	return nil
}

func resourceIPCollectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).IPCollections

	IPCollection := mapResourceDataToIPCollection(d)
	IPCollection.ID = d.Id()

	_, err := proxy.Update(ctx, IPCollection)
	if err != nil {
		return diag.Errorf("error while modifying ip collection: %+v", err)
	}

	return resourceIPCollectionRead(ctx, d, meta)
}

func resourceIPCollectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).IPCollections

	err := proxy.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("ip collection with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting ip collection: %+v", err)
	}

	return nil
}

func mapResourceDataToIPCollection(d *schema.ResourceData) *models.IPCollection {
	return &models.IPCollection{
		DisplayName:           d.Get("display_name").(string),
		ProjectID:             d.Get("project_id").(string),
		IPCollectionAddresses: transformSetToStringSlice(d.Get("ip_addresses").(*schema.Set)),
	}
}
