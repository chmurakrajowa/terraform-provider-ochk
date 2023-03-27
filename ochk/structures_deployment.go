package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

func flattenDeployments(in []*models.DeploymentInstance, deploymentType string) []map[string]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[string]interface{}

	for _, v := range in {
		if deploymentType == "" || deploymentType == v.DeploymentType {
			m := make(map[string]interface{})
			m["deployment_id"] = v.DeploymentID
			m["display_name"] = v.DisplayName
			m["deployment_type"] = v.DeploymentType
			m["initial_size_mb"] = int(v.DeploymentInitialSizeMB)
			out = append(out, m)
		}
	}
	return out
}
