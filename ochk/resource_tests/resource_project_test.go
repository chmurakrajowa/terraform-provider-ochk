package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strconv"
	"testing"
)

type ProjectTestData struct {
	ResourceName          string
	Description           string
	MemoryReservedSizeMB  int64
	DisplayName           string
	StorageReservedSizeGB int64
	VcpuReservedQuantity  int64
	ParentRouter          string
	LimitsEnabled         bool
}

func (c *ProjectTestData) ToString() string {
	return ochk.CastTemplateToString(`
data "ochk_vrf" "vrf" {
  display_name = "{{ .ParentRouter}}"
}

resource "ochk_project" "{{ .ResourceName}}" {
	display_name = "{{.DisplayName}}"
	description = "{{.Description}}"
	limits_enabled = {{.LimitsEnabled}}
	memory_reserved_size_mb = "{{.MemoryReservedSizeMB}}"
	storage_reserved_size_gb = "{{.StorageReservedSizeGB}}"
	vcpu_reserved_quantity = "{{.VcpuReservedQuantity}}"
    vrf_id = data.ochk_vrf.vrf.id
}
`, c)
}

func (c *ProjectTestData) FullResourceName() string {
	return "ochk_project." + c.ResourceName
}

func TestAccProjectResource_create(t *testing.T) {
	fmt.Printf("Project ochk.TestOpenstackData.VRF %v\n", ochk.TestOpenstackData)

	project := ProjectTestData{
		ResourceName:          "default",
		Description:           "tf-test-description",
		ParentRouter:          ochk.TestOpenstackData.VRF,
		MemoryReservedSizeMB:  23,
		DisplayName:           ochk.GenerateRandName(ochk.DevTestDataOpenstackPrefix),
		StorageReservedSizeGB: 150,
		VcpuReservedQuantity:  100,
		LimitsEnabled:         true,
	}

	configInitial := project.ToString()
	fmt.Printf("Project terraform body: %v\n", configInitial)

	fmt.Printf("Project full name: %v\n", project.DisplayName)
	fmt.Printf("Project memeory GB: %v\n", project.MemoryReservedSizeMB)
	projectUpdated := project
	projectUpdated.MemoryReservedSizeMB = 25
	projectUpdated.StorageReservedSizeGB = 200
	projectUpdated.VcpuReservedQuantity = 100
	projectUpdated.Description += " - updated"

	configUpdated := projectUpdated.ToString()

	projectUpdatedWithRecreate := projectUpdated

	configUpdatedWithRecreate := projectUpdatedWithRecreate.ToString()

	projectResourceName := project.FullResourceName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: ochk.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(projectResourceName, "display_name", project.DisplayName),
					resource.TestCheckResourceAttr(projectResourceName, "description", project.Description),
					resource.TestCheckResourceAttr(projectResourceName, "limits_enabled", strconv.FormatBool(project.LimitsEnabled)),
					resource.TestCheckResourceAttr(projectResourceName, "memory_reserved_size_mb", fmt.Sprintf("%d", project.MemoryReservedSizeMB)),
					resource.TestCheckResourceAttr(projectResourceName, "storage_reserved_size_gb", fmt.Sprintf("%d", project.StorageReservedSizeGB)),
					resource.TestCheckResourceAttr(projectResourceName, "vcpu_reserved_quantity", fmt.Sprintf("%d", project.VcpuReservedQuantity)),
				),
			},
			{
				ResourceName:      projectResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: configUpdated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(projectResourceName, "display_name", projectUpdated.DisplayName),
					resource.TestCheckResourceAttr(projectResourceName, "description", projectUpdated.Description),
					resource.TestCheckResourceAttr(projectResourceName, "memory_reserved_size_mb", fmt.Sprintf("%d", projectUpdated.MemoryReservedSizeMB)),
					resource.TestCheckResourceAttr(projectResourceName, "storage_reserved_size_gb", fmt.Sprintf("%d", projectUpdated.StorageReservedSizeGB)),
					resource.TestCheckResourceAttr(projectResourceName, "vcpu_reserved_quantity", fmt.Sprintf("%d", projectUpdated.VcpuReservedQuantity)),
				),
			},
			{
				Config: configUpdatedWithRecreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(projectResourceName, "display_name", projectUpdatedWithRecreate.DisplayName),
					resource.TestCheckResourceAttr(projectResourceName, "description", projectUpdated.Description),
					resource.TestCheckResourceAttr(projectResourceName, "memory_reserved_size_mb", fmt.Sprintf("%d", projectUpdatedWithRecreate.MemoryReservedSizeMB)),
					resource.TestCheckResourceAttr(projectResourceName, "storage_reserved_size_gb", fmt.Sprintf("%d", projectUpdatedWithRecreate.StorageReservedSizeGB)),
					resource.TestCheckResourceAttr(projectResourceName, "vcpu_reserved_quantity", fmt.Sprintf("%d", projectUpdatedWithRecreate.VcpuReservedQuantity)),
				),
			},
		},
		CheckDestroy: testAccProjectResourceNotExists(project.DisplayName),
	})
}

func testAccProjectResourceNotExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := ochk.TestAccProviderDefinition.Meta().(*sdk.Client).Projects

		projects, err := proxy.ListByName(context.Background(), name)
		if err != nil {
			return err
		}

		if len(projects) > 0 {
			return fmt.Errorf("project %s still exists", projects[0].ProjectID)
		}

		return nil
	}
}
