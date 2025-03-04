package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
	"time"
)

const (
	SnapRetryTimeout = 15 * time.Minute
)

func resourceSnapshot() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnapshotCreate,
		ReadContext:   resourceSnapshotRead,
		DeleteContext: resourceSnapshotDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(SnapRetryTimeout),
			Update: schema.DefaultTimeout(SnapRetryTimeout),
			Delete: schema.DefaultTimeout(SnapRetryTimeout),
		},

		Importer: &schema.ResourceImporter{
			StateContext: resourceSnapshotImportState,
		},

		Schema: map[string]*schema.Schema{
			"virtual_machine_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"snapshot_description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"power_state": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"child_id": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
			"ram": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSnapshotImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	parts := strings.SplitN(d.Id(), "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%s), expected format: virtual_machine_id/snapshot_id", d.Id())
	}
	d.SetId(strings.ToLower(parts[1]))
	if err := d.Set("virtual_machine_id", strings.ToLower(parts[0])); err != nil {
		return nil, fmt.Errorf("cannot set virtual_machine_id: (%s)", parts[0])
	}
	return []*schema.ResourceData{d}, nil
}

func resourceSnapshotCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Snapshots

	virtualMachineId := strfmt.UUID(d.Get("virtual_machine_id").(string))
	ram := d.Get("ram").(bool)

	//if ram {
	//	err := d.Set("power_state", models.PowerStatePoweredOn)
	//	if err != nil {
	//		return nil
	//	}
	//} else {
	//	err := d.Set("power_state", models.PowerStatePoweredOff)
	//	if err != nil {
	//		return nil
	//	}
	//}

	snapshot := mapResourceDataToSnapshot(d)

	created, err := proxy.Create(ctx, virtualMachineId, &ram, snapshot)
	if err != nil {
		return diag.Errorf("error while creating snapshot: %+v", err)
	}

	d.SetId(created.SnapshotID.String())
	return resourceSnapshotRead(ctx, d, meta)
}

func resourceSnapshotRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Snapshots
	virtualMachineId := strfmt.UUID(d.Get("virtual_machine_id").(string))

	snapshot, err := proxy.Read(ctx, strfmt.UUID(d.Id()), virtualMachineId)
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("snapshot with id %s not found: %+v", id, err)
		}
		return diag.Errorf("error while reading snpashot: %+v", err)
	}

	if err := d.Set("display_name", snapshot.SnapshotName); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("virtual_machine_id", snapshot.VirtualMachineID); err != nil {
		return diag.Errorf("error setting virtual_machine_id: %+v", err)
	}
	if err := d.Set("snapshot_description", snapshot.SnapshotDescription); err != nil {
		return diag.Errorf("error setting snapshot_description: %+v", err)
	}
	if err := d.Set("power_state", snapshot.PowerState); err != nil {
		return diag.Errorf("error setting power_state: %+v", err)
	}
	if err := d.Set("parent_id", snapshot.ParentSnapshotID); err != nil {
		return diag.Errorf("error setting parent_id: %+v", err)
	}
	if err := d.Set("child_id", flattenChildsListsFromIDs(snapshot.ChildSnapshots)); err != nil {
		return diag.Errorf("error setting child_id: %+v", err)
	}
	return nil
}

func resourceSnapshotDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).Snapshots

	virtualMachineID := strfmt.UUID(d.Get("virtual_machine_id").(string))

	err := proxy.Delete(ctx, virtualMachineID, strfmt.UUID(d.Id()))

	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("snapshot with id %s not found: %+v", id, err)
		}
		return diag.Errorf("error while deleting snapshot: %+v", err)
	}
	return nil
}

func mapResourceDataToSnapshot(d *schema.ResourceData) *models.SnapshotInstance {
	return &models.SnapshotInstance{
		SnapshotName:        d.Get("display_name").(string),
		SnapshotDescription: d.Get("snapshot_description").(string),
		VirtualMachineID:    strfmt.UUID(d.Get("virtual_machine_id").(string)),
		PowerState:          castStringToPowerStateEnum(d.Get("power_state").(string)),
		ParentSnapshotID:    strfmt.UUID(d.Get("parent_id").(string)),
		ChildSnapshots:      expandChildSnapshots(d.Get("child_id").(*schema.Set).List()),
	}
}

func castStringToPowerStateEnum(e string) models.PowerState {
	switch e {
	case "poweredOff":
		return models.PowerStatePoweredOff
	case "poweredOn":
		return models.PowerStatePoweredOn
	default:
		return ""
	}
}
