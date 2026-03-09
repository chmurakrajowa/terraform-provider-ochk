package ochk

import (
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenSnapshot(in []openapi.SnapshotInstance) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		m := make(map[string]interface{})
		m["snapshot_id"] = v.GetSnapshotId()
		m["display_name"] = v.GetSnapshotName()
		m["virtual_machine_id"] = v.GetVirtualMachineId()
		m["parent_id"] = v.GetParentSnapshotId()
		for i := 0; i < len(v.GetChildSnapshots()); i++ {
			m["child_id"] = v.ChildSnapshots[i].GetSnapshotId()
			out = getChildSnap(v.GetChildSnapshots())
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
		n["snapshot_id"] = v.GetSnapshotId()
		n["display_name"] = v.GetSnapshotName()
		n["virtual_machine_id"] = v.GetVirtualMachineId()
		n["parent_id"] = v.GetParentSnapshotId()
		if len(v.GetChildSnapshots()) != 0 {
			for i := 0; i < len(v.GetChildSnapshots()); i++ {
				n["child_id"] = v.ChildSnapshots[i].GetSnapshotId()
				out = getChildSnap(v.GetChildSnapshots())
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
		s.Add(fmt.Sprint(v.GetSnapshotId()))
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
