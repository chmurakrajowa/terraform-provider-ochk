package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type VirtualDiskTestData struct {
	ControllerID int
	LunID        int
	SizeMB       int
	DeviceType   string
}

type VirtualMachineTestData struct {
	ResourceName           string
	DisplayName            string
	DeploymentID           strfmt.UUID
	InitialPassword        string
	PowerState             string
	VirtualDiskSizeMB      string
	MemorySizeMb           string
	CpuCount               string
	StoragePolicy          string
	ProjectID              strfmt.UUID
	AdditionalVirtualDisks []VirtualDiskTestData
	VirtualDiskDevice      *VirtualDiskTestData
	VirtualNetworkDevices  []struct {
		VirtualNetworkID strfmt.UUID
	}
	Encryption                     bool
	EncryptionKeyID                strfmt.UUID
	EncryptionPrivateKeyIDToUnwrap string
	EncryptionRecrypt              string

	PrimaryDNSAddress    string
	SecondaryDNSAddress  string
	DNSSuffix            string
	DNSSearchSuffix      string
	PrimaryWinsAddress   string
	SecondaryWinsAddress string
}

func (c *VirtualMachineTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_backup_plan" "backup_plan" {
  display_name = "`+testData.BackupPlanName+`"
}

data "ochk_backup_list" "backup_list" {
  display_name = "`+testData.BackupListName+`"
  backup_plan_id = data.ochk_backup_plan.backup_plan.id
}

data "ochk_tag" "vm-tag" {
	display_name = "`+testData.TagName+`"
}

resource "ochk_virtual_machine" "{{ .ResourceName}}" {
	display_name = "{{.DisplayName}}"
	deployment_id = {{ UuidTFValue .DeploymentID }}
	initial_password = "{{.InitialPassword}}"
	power_state  = "{{.PowerState}}"
	cpu_count = {{.CpuCount}}
	memory_size_mb = {{.MemorySizeMb}}
	virtual_disk {
		size_mb = {{.VirtualDiskSizeMB}}
	}
	storage_policy  = "{{.StoragePolicy}}"
	project_id  = {{ UuidTFValue .ProjectID}}
	{{range .VirtualNetworkDevices}}
	virtual_network_devices {
		virtual_network_id = {{ UuidTFValue .VirtualNetworkID }}
    }
	{{end}}
	{{range .AdditionalVirtualDisks}}
	additional_virtual_disks {
		controller_id = {{ .ControllerID }}
		lun_id = {{.LunID}}
		size_mb = {{.SizeMB}}
		device_type = "{{.DeviceType}}"
    }
	{{end}}
	{{if .VirtualDiskDevice}}
	virtual_disk {
		controller_id = {{ .VirtualDiskDevice.ControllerID }}
		lun_id = {{.VirtualDiskDevice.LunID}}
		size_mb = {{.VirtualDiskDevice.SizeMB}}
		device_type = "{{.VirtualDiskDevice.DeviceType}}"
	}
	{{end}}
	encryption = {{ .Encryption }}

  	tags = [
    	data.ochk_tag.vm-tag.id
  	]

  	backup_lists = [
    	data.ochk_backup_list.backup_list.id
  	]

	dns_search_suffix = "{{ .DNSSearchSuffix }}"
	dns_suffix = "{{ .DNSSuffix }}"
	primary_dns_address = "{{ .PrimaryDNSAddress }}"
	primary_wins_address = "{{ .PrimaryWinsAddress }}"
	secondary_dns_address = "{{ .SecondaryDNSAddress }}"
	secondary_wins_address = "{{ .SecondaryWinsAddress }}"

	{{if .EncryptionKeyID }}
	encryption_key_id = {{ UuidTFValue .EncryptionKeyID }}
	{{end}}
	{{if .EncryptionPrivateKeyIDToUnwrap }}
	encryption_private_key_id_to_unwrap = {{ UuidTFValue .EncryptionPrivateKeyIDToUnwrap }}
	{{end}}
	{{if .EncryptionRecrypt }}
	encryption_recrypt = "{{ .EncryptionRecrypt }}"
	{{end}}
}
`, c)
}

func (c *VirtualMachineTestData) FullResourceName() string {
	return "ochk_virtual_machine." + c.ResourceName
}

func TestAccVirtualMachineResource_create_update_minimal(t *testing.T) {
	project1 := ProjectDataSourceTestData{
		ResourceName: "project1",
		DisplayName:  testData.ProjectForVMName,
	}

	deployment := DeploymentDataSourceTestData{
		ResourceName: "deployment",
		DisplayName:  testData.Deployment1DisplayName,
	}

	vnet1 := VirtualNetworkDataSourceTestData{
		ResourceName: "vnet1",
		DisplayName:  testData.VirtualNetwork1DisplayName,
	}

	virtualMachine := VirtualMachineTestData{
		ResourceName:      "default",
		DisplayName:       generateShortRandName(devTestDataPrefix),
		DeploymentID:      strfmt.UUID(testDataResourceID(&deployment)),
		InitialPassword:   "50b90880f9f",
		PowerState:        "poweredOn",
		VirtualDiskSizeMB: "40960",
		CpuCount:          "2",
		MemorySizeMb:      "4096",
		StoragePolicy:     "STANDARD_W1",
		ProjectID:         strfmt.UUID(testDataResourceID(&project1)),
		VirtualNetworkDevices: []struct{ VirtualNetworkID strfmt.UUID }{
			{VirtualNetworkID: strfmt.UUID(testDataResourceID(&vnet1))},
		},
		AdditionalVirtualDisks: []VirtualDiskTestData{
			{
				ControllerID: 0,
				LunID:        2,
				SizeMB:       20480,
				DeviceType:   "SCSI",
			},
			{
				ControllerID: 0,
				LunID:        1,
				SizeMB:       40960,
				DeviceType:   "SCSI",
			},
		},
		PrimaryDNSAddress:    "192.168.1.6",
		SecondaryDNSAddress:  "192.168.1.2",
		DNSSuffix:            "test.lcl",
		DNSSearchSuffix:      "test.lcl,prod.lcl",
		PrimaryWinsAddress:   "192.168.1.3",
		SecondaryWinsAddress: "192.168.1.3",
	}

	configInitial := deployment.ToString() + project1.ToString() + vnet1.ToString("-vm-in") + virtualMachine.ToString()
	virtualMachineUpdated := virtualMachine
	virtualMachineUpdated.DisplayName += "upd"
	virtualMachineUpdated.StoragePolicy = "ENTERPRISE"
	virtualMachineUpdated.CpuCount = "4"
	virtualMachineUpdated.MemorySizeMb = "8192"
	virtualMachineUpdated.VirtualDiskSizeMB = "81920"
	//FIXME backend does not support this yet
	//virtualMachineUpdated.PowerState = "poweredOff"
	virtualMachineUpdated.AdditionalVirtualDisks = []VirtualDiskTestData{
		{
			ControllerID: 0,
			LunID:        2,
			SizeMB:       20480,
			DeviceType:   "SCSI",
		},
		{
			ControllerID: 0,
			LunID:        1,
			SizeMB:       40960,
			DeviceType:   "SCSI",
		},
		{
			ControllerID: 0,
			LunID:        3,
			SizeMB:       30960,
			DeviceType:   "SCSI",
		},
	}
	configUpdated := deployment.ToString() + project1.ToString() + vnet1.ToString("-vm-upt") + virtualMachineUpdated.ToString()

	resourceName := virtualMachine.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,

		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", virtualMachine.DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "deployment_id", deployment.FullResourceName(), "id"),
					resource.TestCheckResourceAttr(resourceName, "power_state", virtualMachine.PowerState),
					resource.TestCheckResourceAttr(resourceName, "cpu_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "memory_size_mb", "4096"),
					resource.TestCheckResourceAttr(resourceName, "virtual_disk.0.size_mb", "40960"),
					resource.TestCheckResourceAttr(resourceName, "storage_policy", virtualMachine.StoragePolicy),
					resource.TestCheckResourceAttr(resourceName, "initial_password", "<hidden>"),
					resource.TestCheckResourceAttr(resourceName, "primary_dns_address", virtualMachine.PrimaryDNSAddress),
					resource.TestCheckResourceAttr(resourceName, "secondary_dns_address", virtualMachine.SecondaryDNSAddress),
					resource.TestCheckResourceAttr(resourceName, "dns_suffix", virtualMachine.DNSSuffix),
					resource.TestCheckResourceAttr(resourceName, "dns_search_suffix", virtualMachine.DNSSearchSuffix),
					resource.TestCheckResourceAttr(resourceName, "primary_wins_address", virtualMachine.PrimaryWinsAddress),
					resource.TestCheckResourceAttr(resourceName, "secondary_wins_address", virtualMachine.SecondaryWinsAddress),
					resource.TestCheckResourceAttrPair(resourceName, "project_id", project1.FullResourceName(), "id"),
					resource.TestCheckResourceAttrPair(resourceName, "backup_lists.0", "data.ochk_backup_list.backup_list", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "tags.0", "data.ochk_tag.vm-tag", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "virtual_network_devices.0.virtual_network_id", vnet1.FullResourceName(), "id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_network_devices.0.device_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_disk.0.controller_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_disk.0.lun_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_disk.0.size_mb"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_disk.0.device_type"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.0.controller_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.0.lun_id", "1"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.0.size_mb", "40960"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.0.device_type", "SCSI"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.1.controller_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.1.lun_id", "2"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.1.size_mb", "20480"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.1.device_type", "SCSI"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: configUpdated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", virtualMachineUpdated.DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "deployment_id", deployment.FullResourceName(), "id"),
					resource.TestCheckResourceAttr(resourceName, "power_state", virtualMachineUpdated.PowerState),
					resource.TestCheckResourceAttr(resourceName, "cpu_count", "4"),
					resource.TestCheckResourceAttr(resourceName, "memory_size_mb", "8192"),
					resource.TestCheckResourceAttr(resourceName, "virtual_disk.0.size_mb", "81920"),
					resource.TestCheckResourceAttr(resourceName, "storage_policy", virtualMachineUpdated.StoragePolicy),
					resource.TestCheckResourceAttrPair(resourceName, "project_id", project1.FullResourceName(), "id"),
					resource.TestCheckResourceAttrPair(resourceName, "virtual_network_devices.0.virtual_network_id", vnet1.FullResourceName(), "id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_network_devices.0.device_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_disk.0.controller_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_disk.0.lun_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_disk.0.size_mb"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_disk.0.device_type"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.0.controller_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.0.lun_id", "1"),
					//FIXME backend does not update this properly yet
					//resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.0.size_mb", "50960"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.0.device_type", "SCSI"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.1.controller_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.1.lun_id", "2"),
					//FIXME backend does not update this properly yet
					//resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.1.size_mb", "40480"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.1.device_type", "SCSI"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.2.controller_id", "0"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.2.lun_id", "3"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.2.size_mb", "30960"),
					resource.TestCheckResourceAttr(resourceName, "additional_virtual_disks.2.device_type", "SCSI"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
				),
			},
		},
		CheckDestroy: testAccVirtualMachineResourceNotExists(virtualMachine.DisplayName),
	})
}

func TestAccVirtualMachineResource_create_with_managed_encryption(t *testing.T) {
	project1 := ProjectDataSourceTestData{
		ResourceName: "project1",
		DisplayName:  testData.ProjectForVMName,
	}

	deployment := DeploymentDataSourceTestData{
		ResourceName: "deployment",
		DisplayName:  testData.Deployment1DisplayName,
	}

	vnet1 := VirtualNetworkDataSourceTestData{
		ResourceName: "vnet1",
		DisplayName:  testData.VirtualNetwork1DisplayName,
	}

	kmsAESKey := KMSKeyTestData{
		ResourceName: "key",
		DisplayName:  generateShortRandName(devTestDataPrefix),
		KeyUsage:     []string{"ENCRYPT", "DECRYPT"},
		Algorithm:    "AES",
		Size:         256,
	}

	virtualMachine := VirtualMachineTestData{
		ResourceName:      "default",
		DisplayName:       generateShortRandName(devTestDataPrefix),
		DeploymentID:      strfmt.UUID(testDataResourceID(&deployment)),
		InitialPassword:   "50b90880f9f",
		PowerState:        "poweredOn",
		CpuCount:          "2",
		MemorySizeMb:      "4096",
		VirtualDiskSizeMB: "40960",
		StoragePolicy:     "STANDARD_W1",
		ProjectID:         strfmt.UUID(testDataResourceID(&project1)),
		VirtualNetworkDevices: []struct{ VirtualNetworkID strfmt.UUID }{
			{VirtualNetworkID: strfmt.UUID(testDataResourceID(&vnet1))},
		},
		Encryption: true,
	}

	configInitial := deployment.ToString() + project1.ToString() + vnet1.ToString("-vm-m-in") + kmsAESKey.ToString() + virtualMachine.ToString()
	virtualMachineUpdated := virtualMachine
	virtualMachineUpdated.EncryptionKeyID = strfmt.UUID(testDataResourceID(&kmsAESKey))
	virtualMachineUpdated.EncryptionRecrypt = "SHALLOW"

	configUpdated := deployment.ToString() + project1.ToString() + vnet1.ToString("-vm-m-up") + kmsAESKey.ToString() + virtualMachineUpdated.ToString()

	resourceName := virtualMachine.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,

		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check:  virtualMachineChecks(resourceName, virtualMachine, project1, deployment, vnet1),
			},
			{
				Config: configUpdated,
				Check:  virtualMachineChecks(resourceName, virtualMachineUpdated, project1, deployment, vnet1),
			},
		},
		CheckDestroy: testAccVirtualMachineResourceNotExists(virtualMachine.DisplayName),
	})
}

func TestAccVirtualMachineResource_create_with_own_encryption(t *testing.T) {
	project1 := ProjectDataSourceTestData{
		ResourceName: "project1",
		DisplayName:  testData.ProjectForVMName,
	}

	deployment := DeploymentDataSourceTestData{
		ResourceName: "deployment",
		DisplayName:  testData.Deployment1DisplayName,
	}

	vnet1 := VirtualNetworkDataSourceTestData{
		ResourceName: "vnet1",
		DisplayName:  testData.VirtualNetwork1DisplayName,
	}

	kmsAESKey := KMSKeyTestData{
		ResourceName: "key",
		DisplayName:  generateShortRandName(devTestDataPrefix),
		KeyUsage:     []string{"ENCRYPT", "DECRYPT"},
		Algorithm:    "AES",
		Size:         256,
	}

	virtualMachine := VirtualMachineTestData{
		ResourceName:      "default",
		DisplayName:       generateShortRandName(devTestDataPrefix),
		DeploymentID:      strfmt.UUID(testDataResourceID(&deployment)),
		InitialPassword:   "50b90880f9f",
		PowerState:        "poweredOn",
		CpuCount:          "2",
		MemorySizeMb:      "4096",
		VirtualDiskSizeMB: "40960",
		StoragePolicy:     "STANDARD_W1",
		ProjectID:         strfmt.UUID(testDataResourceID(&project1)),
		VirtualNetworkDevices: []struct{ VirtualNetworkID strfmt.UUID }{
			{VirtualNetworkID: strfmt.UUID(testDataResourceID(&vnet1))},
		},
		Encryption:      true,
		EncryptionKeyID: strfmt.UUID(testDataResourceID(&kmsAESKey)),
	}

	configInitial := deployment.ToString() + project1.ToString() + vnet1.ToString("-vm-enc-in") + kmsAESKey.ToString() + virtualMachine.ToString()

	virtualMachineUpdated := virtualMachine
	virtualMachineUpdated.EncryptionKeyID = ""
	virtualMachineUpdated.EncryptionRecrypt = "SHALLOW"

	configUpdated := deployment.ToString() + project1.ToString() + vnet1.ToString("-vm-enc-up") + kmsAESKey.ToString() + virtualMachineUpdated.ToString()

	resourceName := virtualMachine.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,

		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check:  virtualMachineChecks(resourceName, virtualMachine, project1, deployment, vnet1),
			},
			{
				Config: configUpdated,
				Check:  virtualMachineChecks(resourceName, virtualMachineUpdated, project1, deployment, vnet1),
			},
		},
		CheckDestroy: testAccVirtualMachineResourceNotExists(virtualMachine.DisplayName),
	})
}
func virtualMachineChecks(resourceName string, vm VirtualMachineTestData, project ProjectDataSourceTestData, deployment DeploymentDataSourceTestData, vnet VirtualNetworkDataSourceTestData) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "display_name", vm.DisplayName),
		resource.TestCheckResourceAttrPair(resourceName, "deployment_id", deployment.FullResourceName(), "id"),
		resource.TestCheckResourceAttr(resourceName, "power_state", vm.PowerState),
		resource.TestCheckResourceAttr(resourceName, "initial_password", "<hidden>"),
		resource.TestCheckResourceAttr(resourceName, "cpu_count", vm.CpuCount),
		resource.TestCheckResourceAttr(resourceName, "memory_size_mb", vm.MemorySizeMb),
		resource.TestCheckResourceAttr(resourceName, "virtual_disk.0.size_mb", vm.VirtualDiskSizeMB),
		resource.TestCheckResourceAttr(resourceName, "folder_path", "/"),
		resource.TestCheckResourceAttr(resourceName, "storage_policy", vm.StoragePolicy),
		resource.TestCheckResourceAttrPair(resourceName, "project_id", project.FullResourceName(), "id"),
		resource.TestCheckResourceAttrPair(resourceName, "virtual_network_devices.0.virtual_network_id", vnet.FullResourceName(), "id"),
		resource.TestCheckResourceAttrSet(resourceName, "virtual_network_devices.0.device_id"),
		resource.TestCheckResourceAttrSet(resourceName, "virtual_disk.0.controller_id"),
		resource.TestCheckResourceAttrSet(resourceName, "virtual_disk.0.lun_id"),
		resource.TestCheckResourceAttrSet(resourceName, "virtual_disk.0.size_mb"),
		resource.TestCheckResourceAttrSet(resourceName, "virtual_disk.0.device_type"),
		resource.TestCheckResourceAttrSet(resourceName, "created_by"),
		resource.TestCheckResourceAttrSet(resourceName, "created_at"),
		resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
		resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
	)
}

func testAccVirtualMachineResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).VirtualMachines

		virtualMachines, err := proxy.ListByDisplayName(context.Background(), displayName)
		if err != nil {
			return err
		}

		if len(virtualMachines) > 0 {
			return fmt.Errorf("virtual machine %s still exists", virtualMachines[0].VirtualMachineID)
		}

		return nil
	}
}
