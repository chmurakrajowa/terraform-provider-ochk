package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type MemberType string

const (
	IPCOLLECTION MemberType = "IPCOLLECTION"
	LOGICAL_PORT MemberType = "LOGICAL_PORT"
	IPSET        MemberType = "IPSET"
)

func mapResourceDataToSecurityGroup(d *schema.ResourceData, platformType models.PlatformType) (*models.SecurityGroup, diag.Diagnostics) {

	members, err, wrongMemberType := expandSecurityGroupMembers(d.Get("members").(*schema.Set).List(), platformType)
	if err != nil {
		return nil, diag.Errorf("mapResourceDataToSecurityGroup >>>> error while creating security group. Wrong type member: %+v", wrongMemberType)
	}
	return &models.SecurityGroup{
		DisplayName: d.Get("display_name").(string),
		ProjectID:   strfmt.UUID(d.Get("project_id").(string)),
		Members:     members,
	}, nil
}
