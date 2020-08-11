package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func mapSliceToInterfaceSlice(in []map[string]interface{}) []interface{} {
	if in == nil {
		return nil
	}

	output := make([]interface{}, len(in))
	for i := range in {
		output[i] = in[i]
	}

	return output
}

func mapToInterfaceSlice(in map[string]interface{}) []interface{} {
	if in == nil {
		return nil
	}

	output := make([]interface{}, 1)
	output[0] = in

	return output
}

func mapToMapSlice(in map[string]interface{}) []map[string]interface{} {
	if in == nil {
		return nil
	}

	output := make([]map[string]interface{}, 1)
	output = append(output, in)

	return output
}

func generateRandName() string {
	return fmt.Sprintf("tf-acc_test-%s", acctest.RandStringFromCharSet(8, acctest.CharSetAlphaNum))
}
