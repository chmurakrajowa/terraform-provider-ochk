package ochk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAcctsDataSource_read(t *testing.T) {
	ctx := context.Background()
	client, err := sdk.NewClient(ctx, os.Getenv("TF_VAR_host"), os.Getenv("TF_VAR_platform"), os.Getenv("TF_VAR_api_key"), false, "")
	if err != nil {
		assert.Error(t, err)
	}
	fmt.Printf("Account name: %v\n", client.Accounts)
	proxy := client.Accounts
	accountInstances, err := proxy.ListAccounts(context.Background())
	if err != nil {
		assert.Error(t, err)
	} else {
		assert.Condition(t, func() bool {
			if accountInstances != nil {
				return true
			}
			return false
		})
		assert.Condition(t, func() bool {
			if accountInstances != nil && len(accountInstances) >= 0 {

				return true
			}
			return false
		})
	}

}
