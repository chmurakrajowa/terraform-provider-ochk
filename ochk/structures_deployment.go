package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
)

func flattenDeployments(in []*models.DeploymentInstance, deploymentType models.DeploymentType) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		if deploymentType == "" || deploymentType == v.DeploymentType {
			m := make(map[strfmt.UUID]interface{})
			m["deployment_id"] = v.DeploymentID
			m["display_name"] = v.DisplayName
			m["deployment_type"] = v.DeploymentType
			m["initial_size_gb"] = int(v.DeploymentInitialSizeGB)
			out = append(out, m)
		}
	}
	return out
}
