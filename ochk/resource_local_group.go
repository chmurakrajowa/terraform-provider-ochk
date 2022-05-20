package ochk

import (
	"context"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

const (
	LocalGroupRetryTimeout = 1 * time.Minute
)

func resourceLocalGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLocalGroupCreate,
		ReadContext:   resourceLocalGroupRead,
		UpdateContext: resourceLocalGroupUpdate,
		DeleteContext: resourceLocalGroupDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(LocalGroupRetryTimeout),
			Update: schema.DefaultTimeout(LocalGroupRetryTimeout),
			Delete: schema.DefaultTimeout(LocalGroupRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"users": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"child_groups": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"parent_groups": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceLocalGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).LocalGroups

	name := d.Get("name").(string)

	localGroups, err := proxy.ListByName(ctx, name)
	if err != nil {
		return diag.Errorf("error while listing local groups: %+v", err)
	}

	if len(localGroups) < 1 {
		return diag.Errorf("no local group found for name: %s", name)
	}

	if len(localGroups) > 1 {
		return diag.Errorf("more than one local group with name: %s found!", name)
	}

	d.SetId(localGroups[0].GroupID)

	if err := d.Set("name", localGroups[0].Name); err != nil {
		return diag.Errorf("error setting local group name: %+v", err)
	}
	if err := d.Set("description", localGroups[0].Description); err != nil {
		return diag.Errorf("error setting local group description: %+v", err)
	}
	if err := d.Set("group_type", localGroups[0].GroupType); err != nil {
		return diag.Errorf("error setting group_type: %+v", err)
	}
	if err := d.Set("domain", localGroups[0].Domain); err != nil {
		return diag.Errorf("error setting local group domain: %+v", err)
	}
	if err := d.Set("users", flattenUserInstancesFromIDs(localGroups[0].UserInstanceList)); err != nil {
		return diag.Errorf("error setting users: %+v", err)
	}
	if err := d.Set("parent_groups", flattenGroupInstancesFromIDs(localGroups[0].ParentGroups)); err != nil {
		return diag.Errorf("error setting parent_groups: %+v", err)
	}
	if err := d.Set("child_groups", flattenGroupInstancesFromIDs(localGroups[0].GroupInstanceList)); err != nil {
		return diag.Errorf("error setting child_groups: %+v", err)
	}
	return nil
}

func resourceLocalGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).LocalGroups

	_, err := proxy.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("local group id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting local group: %+v", err)
	}

	return nil
}

func mapResourceDataToLocalGroup(d *schema.ResourceData) *models.LocalGroup {
	localGroup := models.LocalGroup{
		GroupID:           d.Id(),
		Name:              d.Get("name").(string),
		Domain:            d.Get("domain").(string),
		Description:       d.Get("description").(string),
		GroupType:         d.Get("group_type").(string),
		GroupInstanceList: expandGroupsFromIDs(d.Get("child_groups").(*schema.Set).List()),
		ParentGroups:      expandGroupsFromIDs(d.Get("parent_groups").(*schema.Set).List()),
		UserInstanceList:  expandUsersFromIDs(d.Get("users").(*schema.Set).List()),
	}

	return &localGroup
}

func resourceLocalGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*sdk.Client)

	localGroup := mapResourceDataToLocalGroup(d)

	request, err := client.LocalGroups.Create(ctx, localGroup, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return diag.Errorf("error while creating local group: %+v", err)
	}

	err, resourceID := client.Requests.FetchResourceID(ctx, d.Timeout(schema.TimeoutCreate), request)
	if err != nil {
		return diag.Errorf("error while fetching local group request state: %+v", err)
	}

	d.SetId(resourceID)

	return resourceLocalGroupRead(ctx, d, meta)
}

func resourceLocalGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkClient := meta.(*sdk.Client)

	localGroup := mapResourceDataToLocalGroup(d)
	localGroup.GroupID = d.Id()

	request, err := sdkClient.LocalGroups.Update(ctx, localGroup)
	if err != nil {
		return diag.Errorf("error while modifying local group: %+v", err)
	}

	err, resourceID := sdkClient.Requests.FetchResourceID(ctx, d.Timeout(schema.TimeoutCreate), request)
	if err != nil {
		return diag.Errorf("error while fetching local group request state: %+v", err)
	}

	d.SetId(resourceID)

	return resourceLocalGroupRead(ctx, d, meta)
}
