package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type SubtenantTestData struct {
	ResourceName          string
	Description           string
	Email                 string
	MemoryReservedSizeMB  int64
	Name                  string
	StorageReservedSizeGB int64
}

func (c *SubtenantTestData) ToString() string {
	return executeTemplateToString(`
resource "ochk_subtenant" "{{ .ResourceName}}" {
	name = "{{.Name}}"
	email = "{{.Email}}"
	description = "{{.Description}}"
	memory_reserved_size_mb = "{{.MemoryReservedSizeMB}}"
	storage_reserved_size_gb = "{{.StorageReservedSizeGB}}"

	lifecycle {
		ignore_changes = [description]
	}
}
`, c)
}

func (c *SubtenantTestData) FullResourceName() string {
	return "ochk_subtenant." + c.ResourceName
}

func TestAccSubtenantResource_create(t *testing.T) {

	subtenant := SubtenantTestData{
		ResourceName:          "default",
		Description:           "tf-test-description",
		Email:                 "test@example.com",
		MemoryReservedSizeMB:  24000,
		Name:                  generateRandName(devTestDataPrefix),
		StorageReservedSizeGB: 150,
	}

	configInitial := subtenant.ToString()

	subtenantUpdated := subtenant
	subtenantUpdated.MemoryReservedSizeMB = 30000
	subtenantUpdated.StorageReservedSizeGB = 200
	subtenantUpdated.Description += "- updated"

	configUpdated := subtenantUpdated.ToString()

	subtenantUpdatedWithRecreate := subtenantUpdated
	subtenantUpdatedWithRecreate.Email = "email.updated@example.com"

	configUpdatedWithRecreate := subtenantUpdatedWithRecreate.ToString()

	subtenantResourceName := subtenant.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: configInitial,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(subtenantResourceName, "name", subtenant.Name),
					//FIXME waiting for backend fix
					//resource.TestCheckResourceAttr(subtenantResourceName, "description", subtenant.Description),
					resource.TestCheckResourceAttr(subtenantResourceName, "email", subtenant.Email),
					resource.TestCheckResourceAttr(subtenantResourceName, "memory_reserved_size_mb", fmt.Sprintf("%d", subtenant.MemoryReservedSizeMB)),
					resource.TestCheckResourceAttr(subtenantResourceName, "storage_reserved_size_gb", fmt.Sprintf("%d", subtenant.StorageReservedSizeGB)),
				),
			},
			{
				ResourceName:      subtenantResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: configUpdated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(subtenantResourceName, "name", subtenantUpdated.Name),
					//resource.TestCheckResourceAttr(subtenantResourceName, "description", subtenantUpdated.Description),
					resource.TestCheckResourceAttr(subtenantResourceName, "email", subtenantUpdated.Email),
					resource.TestCheckResourceAttr(subtenantResourceName, "memory_reserved_size_mb", fmt.Sprintf("%d", subtenantUpdated.MemoryReservedSizeMB)),
					resource.TestCheckResourceAttr(subtenantResourceName, "storage_reserved_size_gb", fmt.Sprintf("%d", subtenantUpdated.StorageReservedSizeGB)),
				),
			},
			{
				Config: configUpdatedWithRecreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(subtenantResourceName, "name", subtenantUpdatedWithRecreate.Name),
					//resource.TestCheckResourceAttr(subtenantResourceName, "description", subtenantUpdated.Description),
					resource.TestCheckResourceAttr(subtenantResourceName, "email", subtenantUpdatedWithRecreate.Email),
					resource.TestCheckResourceAttr(subtenantResourceName, "memory_reserved_size_mb", fmt.Sprintf("%d", subtenantUpdatedWithRecreate.MemoryReservedSizeMB)),
					resource.TestCheckResourceAttr(subtenantResourceName, "storage_reserved_size_gb", fmt.Sprintf("%d", subtenantUpdatedWithRecreate.StorageReservedSizeGB)),
				),
			},
		},
		CheckDestroy: testAccSubtenantResourceNotExists(subtenant.Name),
	})
}

func testAccSubtenantResourceNotExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).Subtenants

		subtenants, err := proxy.ListByName(context.Background(), name)
		if err != nil {
			return err
		}

		if len(subtenants) > 0 {
			return fmt.Errorf("subtenant %s still exists", subtenants[0].SubtenantID)
		}

		return nil
	}
}
