package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strconv"
	"testing"
)

type ProjectDataSourceTestData struct {
	ResourceName string
	DisplayName  string
}

func (c *ProjectDataSourceTestData) ToString() string {
	return ochk.CastTemplateToString(`
data "ochk_project" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
}
`, c)
}

func (c *ProjectDataSourceTestData) FullResourceName() string {
	return "data.ochk_project." + c.ResourceName
}

func TestAccProjectDataSource_read(t *testing.T) {
	project := &ProjectDataSourceTestData{
		ResourceName: "default",
		DisplayName:  ochk.TestOpenstackData.Project1Name,
	}

	config := project.ToString()

	resourceName := project.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: ochk.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", project.DisplayName),
					resource.TestCheckResourceAttrSet(resourceName, "description"),
					resource.TestCheckResourceAttr(resourceName, "limits_enabled", strconv.FormatBool(true)),
					resource.TestCheckResourceAttr(resourceName, "memory_reserved_size_mb", "1000"),
					resource.TestCheckResourceAttr(resourceName, "storage_reserved_size_gb", "1000"),
					resource.TestCheckResourceAttr(resourceName, "vcpu_reserved_quantity", "50"),
				),
			},
		},
	})
}
