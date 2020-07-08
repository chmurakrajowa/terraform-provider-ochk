package ochk

import (
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client"
	controller "github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/security_group_controller"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	VMRetryTimeout       = 15 * time.Minute
	VMDeleteRetryTimeout = 30 * time.Minute
)

func resourceSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceGroupCreate,
		Read:   resourceOchkVmRead,
		Update: resourceOchkVmUpdate,
		Delete: resourceOchkVmDelete,
		Exists: resourceServiceGroupExists,

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

func resourceServiceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	service := meta.(*client.Ochk).SecurityGroupController

	params := controller.NewCreateUsingPUTParams().WithSecurityGroup()

	put, created, err := service.CreateUsingPUT(params)
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

func resourceServiceGroupExists(d *schema.ResourceData, meta interface{}) (b bool, err error) {
	return false, nil
}
