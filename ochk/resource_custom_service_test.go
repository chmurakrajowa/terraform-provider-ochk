package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

type CustomServiceTestData struct {
	ResourceName string
	DisplayName  string
	Ports        []CustomServicePortTestData
	ProjectID    string
}

type CustomServicePortTestData struct {
	Protocol    string
	Source      []string
	Destination []string
}

func (c *CustomServiceTestData) ToString() string {
	return executeTemplateToString(`

data "ochk_project" "project-cs-1" {
  display_name = "`+testData.Project1Name+`"
}

resource "ochk_custom_service" "{{ .ResourceName}}" {
  display_name = "{{ .DisplayName}}"
  project_id = data.ochk_project.project-cs-1.id

  {{range $port := .Ports}}
  ports {
    protocol = "{{ $port.Protocol }}"
    source = {{ StringsToTFList $port.Source }}
    destination = {{ StringsToTFList $port.Destination }}
  }
  {{end}}
}
`, c)
}

func (c *CustomServiceTestData) FullResourceName() string {
	return "ochk_custom_service." + c.ResourceName
}

func TestAccCustomServiceResource_create(t *testing.T) {
	customService := CustomServiceTestData{
		ResourceName: "default",
		DisplayName:  generateShortRandName(devTestDataPrefix),
		Ports: []CustomServicePortTestData{
			{
				Protocol:    "TCP",
				Source:      []string{"8080", "8081"},
				Destination: []string{"80", "443"},
			},
			{
				Protocol:    "UDP",
				Source:      []string{"53-54"},
				Destination: []string{"8053-8054"},
			},
		},
	}

	customServiceUpdated := customService
	customServiceUpdated.Ports = []CustomServicePortTestData{
		customService.Ports[0],
		customService.Ports[1],
		{
			Protocol:    "TCP",
			Source:      []string{"80"},
			Destination: []string{"81"},
		},
	}

	customServiceUpdated.Ports[0].Source = []string{"8080", "8081", "8082"}
	customServiceUpdated.Ports[0].Destination = []string{"80"}

	customServiceResourceName := customService.FullResourceName()

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: customService.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(customServiceResourceName, "display_name", customService.DisplayName),
					resource.TestCheckResourceAttrPair(customServiceResourceName, "project_id", "data.ochk_project.project-cs-1", "id"),
					resource.TestCheckResourceAttrSet(customServiceResourceName, "ports.0.id"),
					resource.TestCheckResourceAttr(customServiceResourceName, "ports.0.protocol", customService.Ports[0].Protocol),
					testStringInSet(customServiceResourceName, "ports.0.source", customService.Ports[0].Source[0]),
					testStringInSet(customServiceResourceName, "ports.0.source", customService.Ports[0].Source[1]),
					testStringInSet(customServiceResourceName, "ports.0.destination", customService.Ports[0].Destination[0]),
					testStringInSet(customServiceResourceName, "ports.0.destination", customService.Ports[0].Destination[1]),
					resource.TestCheckResourceAttrSet(customServiceResourceName, "ports.1.id"),
					resource.TestCheckResourceAttr(customServiceResourceName, "ports.1.protocol", customService.Ports[1].Protocol),
					testStringInSet(customServiceResourceName, "ports.1.source", customService.Ports[1].Source[0]),
					testStringInSet(customServiceResourceName, "ports.1.destination", customService.Ports[1].Destination[0]),
					resource.TestCheckResourceAttrSet(customServiceResourceName, "created_by"),
					resource.TestCheckResourceAttrSet(customServiceResourceName, "created_at"),
					resource.TestCheckResourceAttrSet(customServiceResourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(customServiceResourceName, "modified_at"),
				),
			},
			{
				ResourceName:      customServiceResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: customServiceUpdated.ToString(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(customServiceResourceName, "display_name", customServiceUpdated.DisplayName),
					resource.TestCheckResourceAttrPair(customServiceResourceName, "project_id", "data.ochk_project.project-cs-1", "id"),
					resource.TestCheckResourceAttrSet(customServiceResourceName, "ports.0.id"),
					resource.TestCheckResourceAttr(customServiceResourceName, "ports.0.protocol", customServiceUpdated.Ports[0].Protocol),
					testStringInSet(customServiceResourceName, "ports.0.source", customServiceUpdated.Ports[0].Source[0]),
					testStringInSet(customServiceResourceName, "ports.0.source", customServiceUpdated.Ports[0].Source[1]),
					testStringInSet(customServiceResourceName, "ports.0.source", customServiceUpdated.Ports[0].Source[2]),
					testStringInSet(customServiceResourceName, "ports.0.destination", customServiceUpdated.Ports[0].Destination[0]),
					resource.TestCheckResourceAttrSet(customServiceResourceName, "ports.1.id"),
					resource.TestCheckResourceAttr(customServiceResourceName, "ports.1.protocol", customServiceUpdated.Ports[1].Protocol),
					testStringInSet(customServiceResourceName, "ports.1.source", customServiceUpdated.Ports[1].Source[0]),
					testStringInSet(customServiceResourceName, "ports.1.destination", customServiceUpdated.Ports[1].Destination[0]),
					resource.TestCheckResourceAttrSet(customServiceResourceName, "ports.2.id"),
					resource.TestCheckResourceAttr(customServiceResourceName, "ports.2.protocol", customServiceUpdated.Ports[2].Protocol),
					testStringInSet(customServiceResourceName, "ports.2.source", customServiceUpdated.Ports[2].Source[0]),
					testStringInSet(customServiceResourceName, "ports.2.destination", customServiceUpdated.Ports[2].Destination[0]),
					resource.TestCheckResourceAttrSet(customServiceResourceName, "created_by"),
					resource.TestCheckResourceAttrSet(customServiceResourceName, "created_at"),
					resource.TestCheckResourceAttrSet(customServiceResourceName, "modified_by"),
					resource.TestCheckResourceAttrSet(customServiceResourceName, "modified_at"),
				),
			},
		},
		CheckDestroy: testAccCustomServiceResourceNotExists(customServiceUpdated.DisplayName),
	})
}

func testAccCustomServiceResourceNotExists(displayNames ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		proxy := testAccProvider.Meta().(*sdk.Client).CustomServices

		for i := 0; i < len(displayNames); i++ {
			customServices, err := proxy.ListByDisplayName(context.Background(), displayNames[i])
			if err != nil {
				return err
			}

			if len(customServices) > 0 {
				return fmt.Errorf("custom service %s still exists", customServices[0].ServiceID)
			}
		}

		return nil
	}
}
