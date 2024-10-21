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
	SecurityGroupRetryTimeout            = 1 * time.Minute
	IPCOLLECTION              MemberType = "IPCOLLECTION"
	LOGICAL_PORT              MemberType = "LOGICAL_PORT"
	IPSET                     MemberType = "IPSET"
)

func resourceSecurityGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSecurityGroupCreate,
		ReadContext:   resourceSecurityGroupRead,
		UpdateContext: resourceSecurityGroupUpdate,
		DeleteContext: resourceSecurityGroupDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(SecurityGroupRetryTimeout),
			Update: schema.DefaultTimeout(SecurityGroupRetryTimeout),
			Delete: schema.DefaultTimeout(SecurityGroupRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: resourceSecurityGroupImportState,
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
			"members": {
				Type:     schema.TypeSet,
				MinItems: 0,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
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

func resourceSecurityGroupImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	d.SetId(strings.ToLower(d.Id()))
	return []*schema.ResourceData{d}, nil
}

func resourceSecurityGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SecurityGroups

	proxy_pt := meta.(*sdk.Client).PlatformType
	platformType, _ := proxy_pt.Read(ctx)

	securityGroup, err_sg := mapResourceDataToSecurityGroup(d, platformType)
	if err_sg != nil {
		return diag.Errorf("ResourceSecurityGroupCreate >>>> error while mapping security group to resource: %+v", err_sg)

	}

	created, err := proxy.Create(ctx, securityGroup)
	if err != nil {
		return diag.Errorf("ResourceSecurityGroupCreate >>>> error while creating security group: %+v", err)
	}

	d.SetId(created.ID.String())

	return resourceSecurityGroupRead(ctx, d, meta)
}

func resourceSecurityGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SecurityGroups

	securityGroup, err := proxy.Read(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("security group with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading security group: %+v", err)
	}

	if err := d.Set("display_name", securityGroup.DisplayName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("project_id", securityGroup.ProjectID); err != nil {
		return diag.Errorf("error setting project_id: %+v", err)
	}

	if err := d.Set("members", flattenSecurityGroupMembers(securityGroup.Members)); err != nil {
		return diag.Errorf("error setting members: %+v", err)
	}

	if err := d.Set("created_by", securityGroup.CreatedBy); err != nil {
		return diag.Errorf("error setting created_by: %+v", err)
	}

	if err := d.Set("created_at", securityGroup.CreationDate.String()); err != nil {
		return diag.Errorf("error setting created_at: %+v", err)
	}

	if err := d.Set("modified_by", securityGroup.ModifiedBy); err != nil {
		return diag.Errorf("error setting modified_by: %+v", err)
	}

	if err := d.Set("modified_at", securityGroup.ModificationDate.String()); err != nil {
		return diag.Errorf("error setting modified_at: %+v", err)
	}

	return nil
}

func resourceSecurityGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SecurityGroups
	proxy_pt := meta.(*sdk.Client).PlatformType
	platformType, _ := proxy_pt.Read(ctx)
	securityGroup, err_pt := mapResourceDataToSecurityGroup(d, platformType)

	if err_pt != nil {
		return diag.Errorf("resourceSecurityGroupUpdate >>>> error while update security group: %+v", err_pt)

	}
	securityGroup.ID = strfmt.UUID(d.Id())

	_, err := proxy.Update(ctx, securityGroup)
	if err != nil {
		return diag.Errorf("error while modifying security group: %+v", err)
	}

	return resourceSecurityGroupRead(ctx, d, meta)
}

func resourceSecurityGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).SecurityGroups

	err := proxy.Delete(ctx, strfmt.UUID(d.Id()))
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("security group with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting security group: %+v", err)
	}

	return nil
}

func mapResourceDataToSecurityGroup(d *schema.ResourceData, platformType models.PlatformType) (*models.SecurityGroup, diag.Diagnostics) {
	members, err, wrongMemberType := expandSecurityGroupMembers(d.Get("members").(*schema.Set).List(), platformType)
	if err != nil {
		return nil, diag.Errorf("Wrong type member: %+v", wrongMemberType)
	}
	return &models.SecurityGroup{
		DisplayName: d.Get("display_name").(string),
		ProjectID:   strfmt.UUID(d.Get("project_id").(string)),
		Members:     members,
	}, nil
}
