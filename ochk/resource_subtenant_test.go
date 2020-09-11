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
	network1 := NetworkDataSourceTestData{
		ResourceName: "network1",
		Name:         testDataNetwork1Name,
	}

	user1 := UserDataSourceTestData{
		ResourceName: "user1",
		Name:         "devel-dstawicki",
	}

	subtenant := SubtenantTestData{
		ResourceName:          "default",
		Description:           "tf-test-description",
		Email:                 "test@example.com",
		MemoryReservedSizeMB:  24000,
		Name:                  "tf-test-name-" + generateRandName(),
		NetworkIDs:            []string{network1.FullResourceName() + ".id"},
		StorageReservedSizeGB: 150,
		UserIDs:               []string{user1.FullResourceName() + ".id"},
	}

	configInitial := user1.ToString() + network1.ToString() + subtenant.ToString()

	network2 := NetworkDataSourceTestData{
		ResourceName: "network2",
		Name:         testDataNetwork2Name,
	}

	user2 := UserDataSourceTestData{
		ResourceName: "user2",
		Name:         "devel-firstuserpb1",
	}

	subtenantUpdated := subtenant
	subtenantUpdated.MemoryReservedSizeMB = 30000
	subtenantUpdated.StorageReservedSizeGB = 200
	subtenantUpdated.NetworkIDs = []string{network2.FullResourceName() + ".id"}
	subtenantUpdated.UserIDs = []string{user2.FullResourceName() + ".id"}
	subtenantUpdated.Email = "email.updated@example.com"
	subtenantUpdated.Description += "- updated"

	configUpdated := user2.ToString() + network2.ToString() + subtenantUpdated.ToString()

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
					resource.TestCheckResourceAttrPair(subtenantResourceName, "users.0", user1.FullResourceName(), "id"),
					resource.TestCheckResourceAttrPair(subtenantResourceName, "networks.0", network1.FullResourceName(), "id"),
				),
			},
			{
				Config: configUpdated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(subtenantResourceName, "name", subtenantUpdated.Name),
					//FIXME waiting for backend fix
					//resource.TestCheckResourceAttr(subtenantResourceName, "description", subtenantUpdated.Description),
					resource.TestCheckResourceAttr(subtenantResourceName, "email", subtenantUpdated.Email),
					resource.TestCheckResourceAttr(subtenantResourceName, "memory_reserved_size_mb", fmt.Sprintf("%d", subtenantUpdated.MemoryReservedSizeMB)),
					resource.TestCheckResourceAttr(subtenantResourceName, "storage_reserved_size_gb", fmt.Sprintf("%d", subtenantUpdated.StorageReservedSizeGB)),
					resource.TestCheckResourceAttrPair(subtenantResourceName, "users.0", user2.FullResourceName(), "id"),
					resource.TestCheckResourceAttrPair(subtenantResourceName, "networks.0", network2.FullResourceName(), "id"),
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
