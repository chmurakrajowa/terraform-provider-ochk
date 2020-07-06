package ochk

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateHost,
				Description:  "host value",
			},
			"tenant": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "tenant value",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "username value",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "password value",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"virtual_machine": resourceVirtualMachine(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			transportConfig := &client.TransportConfig{
				Host:     d.Get("host").(string),
				BasePath: "",
				Schemes:  nil,
			}

			return client.NewHTTPClientWithConfig(nil, transportConfig), nil
		},
	}
}

func validateHost(val interface{}, key string) (warns []string, errs []error) {
	if val == nil || val.(string) == "" {
		errs = append(errs, fmt.Errorf("%s value is not valid: %s", key, val.(string)))
	}

	return nil, errs
}
