package ochk

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	VMRetryTimeout       = 15 * time.Minute
	VMDeleteRetryTimeout = 30 * time.Minute
)

func resourceVirtualMachine() *schema.Resource {
	return &schema.Resource{
		Create: resourceOchkVmCreate,
		Read:   resourceOchkVmRead,
		Update: resourceOchkVmUpdate,
		Delete: resourceOchkVmDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(VMRetryTimeout),
			Update: schema.DefaultTimeout(VMRetryTimeout),
			Delete: schema.DefaultTimeout(VMDeleteRetryTimeout),
		},

		Schema: map[string]*schema.Schema{
			"count": {
				Type:     schema.TypeInt,
				Required: true,
				Default:  1,
			},
			"catalog_item_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_configuration": {
				Type:     schema.TypeMap,
				Required: true,
				Elem:     schema.TypeString,
			},
		},
	}
}

func resourceOchkVmCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceOchkVmRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceOchkVmUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceOchkVmDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
