package ochk

func toInterfaceSlice(in []map[string]interface{}) []interface{} {
	output := make([]interface{}, len(in))
	for i := range in {
		output[i] = in[i]
	}

	return output
}
