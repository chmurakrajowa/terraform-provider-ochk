package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
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
	DeploymentID           string
	InitialPassword        string
	PowerState             string
	ResourceProfile        string
	StoragePolicy          string
	SubtenantID            string
	AdditionalVirtualDisks []VirtualDiskTestData
	VirtualDiskDevice      *VirtualDiskTestData
	VirtualNetworkDevices  []struct {
		VirtualNetworkID string
	}
	Encryption                     bool
	EncryptionKeyID                string
	EncryptionPrivateKeyIDToUnwrap string
	EncryptionRecrypt              string
}

func (c *VirtualMachineTestData) ToString() string {
	return executeTemplateToString(`
resource "ochk_virtual_machine" "{{ .ResourceName}}" {
	display_name = "{{.DisplayName}}"
	deployment_id = {{ StringTFValue .DeploymentID }}
	initial_password = "{{.InitialPassword}}"
	power_state  = "{{.PowerState}}"
	resource_profile  = "{{.ResourceProfile}}"
	storage_policy  = "{{.StoragePolicy}}"
	subtenant_id  = {{ StringTFValue .SubtenantID}}
	{{range .VirtualNetworkDevices}}
	virtual_network_devices {
		virtual_network_id = {{ StringTFValue .VirtualNetworkID }}
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
	
	{{if .EncryptionKeyID }}
	encryption_key_id = {{ StringTFValue .EncryptionKeyID }}
	{{end}}
	{{if .EncryptionPrivateKeyIDToUnwrap }}
	encryption_private_key_id_to_unwrap = {{ StringTFValue .EncryptionPrivateKeyIDToUnwrap }}
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
	subtenant1 := SubtenantDataSourceTestData{
		ResourceName: "subtenant1",
		Name:         testData.SubtenantForVMName,
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
		ResourceName:    "default",
		DisplayName:     generateShortRandName(),
		DeploymentID:    testDataResourceID(&deployment),
		InitialPassword: "50b90880f9f",
		PowerState:      "poweredOn",
		ResourceProfile: "SIZE_S",
		StoragePolicy:   "STANDARD",
		SubtenantID:     testDataResourceID(&subtenant1),
		VirtualNetworkDevices: []struct{ VirtualNetworkID string }{
			{VirtualNetworkID: testDataResourceID(&vnet1)},
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
	}

	configInitial := deployment.ToString() + subtenant1.ToString() + vnet1.ToString() + virtualMachine.ToString()

	virtualMachineUpdated := virtualMachine
	virtualMachineUpdated.DisplayName += "upd"
	virtualMachineUpdated.StoragePolicy = "ENTERPRISE"
	virtualMachineUpdated.ResourceProfile = "SIZE_M"
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
	configUpdated := deployment.ToString() + subtenant1.ToString() + vnet1.ToString() + virtualMachineUpdated.ToString()

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
					resource.TestCheckResourceAttr(resourceName, "resource_profile", virtualMachine.ResourceProfile),
					resource.TestCheckResourceAttr(resourceName, "storage_policy", virtualMachine.StoragePolicy),
					resource.TestCheckResourceAttrPair(resourceName, "subtenant_id", subtenant1.FullResourceName(), "id"),
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
				Config: configUpdated,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", virtualMachineUpdated.DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "deployment_id", deployment.FullResourceName(), "id"),
					resource.TestCheckResourceAttr(resourceName, "power_state", virtualMachineUpdated.PowerState),
					resource.TestCheckResourceAttr(resourceName, "resource_profile", virtualMachineUpdated.ResourceProfile),
					resource.TestCheckResourceAttr(resourceName, "storage_policy", virtualMachineUpdated.StoragePolicy),
					resource.TestCheckResourceAttrPair(resourceName, "subtenant_id", subtenant1.FullResourceName(), "id"),
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
	subtenant1 := SubtenantDataSourceTestData{
		ResourceName: "subtenant1",
		Name:         testData.SubtenantForVMName,
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
		DisplayName:  generateShortRandName(),
		KeyUsage:     []string{"ENCRYPT", "DECRYPT"},
		Algorithm:    "AES",
		Size:         256,
	}

	virtualMachine := VirtualMachineTestData{
		ResourceName:    "default",
		DisplayName:     generateShortRandName(),
		DeploymentID:    testDataResourceID(&deployment),
		InitialPassword: "50b90880f9f",
		PowerState:      "poweredOn",
		ResourceProfile: "SIZE_S",
		StoragePolicy:   "STANDARD",
		SubtenantID:     testDataResourceID(&subtenant1),
		VirtualNetworkDevices: []struct{ VirtualNetworkID string }{
			{VirtualNetworkID: testDataResourceID(&vnet1)},
		},
		Encryption: true,
	}

	configInitial := deployment.ToString() + subtenant1.ToString() + vnet1.ToString() + kmsAESKey.ToString() + virtualMachine.ToString()

	virtualMachineUpdated := virtualMachine
	virtualMachineUpdated.EncryptionKeyID = testDataResourceID(&kmsAESKey)
	virtualMachineUpdated.EncryptionRecrypt = "SHALLOW"

	configUpdated := deployment.ToString() + subtenant1.ToString() + vnet1.ToString() + kmsAESKey.ToString() + virtualMachineUpdated.ToString()

	resourceName := virtualMachine.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,

		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check:  virtualMachineChecks(resourceName, virtualMachine, subtenant1, deployment, vnet1),
			},
			{
				Config: configUpdated,
				Check:  virtualMachineChecks(resourceName, virtualMachineUpdated, subtenant1, deployment, vnet1),
			},
		},
		CheckDestroy: testAccVirtualMachineResourceNotExists(virtualMachine.DisplayName),
	})
}

func TestAccVirtualMachineResource_create_with_own_encryption(t *testing.T) {
	subtenant1 := SubtenantDataSourceTestData{
		ResourceName: "subtenant1",
		Name:         testData.SubtenantForVMName,
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
		DisplayName:  generateShortRandName(),
		KeyUsage:     []string{"ENCRYPT", "DECRYPT"},
		Algorithm:    "AES",
		Size:         256,
	}

	virtualMachine := VirtualMachineTestData{
		ResourceName:    "default",
		DisplayName:     generateShortRandName(),
		DeploymentID:    testDataResourceID(&deployment),
		InitialPassword: "50b90880f9f",
		PowerState:      "poweredOn",
		ResourceProfile: "SIZE_S",
		StoragePolicy:   "STANDARD",
		SubtenantID:     testDataResourceID(&subtenant1),
		VirtualNetworkDevices: []struct{ VirtualNetworkID string }{
			{VirtualNetworkID: testDataResourceID(&vnet1)},
		},
		Encryption:      true,
		EncryptionKeyID: testDataResourceID(&kmsAESKey),
	}

	configInitial := deployment.ToString() + subtenant1.ToString() + vnet1.ToString() + kmsAESKey.ToString() + virtualMachine.ToString()

	virtualMachineUpdated := virtualMachine
	virtualMachineUpdated.EncryptionKeyID = ""
	virtualMachineUpdated.EncryptionRecrypt = "SHALLOW"

	configUpdated := deployment.ToString() + subtenant1.ToString() + vnet1.ToString() + kmsAESKey.ToString() + virtualMachineUpdated.ToString()

	resourceName := virtualMachine.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,

		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check:  virtualMachineChecks(resourceName, virtualMachine, subtenant1, deployment, vnet1),
			},
			{
				Config: configUpdated,
				Check:  virtualMachineChecks(resourceName, virtualMachineUpdated, subtenant1, deployment, vnet1),
			},
		},
		CheckDestroy: testAccVirtualMachineResourceNotExists(virtualMachine.DisplayName),
	})
}
func virtualMachineChecks(resourceName string, vm VirtualMachineTestData, subtenant SubtenantDataSourceTestData, deployment DeploymentDataSourceTestData, vnet VirtualNetworkDataSourceTestData) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "display_name", vm.DisplayName),
		resource.TestCheckResourceAttrPair(resourceName, "deployment_id", deployment.FullResourceName(), "id"),
		resource.TestCheckResourceAttr(resourceName, "power_state", vm.PowerState),
		resource.TestCheckResourceAttr(resourceName, "resource_profile", vm.ResourceProfile),
		resource.TestCheckResourceAttr(resourceName, "storage_policy", vm.StoragePolicy),
		resource.TestCheckResourceAttrPair(resourceName, "subtenant_id", subtenant.FullResourceName(), "id"),
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
