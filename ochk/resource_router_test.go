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
}

func (c *RouterTestData) ToString(routerDataSourceName string) string {
	return executeTemplateToString(`
data "ochk_router" "vrf-`+routerDataSourceName+`" {
  display_name = "{{ .ParentRouter}}"
}

resource "ochk_router" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  parent_router_id = data.ochk_router.vrf-`+routerDataSourceName+`.id
}
`, c)
}

func (c *RouterTestData) FullResourceName() string {
	return "ochk_router." + c.ResourceName
}

func TestAccRouterResource_create_update(t *testing.T) {
	router := RouterTestData{
		ResourceName: "router",
		DisplayName:  generateShortRandName(devTestDataPrefix),
		ParentRouter: testData.VRF,
	}

	routerUpdated := RouterTestData{
		ResourceName: "router",
		DisplayName:  router.DisplayName + "-upd",
		ParentRouter: testData.VRF,
	}

	RouterResourceName := router.FullResourceName()
	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: router.ToString("r-test1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(RouterResourceName, "display_name", router.DisplayName),
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
			return fmt.Errorf("router %s still exists", Routers[0].RouterID)
		}

		return nil
	}
}
