package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccCustomServiceDataSource_read(t *testing.T) {
	resourceName := "data.ochk_custom_service.http"
	dataSourceProject := "data.ochk_project.project"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCustomServiceDataSourceConfig(testData.CustomService1DisplayName, testData.Project1Name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_name", testData.CustomService1DisplayName),
					resource.TestCheckResourceAttrPair(resourceName, "project_id", dataSourceProject, "id"),
					resource.TestCheckResourceAttr(resourceName, "ports.0.protocol", "TCP"),
					testStringInSet(resourceName, "ports.0.source", "443"),
					testStringInSet(resourceName, "ports.0.destination", "1-65535"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
				),
			},
		},
	})
}

func testAccCustomServiceDataSourceConfig(displayName string, projectName string) string {
	return fmt.Sprintf(`

data "ochk_project" "project" {
  display_name = %[2]q
}

data "ochk_custom_service" "http" {
  display_name = %[1]q
}
`, displayName, projectName)
}
