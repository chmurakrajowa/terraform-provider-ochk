package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenStringSlice(in []string) *schema.Set {
	if len(in) == 0 {
		return nil
	}
	out := &schema.Set{
		F: schema.HashString,
	}

	for _, v := range in {
		out.Add(v)
	}
	return out
}

func flattenVirtualNetworks(in []*models.VirtualNetworkInstance) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["virtual_network_id"] = v.VirtualNetworkID
		m["display_name"] = v.DisplayName
		m["ipam_enabled"] = v.IpamEnabled
		m["vpc_id"] = v.RouterRefID
		m["folder_path"] = v.FolderPath
		m["project_id"] = v.ProjectID
		out = append(out, m)
	}
	return out
}

func flattenDnsServers(in []*models.DNSServerInstance) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}
	var out []map[strfmt.UUID]interface{}
	for _, v := range in {
		m := make(map[strfmt.UUID]interface{})
		m["id"] = v.ID
		m["address"] = v.Address
		out = append(out, m)
	}
	return out
}
