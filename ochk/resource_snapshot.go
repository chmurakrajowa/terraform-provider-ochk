package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
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
		UpdateContext: resourceSnapshotUpdate,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {

			if d.Id() == "" {
				return nil
			}

			//if d.Id() != "" && d.HasChange("ram") {
			//	return fmt.Errorf("field1 cannot be updated")
			//}
			//if d.Id() == "" { // only for update if d.Id() exists it means we can run validate method
			//	return nil
			//}
			//snapshot_id := d.Id()
			//
			//fmt.Println("######### CustomizeDiff >>> " + d.Id() + " ")
			//fmt.Println("######### CustomizeDiff >>> ram %s", d.HasChange("ram"))
			//if d.HasChange("ram") {
			//	return fmt.Errorf("updating snapshot %s%s%s", snapshot_id, d.HasChange("ram"), d.Get("display_name"), " is not supported by API, field ram can not be changed")
			//}
			//} else if d.HasChange("display_name") {
			//	return fmt.Errorf("updating snapshot %s", snapshot_id, " is not supported by API, field display_name can not be changed")
			//} else if d.HasChange("snapshot_description") {
			//	return fmt.Errorf("updating snapshot %s", snapshot_id, " is not supported by API, field snapshot_description can not be changed")
			//} else if d.HasChange("virtual_machine_id") {
			//	return fmt.Errorf("updating snapshot %s", snapshot_id, " is not supported by API, field virtual_machine_id can not be changed")
			//}
			return nil
		},

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
				ForceNew: false,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"snapshot_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"power_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"child_id": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ram": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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
	proxy_vm := meta.(*sdk.Client).VirtualMachines

	virtualMachineId := strfmt.UUID(d.Get("virtual_machine_id").(string))
	ram := d.Get("ram").(bool)

	vm, _, err := proxy_vm.Read(ctx, virtualMachineId)

	if ram && vm.GetPowerState() == openapi.POWERSTATE_POWERED_OFF {
		return diag.Errorf("error while creating snapshot with ram for virtual machine %s. Machine is powered off.", vm.GetVirtualMachineName())
	}

	if ram {
		err := d.Set("power_state", openapi.POWERSTATE_POWERED_ON)
		if err != nil {
			return nil
		}
	} else {
		err := d.Set("power_state", openapi.POWERSTATE_POWERED_OFF)
		if err != nil {
			return nil
		}
	}

	snapshot := mapResourceDataToSnapshot(d)

	created, _, err := proxy.Create(ctx, virtualMachineId, ram, snapshot)
	if err != nil {
		return diag.Errorf("error while creating snapshot: %+v", err)
	}

	d.SetId(created.GetSnapshotId())
	return resourceSnapshotRead(ctx, d, meta)
}

func resourceSnapshotUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	//return diag.Errorf("error while creating snapshot: %+v", "ZLE")
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
		return diag.Errorf("error while reading snapshot sss : %+v", err)
	}

	if err := d.Set("display_name", snapshot.GetSnapshotName()); err != nil {
		return diag.Errorf("error setting display_name: %+v", err)
	}

	if err := d.Set("virtual_machine_id", snapshot.GetVirtualMachineId()); err != nil {
		return diag.Errorf("error setting virtual_machine_id: %+v", err)
	}
	if err := d.Set("snapshot_description", snapshot.GetSnapshotDescription()); err != nil {
		return diag.Errorf("error setting snapshot_description: %+v", err)
	}
	if err := d.Set("power_state", snapshot.GetPowerState()); err != nil {
		return diag.Errorf("error setting power_state: %+v", err)
	}
	if err := d.Set("parent_id", snapshot.GetParentSnapshotId()); err != nil {
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

func mapResourceDataToSnapshot(d *schema.ResourceData) openapi.SnapshotInstance {
	return openapi.SnapshotInstance{
		SnapshotName:        NewNullableString(d.Get("display_name").(string)),
		SnapshotDescription: NewNullableString(d.Get("snapshot_description").(string)),
		VirtualMachineId:    NewNullableString(d.Get("virtual_machine_id").(string)),
		PowerState:          castStringToPowerStateEnum(d.Get("power_state").(string)).Ptr(),
		ParentSnapshotId:    openapi.NullableString{},
		ChildSnapshots:      expandChildSnapshots(d.Get("child_id").(*schema.Set).List()),
		SnapshotId:          openapi.NullableString{},
	}
}

func castStringToPowerStateEnum(e string) openapi.PowerState {
	switch e {
	case "poweredOff":
		return openapi.POWERSTATE_POWERED_OFF
	case "poweredOn":
		return openapi.POWERSTATE_POWERED_ON
	default:
		return ""
	}
}
