package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenTags(in []*models.Tag) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["tag_id"] = fmt.Sprint(v.TagID)
		m["display_name"] = v.TagValue
		m["project_id"] = v.ProjectID
		m["related_virtual_machines"] = v.RelatedVirtualMachines
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
