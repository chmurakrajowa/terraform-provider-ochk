package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccSecurityGroupDataSource_read(t *testing.T) {
	checkVariables := CheckInputVariables()
	if checkVariables == "" {
		ctx := context.Background()
		client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "", os.Getenv("TF_VAR_platform_type"))
		if err != nil {
			assert.Error(t, err)
		}
		fmt.Printf("Security group name: %v\n", client.SecurityGroups)
		proxy := client.SecurityGroups
		securityGroupInstance, err := proxy.Read(context.Background(), "1f59e0d4-94c2-4944-b1b0-62a378dabcb3")
		if err != nil {
			assert.Error(t, err)
		} else {
			assert.Condition(t, func() bool {
				if securityGroupInstance.DisplayName == "admin_gr00001" {
					return true
				}
				return false
			})
		}

		/*
			resourceName := "data.ochk_security_group.one_member"
			displayName := generateRandName(devTestDataPrefix)

			resource.ParallelTest(t, resource.TestCase{
				ProviderFactories: testAccProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccSecurityGroupDataSourceConfig(displayName),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(resourceName, "display_name", displayName),
							resource.TestCheckResourceAttrPair(resourceName, "project_id", "data.ochk_project.project-sg-1", "id"),
							resource.TestCheckResourceAttrSet(resourceName, "created_by"),
							resource.TestCheckResourceAttrSet(resourceName, "created_at"),
							resource.TestCheckResourceAttrSet(resourceName, "modified_by"),
							resource.TestCheckResourceAttrSet(resourceName, "modified_at"),
							resource.TestCheckResourceAttrPair(resourceName, "members.0.id", "data.ochk_virtual_machine.default", "id"),
							resource.TestCheckResourceAttr(resourceName, "members.0.type", "VIRTUAL_MACHINE"),
							resource.TestCheckResourceAttrSet(resourceName, "members.0.display_name"),
						),
					},
				},
			})*/
	} else {
		fmt.Printf("ERROR:: %s", checkVariables)
	}
}
