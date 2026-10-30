package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
)

func flattenDeployments(in []openapi.DeploymentInstance, deploymentType openapi.DeploymentType) []map[strfmt.UUID]interface{} {
	if len(in) == 0 {
		return nil
	}

	var out []map[strfmt.UUID]interface{}

	for _, v := range in {
		if deploymentType == "" || deploymentType == *v.DeploymentType {
			m := make(map[strfmt.UUID]interface{})
			m["deployment_id"] = v.GetDeploymentId()
			m["display_name"] = v.GetDisplayName()
			m["deployment_type"] = v.GetDeploymentType()
			m["deployment_category"] = v.GetDeploymentCategory()
			m["initial_size_gb"] = v.GetDeploymentInitialSizeGB()
			out = append(out, m)
		}
	}
	return out
}
