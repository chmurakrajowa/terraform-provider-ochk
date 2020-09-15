package ochk

import (
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

func mapInterfaceSliceToUUIDSlice(in []interface{}) []strfmt.UUID {
	if in == nil {
		return nil
	}

	output := make([]strfmt.UUID, len(in))
	for i := range in {
		output[i] = strfmt.UUID(in[i].(string))
	}

	return output
}

func transformSetToStringSlice(set *schema.Set) []string {
	list := set.List()
	return transformToStringSlice(len(list), func(idx int) string {
		return list[idx].(string)
	})
}

func transformToStringSlice(size int, mapFunc func(idx int) string) []string {
	resultSlice := make([]string, size)
	for i := 0; i < size; i++ {
		resultSlice[i] = mapFunc(i)
	}

	return resultSlice
}

func findIndexOfFirstMatch(size int, matchFunc func(idx int) bool) int {
	for i := 0; i < size; i++ {
		if matchFunc(i) {
			return i
		}
	}

	return -1
}

func slicesEqual(lhsSize int, rhsSize int, matchFunc func(lhsIdx int, rhsIdx int) bool) bool {
	if lhsSize != rhsSize {
		return false
	}

	for i := 0; i < lhsSize; i++ {
		if !matchFunc(i, i) {
			return false
		}
	}

	return true
}
