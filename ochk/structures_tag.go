package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
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
		out = append(out, m)
	}
	return out
}
