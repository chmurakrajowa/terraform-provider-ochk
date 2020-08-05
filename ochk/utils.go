package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func toInterfaceSlice(in []map[string]interface{}) []interface{} {
	output := make([]interface{}, len(in))
	for i := range in {
		output[i] = in[i]
	}

	return output
}

func generateRandName() string {
	return fmt.Sprintf("tf-acc_test-%s", acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))
}
