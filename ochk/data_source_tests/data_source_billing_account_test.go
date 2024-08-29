package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAccDataSource_read(t *testing.T) {

	ctx := context.Background()
	client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "")
	if err != nil {
		assert.Error(t, err)
	}
	fmt.Printf("Account name: %v\n", client.Accounts)
	proxy := client.Accounts
	accountInstance, err := proxy.Read(context.Background(), "ec6db72b-0de8-47f3-9bd3-b4c9f387a036")
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Condition(t, func() bool {
			if accountInstance.AccountName == "Rachunek Platformowy" {
				return true
			}
			return false
		})
	}
}
