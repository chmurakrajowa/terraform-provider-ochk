package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/openapi/v3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenSnapshot(in []openapi.SnapshotInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["snapshot_id"] = v.SnapshotId
		m["display_name"] = v.SnapshotName
		m["virtual_machine_id"] = v.VirtualMachineId
		m["parent_id"] = v.ParentSnapshotId
		for i := 0; i < len(v.ChildSnapshots); i++ {
			m["child_id"] = v.ChildSnapshots[i].SnapshotId
			out = getChildSnap(v.ChildSnapshots)
		}
		out = append(out, m)
	}
	return out
}

func getChildSnap(in []openapi.SnapshotInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		n := make(map[string]interface{})
		n["snapshot_id"] = v.SnapshotId
		n["display_name"] = v.SnapshotName
		n["virtual_machine_id"] = v.VirtualMachineId
		n["parent_id"] = v.ParentSnapshotId
		if len(v.ChildSnapshots) != 0 {
			for i := 0; i < len(v.ChildSnapshots); i++ {
				n["child_id"] = v.ChildSnapshots[i].SnapshotId
				out = getChildSnap(v.ChildSnapshots)
			}
		}
		out = append(out, n)
	}
	return out
}

func flattenChildsListsFromIDs(m []openapi.SnapshotInstance) *schema.Set {
	if len(m) == 0 {
		return nil
	}

	s := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range m {
		s.Add(fmt.Sprint(v.SnapshotId))
	}

	return s
}

func expandChildSnapshots(in []interface{}) []openapi.SnapshotInstance {
	if len(in) == 0 {
		return nil
	}

	var out = make([]openapi.SnapshotInstance, len(in))
	for i, v := range in {
		var snapID string
		_, err := fmt.Sscan(v.(string), &snapID)
		if err != nil {
			return nil
		}

		snapInstance := openapi.SnapshotInstance{
			SnapshotId: NewNullableString(snapID),
		}

		out[i] = snapInstance
	}
	return out
}
