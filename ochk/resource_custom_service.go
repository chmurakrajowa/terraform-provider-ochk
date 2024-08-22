package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	CustomServiceRetryTimeout = 1 * time.Minute
)

func resourceCustomService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCustomServiceCreate,
		ReadContext:   resourceCustomServiceRead,
		UpdateContext: resourceCustomServiceUpdate,
		DeleteContext: resourceCustomServiceDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(CustomServiceRetryTimeout),
			Update: schema.DefaultTimeout(CustomServiceRetryTimeout),
			Delete: schema.DefaultTimeout(CustomServiceRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: resourceCustomServiceImportState,
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
			"ports": {
				Type:     schema.TypeList,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},
						"source": {
							Type:     schema.TypeSet,
							Required: true,
							MinItems: 1,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"destination": {
							Type:     schema.TypeSet,
							Required: true,
							MinItems: 1,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
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

func resourceCustomServiceImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	d.SetId(strings.ToLower(d.Id()))
	return []*schema.ResourceData{d}, nil
}

func resourceCustomServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).CustomServices

	customService := mapResourceDataToCustomService(d)

	created, err := proxy.Create(ctx, customService)
	if err != nil {
		return diag.Errorf("error while creating custom service: %+v", err)
	}

	d.SetId(created.ServiceID.String())

	return resourceCustomServiceRead(ctx, d, meta)
}

func resourceCustomServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).CustomServices

	customService, err := proxy.Read(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("custom service with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading custom service: %+v", err)
	}

	if err := d.Set("display_name", customService.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("project_id", customService.ProjectID); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
	}

	if err := d.Set("ports", flattenCustomServicePorts(customService.L4PortSetEntries)); err != nil {
		return diag.Errorf("error setting members: %+v", err)
	}

	if err := d.Set("created_by", customService.CreatedBy); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", customService.CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", customService.ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", customService.ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	return nil
}

func resourceCustomServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).CustomServices

	customService := mapResourceDataToCustomService(d)
	customService.ServiceID = strfmt.UUID(d.Id())

	_, err := proxy.Update(ctx, customService)
	if err != nil {
		return diag.Errorf("error while modifying custom service: %+v", err)
	}

	return resourceCustomServiceRead(ctx, d, meta)
}

func resourceCustomServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).CustomServices

	err := proxy.Delete(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("custom service with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting custom service: %+v", err)
	}

	return nil
}

func mapResourceDataToCustomService(d *schema.ResourceData) *models.CustomServiceInstance {
	return &models.CustomServiceInstance{
		DisplayName:      d.Get("display_name").(string),
		ProjectID:        strfmt.UUID(d.Get("project_id").(string)),
		L4PortSetEntries: expandCustomServicePorts(d.Get("ports").([]interface{})),
	}
}
