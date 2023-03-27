package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

func flattenTags(in []*models.Tag) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["tag_id"] = fmt.Sprint(v.TagID)
		m["display_name"] = v.TagValue
		m["project_id"] = v.ProjectID
		out = append(out, m)
	}
	return out
}
