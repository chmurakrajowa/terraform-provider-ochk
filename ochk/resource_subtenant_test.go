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
	NetworkIDs            []string
	StorageReservedSizeGB int64
	UserIDs               []string
}

func (c *SubtenantTestData) ToString() string {
	return executeTemplateToString(`
resource "ochk_subtenant" "{{ .ResourceName}}" {
	name = "{{.Name}}"
	email = "{{.Email}}"
	description = "{{.Description}}"
	memory_reserved_size_mb = "{{.MemoryReservedSizeMB}}"
	storage_reserved_size_gb = "{{.StorageReservedSizeGB}}"
	users = {{ StringsToTFList .UserIDs}}
	networks = {{ StringsToTFList .NetworkIDs}}

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
		Name:                  "tf-test-name-" + generateRandName(),
		NetworkIDs:            []string{"258313ca-365e-4f5c-8510-4a42f6595651"},
		StorageReservedSizeGB: 150,
		UserIDs:               []string{"bf5e40c6-191c-40f5-b4d1-9332a9e4ed48"},
	}

	subtenantUpdated := subtenant
	subtenantUpdated.MemoryReservedSizeMB = 30000
	subtenantUpdated.StorageReservedSizeGB = 200
	subtenantUpdated.NetworkIDs = []string{"bd814070-18f3-4182-b2af-edaa72a50fee"}
	subtenantUpdated.UserIDs = []string{"dbeb5be9-71ff-4d34-b64d-6e0fced6ed52"}
	subtenantUpdated.Email = "email.updated@example.com"
	subtenantUpdated.Description += "- updated"

	subtenantResourceName := subtenant.FullResourceName()
	resource.Test(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: subtenant.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(subtenantResourceName, "name", subtenant.Name),
					//FIXME waiting for backend fix
					//resource.TestCheckResourceAttr(subtenantResourceName, "description", subtenant.Description),
					resource.TestCheckResourceAttr(subtenantResourceName, "email", subtenant.Email),
					resource.TestCheckResourceAttr(subtenantResourceName, "memory_reserved_size_mb", fmt.Sprintf("%d", subtenant.MemoryReservedSizeMB)),
					resource.TestCheckResourceAttr(subtenantResourceName, "storage_reserved_size_gb", fmt.Sprintf("%d", subtenant.StorageReservedSizeGB)),
					resource.TestCheckResourceAttr(subtenantResourceName, "users.0", subtenant.UserIDs[0]),
					resource.TestCheckResourceAttr(subtenantResourceName, "networks.0", subtenant.NetworkIDs[0]),
				),
			},
			{
				Config: subtenantUpdated.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(subtenantResourceName, "name", subtenantUpdated.Name),
					//FIXME waiting for backend fix
					//resource.TestCheckResourceAttr(subtenantResourceName, "description", subtenantUpdated.Description),
					resource.TestCheckResourceAttr(subtenantResourceName, "email", subtenantUpdated.Email),
					resource.TestCheckResourceAttr(subtenantResourceName, "memory_reserved_size_mb", fmt.Sprintf("%d", subtenantUpdated.MemoryReservedSizeMB)),
					resource.TestCheckResourceAttr(subtenantResourceName, "storage_reserved_size_gb", fmt.Sprintf("%d", subtenantUpdated.StorageReservedSizeGB)),
					resource.TestCheckResourceAttr(subtenantResourceName, "users.0", subtenantUpdated.UserIDs[0]),
					resource.TestCheckResourceAttr(subtenantResourceName, "networks.0", subtenantUpdated.NetworkIDs[0]),
				),
			},
		},
		CheckDestroy: testAccSubtenantResourceNotExists(subtenant.Name),
	})
}

func testAccSubtenantResourceNotExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).Subtenants

		subtenants, err := proxy.ListByDisplayName(context.Background(), name)
		if err != nil {
			return err
		}

		if len(subtenants) > 0 {
			return fmt.Errorf("subtenant %s still exists", subtenants[0].SubtenantID)
		}

		return nil
	}
}
