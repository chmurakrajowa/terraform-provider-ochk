package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccProjectDataSource_read(t *testing.T) {

	ctx := context.Background()
	client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "", os.Getenv("TF_VAR_platform_type"))
	if err != nil {
		assert.Error(t, err)
	}
	fmt.Printf("Project name: %v\n", client.Projects)
	proxy := client.Projects
	projectInstance, err := proxy.Read(context.Background(), "3cda830c-b37f-46dc-be54-f649d31bec66")
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Condition(t, func() bool {
			if projectInstance.Name == "tf-gojl-project-01" {
				return true
			}
			return false
		})
	}

}
