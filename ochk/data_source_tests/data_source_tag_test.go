package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccTagDataSource_read(t *testing.T) {

	ctx := context.Background()
	client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "", os.Getenv("TF_VAR_platform_type"))
	if err != nil {
		assert.Error(t, err)
	}
	fmt.Printf("Tag name: %v\n", client.Tags)
	proxy := client.Tags
	tagInstance, err := proxy.Read(context.Background(), 354)
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Condition(t, func() bool {
			if tagInstance.TagValue == "tf-gojl-fvs8" {
				return true
			}
			return false
		})
	}

}
