package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

type UserDataSourceTestData struct {
	ResourceName string
	Name         string
}

func (c *UserDataSourceTestData) ToString() string {
	return executeTemplateToString(`
data "ochk_user" "{{ .ResourceName}}" {
  name = "{{ .Name}}"
}
`, c)
}

func (c *UserDataSourceTestData) FullResourceName() string {
	return "data.ochk_user." + c.ResourceName
}

func TestAccUserDataSource_read(t *testing.T) {
	user := UserDataSourceTestData{
		ResourceName: "default",
		Name:         testDataUser1Name,
	}

	resourceName := user.FullResourceName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: user.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", user.Name),
					resource.TestCheckResourceAttr(resourceName, "email_address", "test@ochk.pl"),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "first_name", "Test"),
					resource.TestCheckResourceAttr(resourceName, "last_name", "Test"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
					resource.TestCheckResourceAttr(resourceName, "user_principal_name", "devel-jpuser@under.test"),
					resource.TestCheckResourceAttr(resourceName, "principal_id", "487"),
					resource.TestCheckResourceAttr(resourceName, "principal_name", "devel-jpuser"),
					resource.TestCheckResourceAttr(resourceName, "principal_domain", "under.test"),
					resource.TestCheckResourceAttr(resourceName, "principal_type", "USER"),
				),
			},
		},
	})
}
