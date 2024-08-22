package ochk

import (
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlattenDeployments(t *testing.T) {
	cases := []struct {
		expanded  []*models.DeploymentInstance
		flattened []map[strfmt.UUID]interface{}
	}{
		// nil values
		{
			expanded:  nil,
			flattened: nil,
		},
		{
			expanded: []*models.DeploymentInstance{
				{
					DeploymentID:            "3e8b5314-6c6c-4848-9830-6af905c3d878",
					DisplayName:             "rchlinux.iso",
					DeploymentType:          "ISO",
					DeploymentInitialSizeGB: 30,
				},
				{
					DeploymentID:            "198405db-6d29-41db-b2fa-5f8a0b33de39",
					DisplayName:             "centOS",
					DeploymentType:          "ISO",
					DeploymentInitialSizeGB: 40,
				},
				{
					DeploymentID:            "564e684a-b2f2-4236-ba48-ae9bf32d0f1a",
					DisplayName:             "Rocky 8",
					DeploymentType:          "ISO",
					DeploymentInitialSizeGB: 50,
				},
			},
			flattened: []map[strfmt.UUID]interface{}{
				{
					"deployment_id":   strfmt.UUID("3e8b5314-6c6c-4848-9830-6af905c3d878"),
					"display_name":    "rchlinux.iso",
					"deployment_type": models.DeploymentType("ISO"),
					"initial_size_gb": 30,
				},
				{
					"deployment_id":   strfmt.UUID("198405db-6d29-41db-b2fa-5f8a0b33de39"),
					"display_name":    "centOS",
					"deployment_type": models.DeploymentType("ISO"),
					"initial_size_gb": 40,
				},
				{
					"deployment_id":   strfmt.UUID("564e684a-b2f2-4236-ba48-ae9bf32d0f1a"),
					"display_name":    "Rocky 8",
					"deployment_type": models.DeploymentType("ISO"),
					"initial_size_gb": 50,
				},
			},
		},
	}
	for _, c := range cases {
		flattenedType := mapSliceToInterfaceSlice(c.flattened)
		outFlattened := mapSliceToInterfaceSlice(flattenDeployments(c.expanded, "ISO"))
		assert.EqualValues(t, flattenedType, outFlattened, "Error matching output and flattened: %#v vs %#v", outFlattened, flattenedType)
	}
}
