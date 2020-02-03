package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {

	// The actual provider
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateHost,
				Description:  "host value",
			},
			"tenant": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "tenant value",
			},
			"username": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "username value",
			},
			"password": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "password value",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"virtual_machine": resourceVirtualMachine(),
		},
	}

	return provider
}

func validateHost(val interface{}, key string) (warns []string, errs []error) {
	if val == nil || val.(string) == "" {
		errs = append(errs, fmt.Errorf("%s value is not valid: %s", key, val.(string)))
	}

	return nil, errs
}
