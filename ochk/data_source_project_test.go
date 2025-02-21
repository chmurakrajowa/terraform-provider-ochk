package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strconv"
	"testing"
)

type ProjectDataSourceTestData struct {
	ResourceName string
	DisplayName  string
}

func (c *ProjectDataSourceTestData) ToString() string {
	return executeTemplateToString(`
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
		DisplayName:  testData.Project1Name,
	}

	config := project.ToString()
	platformType := checkPlatformType()
	fmt.Printf("checkPlatformType%s ", platformType)
	memory_reserved_size_mb := "1000"
	storage_reserved_size_gb := "1000"
	vcpu_reserved_quantity := "50"
	if platformType == "VMWARE" {
		memory_reserved_size_mb = "255000"
		storage_reserved_size_gb = "10000"
		vcpu_reserved_quantity = "50"
	}
	resourceName := project.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", project.DisplayName),
					resource.TestCheckResourceAttrSet(resourceName, "description"),
					resource.TestCheckResourceAttr(resourceName, "limits_enabled", strconv.FormatBool(true)),
					resource.TestCheckResourceAttr(resourceName, "memory_reserved_size_mb", memory_reserved_size_mb),
					resource.TestCheckResourceAttr(resourceName, "storage_reserved_size_gb", storage_reserved_size_gb),
					resource.TestCheckResourceAttr(resourceName, "vcpu_reserved_quantity", vcpu_reserved_quantity),
				),
			},
		},
	})
}
