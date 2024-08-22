package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenSnapshot(in []*models.SnapshotInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["snapshot_id"] = v.SnapshotID
		m["display_name"] = v.SnapshotName
		m["virtual_machine_id"] = v.VirtualMachineID
		m["parent_id"] = v.ParentSnapshotID
		for i := 0; i < len(v.ChildSnapshots); i++ {
			m["child_id"] = v.ChildSnapshots[i].SnapshotID
			out = getChildSnap(v.ChildSnapshots)
		}
		out = append(out, m)
	}
	return out
}

func getChildSnap(in []*models.SnapshotInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		n := make(map[string]interface{})
		n["snapshot_id"] = v.SnapshotID
		n["display_name"] = v.SnapshotName
		n["virtual_machine_id"] = v.VirtualMachineID
		n["parent_id"] = v.ParentSnapshotID
		if len(v.ChildSnapshots) != 0 {
			for i := 0; i < len(v.ChildSnapshots); i++ {
				n["child_id"] = v.ChildSnapshots[i].SnapshotID
				out = getChildSnap(v.ChildSnapshots)
			}
		}
		out = append(out, n)
	}
	return out
}

func flattenChildsListsFromIDs(m []*models.SnapshotInstance) *schema.Set {
	if len(m) == 0 {
		return nil
	}

	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(fmt.Sprint(v.SnapshotID))
	}

	return s
}

func expandChildSnapshots(in []interface{}) []*models.SnapshotInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]*models.SnapshotInstance, len(in))
	for i, v := range in {
		var snapID string
		_, err := fmt.Sscan(v.(string), &snapID)
		if err != nil {
			return nil
		}

		snapInstance := &models.SnapshotInstance{
			SnapshotID: strfmt.UUID(snapID),
		}

		out[i] = snapInstance
	}
	return out
}
