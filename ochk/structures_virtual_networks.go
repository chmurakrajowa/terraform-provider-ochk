package ochk

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenStringSlice(in []string) *schema.Set {
	out := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range in {
		out.Add(v)
	}
	return out
}
