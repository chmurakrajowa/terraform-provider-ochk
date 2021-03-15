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
	VirtualMachineRetryTimeout = 20 * time.Minute
)

func resourceVirtualMachine() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVirtualMachineCreate,
		ReadContext:   resourceVirtualMachineRead,
		UpdateContext: resourceVirtualMachineUpdate,
		DeleteContext: resourceVirtualMachineDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(VirtualMachineRetryTimeout),
			Update: schema.DefaultTimeout(VirtualMachineRetryTimeout),
			Delete: schema.DefaultTimeout(VirtualMachineRetryTimeout),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deployment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"initial_password": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				// Setting password is possible only on create, subsequent read gets empty string,
				// which causes config drift. This suppresses any reported differences.
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return true
				},
			},
			"power_state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_profile": {
				Type:     schema.TypeString,
				Required: true,
			},
			"storage_policy": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subtenant_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"virtual_network_devices": {
				Type:     schema.TypeList,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"virtual_network_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"device_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"additional_virtual_disks": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"controller_id": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  "0",
						},
						"lun_id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"size_mb": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"device_type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "SCSI",
						},
					},
				},
			},
			"virtual_disk": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"controller_id": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  "0",
						},
						"lun_id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"size_mb": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"device_type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "SCSI",
						},
					},
				},
			},
			"encryption": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"encryption_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encryption_recrypt": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceVirtualMachineCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*sdk.Client)

	virtualMachine := mapResourceDataToVirtualMachine(d)

	request, err := client.VirtualMachines.Create(ctx, virtualMachine)
	if err != nil {
		return diag.Errorf("error while creating virtual machine: %+v", err)
	}

	err, resourceID := client.Requests.FetchResourceID(ctx, d.Timeout(schema.TimeoutCreate), request)
	if err != nil {
		return diag.Errorf("error while fetching virtual machine request state: %+v", err)
	}

	d.SetId(resourceID)

	return resourceVirtualMachineRead(ctx, d, meta)
}

func resourceVirtualMachineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).VirtualMachines

	virtualMachine, err := proxy.Read(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("virtual machine with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while reading virtual machine: %+v", err)
	}

	if err := mapVirtualMachineToResourceData(d, virtualMachine); err != nil {
		return diag.Errorf("error setting virtual machine: %v", err)
	}

	return nil
}

func resourceVirtualMachineUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkClient := meta.(*sdk.Client)

	virtualMachine := mapResourceDataToVirtualMachine(d)
	virtualMachine.VirtualMachineID = d.Id()

	request, err := sdkClient.VirtualMachines.Update(ctx, virtualMachine)
	if err != nil {
		return diag.Errorf("error while modifying virtual machine: %+v", err)
	}

	err, resourceID := sdkClient.Requests.FetchResourceID(ctx, d.Timeout(schema.TimeoutCreate), request)
	if err != nil {
		return diag.Errorf("error while fetching virtual machine request state: %+v", err)
	}

	d.SetId(resourceID)

	return resourceVirtualMachineRead(ctx, d, meta)
}

func resourceVirtualMachineDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkClient := meta.(*sdk.Client)

	request, err := sdkClient.VirtualMachines.Delete(ctx, d.Id())
	if err != nil {
		if sdk.IsNotFoundError(err) {
			id := d.Id()
			d.SetId("")
			return diag.Errorf("virtual machine with id %s not found: %+v", id, err)
		}

		return diag.Errorf("error while deleting virtual machine: %+v", err)
	}

	err, _ = sdkClient.Requests.FetchResourceID(ctx, d.Timeout(schema.TimeoutCreate), request)
	if err != nil {
		return diag.Errorf("error while fetching virtual machine request state: %+v", err)
	}

	return nil
}

func mapVirtualMachineToResourceData(d *schema.ResourceData, virtualMachine *models.VcsVirtualMachineInstance) error {
	if err := d.Set("display_name", virtualMachine.VirtualMachineName); err != nil {
		return fmt.Errorf("error setting display_name: %w", err)
	}
	if virtualMachine.DeploymentInstance != nil {
		if err := d.Set("deployment_id", virtualMachine.DeploymentInstance.DeploymentID); err != nil {
			return fmt.Errorf("error setting deployment_id: %w", err)
		}
	}
	if err := d.Set("initial_password", virtualMachine.InitialPassword); err != nil {
		return fmt.Errorf("error setting initial_password: %w", err)
	}
	if err := d.Set("power_state", virtualMachine.PowerState); err != nil {
		return fmt.Errorf("error setting power_state: %w", err)
	}
	if err := d.Set("resource_profile", virtualMachine.ResourceProfile); err != nil {
		return fmt.Errorf("error setting resource_profile: %w", err)
	}
	if err := d.Set("storage_policy", virtualMachine.StoragePolicy); err != nil {
		return fmt.Errorf("error setting storage_policy: %w", err)
	}
	if err := d.Set("subtenant_id", virtualMachine.SubtenantRefID); err != nil {
		return fmt.Errorf("error setting subtenant_id: %w", err)
	}
	if err := d.Set("virtual_network_devices", flattenVirtualNetworkDevice(virtualMachine.VirtualNetworkDevices)); err != nil {
		return fmt.Errorf("error setting virtual_network_devices: %w", err)
	}
	if err := d.Set("additional_virtual_disks", flattenVirtualDisks(virtualMachine.AdditionalVirtualDiskDeviceCollection)); err != nil {
		return fmt.Errorf("error setting additional_virtual_disks: %w", err)
	}

	var virtualDisks []*models.VirtualDiskDevice
	if virtualMachine.OsVirtualDiskDevice != nil {
		virtualDisks = append(virtualDisks, virtualMachine.OsVirtualDiskDevice)
	}
	if err := d.Set("virtual_disk", flattenVirtualDisks(virtualDisks)); err != nil {
		return fmt.Errorf("error setting virtual_disk: %w", err)
	}

	if virtualMachine.EncryptionInstance != nil {
		if err := d.Set("encryption", virtualMachine.EncryptionInstance.Encrypt); err != nil {
			return fmt.Errorf("error setting created_by: %w", err)
		}
		if virtualMachine.EncryptionInstance.EncryptionKeyID != "" {
			if err := d.Set("encryption_key_id", virtualMachine.EncryptionInstance.EncryptionKeyID); err != nil {
				return fmt.Errorf("error setting created_by: %w", err)
			}
		}
	}

	if err := d.Set("created_by", virtualMachine.CreatedBy); err != nil {
		return fmt.Errorf("error setting created_by: %w", err)
	}
	if err := d.Set("created_at", virtualMachine.CreationDate.String()); err != nil {
		return fmt.Errorf("error setting created_at: %w", err)
	}
	if err := d.Set("modified_at", virtualMachine.ModificationDate.String()); err != nil {
		return fmt.Errorf("error setting modified_at: %w", err)
	}
	if err := d.Set("modified_by", virtualMachine.ModifiedBy); err != nil {
		return fmt.Errorf("error setting modified_by: %w", err)
	}

	return nil
}

func mapResourceDataToVirtualMachine(d *schema.ResourceData) *models.VcsVirtualMachineInstance {
	virtualMachineInstance := models.VcsVirtualMachineInstance{
		AdditionalVirtualDiskDeviceCollection: expandVirtualDisks(d.Get("additional_virtual_disks").(*schema.Set).List()),
		DeploymentInstance: &models.DeploymentInstance{
			DeploymentID: d.Get("deployment_id").(string),
		},
		InitialPassword:       d.Get("initial_password").(string),
		PowerState:            d.Get("power_state").(string),
		ResourceProfile:       d.Get("resource_profile").(string),
		StoragePolicy:         d.Get("storage_policy").(string),
		SubtenantRefID:        d.Get("subtenant_id").(string),
		VirtualMachineID:      d.Id(),
		VirtualMachineName:    d.Get("display_name").(string),
		VirtualNetworkDevices: expandVirtualNetworkDevices(d.Get("virtual_network_devices").([]interface{})),
	}

	encryptionInstance := &models.EncryptionInstance{
		Encrypt: d.Get("encryption").(bool),
	}

	if recryptOperation, ok := d.GetOk("encryption_recrypt"); ok && recryptOperation.(string) != "" {
		encryptionInstance.RecryptOperation = d.Get("encryption_recrypt").(string)
	} else {
		encryptionInstance.RecryptOperation = "NONE"
	}

	if encryptionKeyId, ok := d.GetOk("encryption_key_id"); ok && encryptionKeyId.(string) != "" {
		encryptionInstance.EncryptionKeyID = encryptionKeyId.(string)
		encryptionInstance.Managed = false
	} else {
		encryptionInstance.Managed = true
	}

	if !encryptionInstance.Managed || encryptionInstance.Encrypt {
		virtualMachineInstance.EncryptionInstance = encryptionInstance
	}

	virtualDisks := expandVirtualDisks(d.Get("virtual_disk").(*schema.Set).List())
	if len(virtualDisks) > 0 {
		virtualMachineInstance.OsVirtualDiskDevice = virtualDisks[0]
	}

	return &virtualMachineInstance
}
