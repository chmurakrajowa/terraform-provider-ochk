package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenTags(in []openapi.Tag) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["tag_id"] = fmt.Sprint(v.GetTagId())
		m["display_name"] = v.GetTagValue()
		m["project_id"] = v.GetProjectId()
		m["related_virtual_machines"] = v.GetRelatedVirtualMachines()
		out = append(out, m)
	}
	return out
}

func flattenTagsLists(m []strfmt.UUID) *schema.Set {

	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(fmt.Sprint(v.Value()))
	}

	return s
}
