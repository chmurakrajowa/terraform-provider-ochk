package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type RouterTestData struct {
	ResourceName string
	DisplayName  string
	ParentRouter string
	ProjectName  string
}

func (c *RouterTestData) ToString(routerDataSourceName string) string {
	return executeTemplateToString(`
data "ochk_project" "project-`+routerDataSourceName+`" {
  display_name = "{{ .ProjectName}}"
}

data "ochk_vrf" "vrf-`+routerDataSourceName+`" {
  display_name = "{{ .ParentRouter}}"
}

resource "ochk_vpc" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  vrf_id = data.ochk_vrf.vrf-`+routerDataSourceName+`.id
  project_id = data.ochk_project.project-`+routerDataSourceName+`.id
}
`, c)
}

func (c *RouterTestData) FullResourceName() string {
	return "ochk_vpc." + c.ResourceName
}

func TestAccRouterResource_create_update(t *testing.T) {

	router := RouterTestData{
		ResourceName: "vpc",
		DisplayName:  generateShortRandName(devTestDataPrefix),
		ParentRouter: testData.VRF,
		ProjectName:  testData.Project1Name,
	}

	routerUpdated := RouterTestData{
		ResourceName: "vpc",
		DisplayName:  router.DisplayName + "-upd",
		ParentRouter: testData.VRF,
		ProjectName:  testData.Project1Name,
	}

	RouterResourceName := router.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: router.ToString("r-test1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(RouterResourceName, "display_name", router.DisplayName),
					resource.TestCheckResourceAttrPair(RouterResourceName, "vrf_id", "data.ochk_vrf.vrf-r-test1", "id"),
					resource.TestCheckResourceAttrPair(RouterResourceName, "project_id", "data.ochk_project.project-r-test1", "id"),
					resource.TestCheckResourceAttr(RouterResourceName, "folder_path", "/"),
					resource.TestCheckResourceAttrSet(RouterResourceName, "created_by"),
					resource.TestCheckResourceAttrSet(RouterResourceName, "created_at"),
					resource.TestCheckResourceAttrSet(RouterResourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(RouterResourceName, "modified_at"),
				),
			},
			{
				ResourceName:      RouterResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: routerUpdated.ToString("r-test1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(RouterResourceName, "display_name", routerUpdated.DisplayName),
					resource.TestCheckResourceAttrPair(RouterResourceName, "project_id", "data.ochk_project.project-r-test1", "id"),
				),
			},
		},
		CheckDestroy: testAccRouterResourceNotExists(routerUpdated.DisplayName),
	})
}

func testAccRouterResourceNotExists(displayName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).Routers

		Routers, err := proxy.ListByDisplayName(context.Background(), displayName)
		if err != nil {
			return err
		}

		if len(Routers) > 0 {
			return fmt.Errorf("vpc %s still exists", Routers[0].RouterID)
		}

		return nil
	}
}
