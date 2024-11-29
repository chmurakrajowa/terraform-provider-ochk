package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	VirtualMachineRetryTimeout = 20 * time.Minute
	E1001                      = "TF_ERROR{1001}: Error during creating virtual machine. %s"
	E1002                      = "Input value MB_SIZE %d MB is less then 1024 MB"
	E1003                      = "Resource: %s is not supported in Openstack. Please remove %s from your terraform file."
	E1004                      = "KMS is not supported in Openstack. Please remove %s field from your terraform file."
	E1001_UPDATE               = "TF_ERROR{1001}: Error during updating virtual machine. %s"
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

		Importer: &schema.ResourceImporter{
			StateContext: resourceVirtualMachineImportState,
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
			"folder_path": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/",
			},
			"initial_password": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				// Setting password is possible only on create, subsequent read gets empty string,
				// which causes config drift. This suppresses any reported differences.
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return old == "<hidden>"
				},
			},
			"power_state": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"memory_size_mb": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"storage_policy": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				StateFunc: func(val any) string {
					return strings.ToLower(val.(string))
				},
			},
			"ssh_key": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"backup_lists": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"deployment_params": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"param_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"param_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"param_value": {
							Type:     schema.TypeString,
							ForceNew: true,
							Required: true,
						},
					},
				},
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
							StateFunc: func(val any) string {
								return strings.ToLower(val.(string))
							},
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
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"controller_id": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  "0",
						},
						"lun_id": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  "0",
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
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ovf_ip_configuration": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"initial_user_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary_dns_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_dns_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_suffix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_search_suffix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"primary_wins_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_wins_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVirtualMachineImportState(_ context.Context, d *schema.ResourceData, _ interface{}) ([]*schema.ResourceData, error) {
	d.SetId(strings.ToLower(d.Id()))
	return []*schema.ResourceData{d}, nil
}

func resourceVirtualMachineCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*sdk.Client)

	result := validateVirtualMachine(d, client.PType)
	if result != "" {
		return diag.Errorf(E1001, result)
	}
	virtualMachine := mapResourceDataToVirtualMachine(d)

	request, err := client.VirtualMachines.Create(ctx, virtualMachine)
	if err != nil {
		return diag.Errorf("error while creating virtual machine: %+v", err)
	}

	err, resourceID := client.Requests.FetchResourceID(ctx, d.Timeout(schema.TimeoutCreate), request)
	if err != nil {
		return diag.Errorf(E1001, err)
	}

	d.SetId(resourceID.String())

	return resourceVirtualMachineRead(ctx, d, meta)
}

func resourceVirtualMachineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	proxy := meta.(*sdk.Client).VirtualMachines

	virtualMachine, err := proxy.Read(ctx, strfmt.UUID(d.Id()))

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

	result := validateVirtualMachine(d, sdkClient.PType)
	if result != "" {
		return diag.Errorf(E1001_UPDATE, result)
	}

	virtualMachine := mapResourceDataToVirtualMachine(d)
	virtualMachine.VirtualMachineID = strfmt.UUID(d.Id())

	request, err := sdkClient.VirtualMachines.Update(ctx, virtualMachine)
	if err != nil {
		return diag.Errorf("error while modifying virtual machine: %+v", err)
	}

	err, resourceID := sdkClient.Requests.FetchResourceID(ctx, d.Timeout(schema.TimeoutCreate), request)
	if err != nil {
		return diag.Errorf("Error while fetching virtual machine request state: %+v", err)
	}

	d.SetId(resourceID.String())

	return resourceVirtualMachineRead(ctx, d, meta)
}

func resourceVirtualMachineDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sdkClient := meta.(*sdk.Client)

	request, err := sdkClient.VirtualMachines.Delete(ctx, strfmt.UUID(d.Id()))
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

func mapVirtualMachineToResourceData(d *schema.ResourceData, virtualMachine *models.VirtualMachineInstance) error {
	if err := d.Set("display_name", virtualMachine.VirtualMachineName); err != nil {
		return fmt.Errorf("error setting display_name: %w", err)
	}
	if virtualMachine.DeploymentInstance != nil {
		if err := d.Set("deployment_id", virtualMachine.DeploymentInstance.DeploymentID); err != nil {
			return fmt.Errorf("error setting deployment_id: %w", err)
		}
	}

	if err := d.Set("folder_path", virtualMachine.FolderPath); err != nil {
		return fmt.Errorf("error setting folder_path: %w", err)
	}

	if err := d.Set("initial_password", "<hidden>"); err != nil {
		return fmt.Errorf("error setting initial_password: %w", err)
	}
	if err := d.Set("power_state", virtualMachine.PowerState); err != nil {
		return fmt.Errorf("error setting power_state: %w", err)
	}

	if err := d.Set("cpu_count", int(virtualMachine.CPUCount)); err != nil {
		return fmt.Errorf("error setting cpu_count: %+v", err)
	}

	if err := d.Set("memory_size_mb", int(virtualMachine.MemorySizeMB)); err != nil {
		return fmt.Errorf("error setting memory_size_mb: %w", err)
	}

	if err := d.Set("storage_policy", virtualMachine.StoragePolicy); err != nil {
		return fmt.Errorf("error setting storage_policy: %w", err)
	}
	if err := d.Set("project_id", strings.ToLower(virtualMachine.ProjectID.String())); err != nil {
		return fmt.Errorf("error setting project_id: %w", err)
	}
	if err := d.Set("virtual_network_devices", flattenVirtualNetworkDevice(virtualMachine.VirtualNetworkDevices)); err != nil {
		return fmt.Errorf("error setting virtual_network_devices: %w", err)
	}

	if err := d.Set("additional_virtual_disks", flattenVirtualDisks(virtualMachine.AdditionalVirtualDiskDeviceCollection)); err != nil {
		return fmt.Errorf("error setting additional_virtual_disks: %w", err)
	}

	//virtualMachine.DeploymentParams = expandVDeploymentParams(d.Get("deployment_params").([]interface{}))
	//if err := d.Set("deployment_params", flattenDeploymentParams(virtualMachine.DeploymentParams)); err != nil {
	//	return fmt.Errorf("error setting deployment_params: %w", err)
	//}

	if err := d.Set("ip_address", virtualMachine.IPAddress); err != nil {
		return fmt.Errorf("error setting ip_address: %w", err)
	}

	var virtualDisks []*models.VirtualDiskDevice
	if virtualMachine.OsVirtualDiskDevice != nil {
		virtualDisks = append(virtualDisks, virtualMachine.OsVirtualDiskDevice)
	}

	if err := d.Set("virtual_disk", flattenVirtualDisks(virtualDisks)); err != nil {
		return fmt.Errorf("error setting virtual_disk: %w", err)
	}

	if virtualMachine.SSHKey != "" {
		if err := d.Set("ssh_key", virtualMachine.SSHKey); err != nil {
			return fmt.Errorf("error setting ssh_key: %w", err)
		}
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

	if len(virtualMachine.BackupListCollection) > 0 {
		if err := d.Set("backup_lists", flattenBackupListsFromIDs(virtualMachine.BackupListCollection)); err != nil {
			return fmt.Errorf("error setting backup lists: %w", err)
		}
	}
	if err := d.Set("tags", flattenTagsListsFromIDs(virtualMachine.Tags)); err != nil {
		return fmt.Errorf("error setting tags: %w", err)
	}
	if err := d.Set("primary_dns_address", virtualMachine.PrimaryDNSAddress); err != nil {
		return fmt.Errorf("error setting primary_dns_address: %+v", err)
	}
	if err := d.Set("secondary_dns_address", virtualMachine.SecondaryDNSAddress); err != nil {
		return fmt.Errorf("error setting secondary_dns_address: %+v", err)
	}
	if err := d.Set("dns_suffix", virtualMachine.DNSSuffix); err != nil {
		return fmt.Errorf("error setting dns_suffix: %+v", err)
	}
	if err := d.Set("dns_search_suffix", virtualMachine.DNSSearchSuffix); err != nil {
		return fmt.Errorf("error setting dns_search_suffix: %+v", err)
	}
	if err := d.Set("primary_wins_address", virtualMachine.PrimaryWinsAddress); err != nil {
		return fmt.Errorf("error setting primary_wins_address: %+v", err)
	}
	if err := d.Set("secondary_wins_address", virtualMachine.SecondaryWinsAddress); err != nil {
		return fmt.Errorf("error setting secondary_wins_address: %+v", err)
	}

	return nil
}

func mapResourceDataToVirtualMachine(d *schema.ResourceData) *models.VirtualMachineInstance {
	var virtualMachineInstance = models.VirtualMachineInstance{
		AdditionalVirtualDiskDeviceCollection: expandVirtualDisks(d.Get("additional_virtual_disks").(*schema.Set).List()),
		DeploymentInstance: &models.DeploymentInstance{
			DeploymentID: strfmt.UUID(d.Get("deployment_id").(string)),
		},
		InitialPassword:       d.Get("initial_password").(string),
		PowerState:            castStringToPowerStateEnum(d.Get("power_state").(string)),
		StoragePolicy:         castStringToStorageEnum(d.Get("storage_policy").(string)),
		ProjectID:             strfmt.UUID(d.Get("project_id").(string)),
		VirtualMachineID:      strfmt.UUID(d.Id()),
		VirtualMachineName:    d.Get("display_name").(string),
		SSHKey:                d.Get("ssh_key").(string),
		VirtualNetworkDevices: expandVirtualNetworkDevices(d.Get("virtual_network_devices").([]interface{})),
		//DeploymentParams:      expandVDeploymentParams(d.Get("deployment_params").([]interface{})),
		BackupListCollection: expandBackupListsFromIDs(d.Get("backup_lists").(*schema.Set).List()),
		Tags:                 expandTagsListsFromIDs(d.Get("tags").(*schema.Set).List()),
		OsType:               castStringToOsTypeEnum(d.Get("os_type").(string)),
		OvfIPConfiguration:   d.Get("ovf_ip_configuration").(bool),
		InitialUserName:      d.Get("initial_user_name").(string),
		FolderPath:           d.Get("folder_path").(string),
		DNSSearchSuffix:      d.Get("dns_search_suffix").(string),
		DNSSuffix:            d.Get("dns_suffix").(string),
		PrimaryDNSAddress:    d.Get("primary_dns_address").(string),
		PrimaryWinsAddress:   d.Get("primary_wins_address").(string),
		SecondaryDNSAddress:  d.Get("secondary_dns_address").(string),
		SecondaryWinsAddress: d.Get("secondary_wins_address").(string),
		CPUCount:             int32(d.Get("cpu_count").(int)),
		MemorySizeMB:         int32(d.Get("memory_size_mb").(int)),
	}
	encryptionInstance := &models.EncryptionInstance{
		Encrypt: d.Get("encryption").(bool),
	}

	if recryptOperation, ok := d.GetOk("encryption_recrypt"); ok && recryptOperation.(string) != "" {
		encryptionInstance.RecryptOperation = d.Get("encryption_recrypt").(models.RecryptOperation)
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

func castStringToOsTypeEnum(e string) models.OsType {
	switch e {
	case "WINDOWS":
		return models.OsTypeWINDOWS
	case "LINUX":
		return models.OsTypeLINUX
	default:
		return ""
	}
}

func castStringToStorageEnum(e string) models.StoragePolicy {
	switch e {
	case "UNKNOWN":
		return models.StoragePolicyUNKNOWN
	case "STANDARD":
		return models.StoragePolicySTANDARD
	case "STANDARD_W1":
		return models.StoragePolicySTANDARDW1
	case "STANDARD_W2":
		return models.StoragePolicySTANDARDW2
	case "ENTERPRISE":
		return models.StoragePolicyENTERPRISE
	case "STANDARDENCRYPTION":
		return models.StoragePolicySTANDARDENCRYPTION
	case "ENTERPRISEENCRYPTION":
		return models.StoragePolicyENTERPRISEENCRYPTION
	case "STANDARD_W1_ENCRYPTION":
		return models.StoragePolicySTANDARDW1ENCRYPTION
	case "STANDARD_W2_ENCRYPTION":
		return models.StoragePolicySTANDARDW2ENCRYPTION
	default:
		return ""
	}
}
