package ochk

import (
	"fmt"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client"
	controller "github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/security_group_controller"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
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
		Read:   resourceServiceGroupRead,
		Update: resourceServiceGroupUpdate,
		Delete: resourceServiceGroupDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(VMRetryTimeout),
			Update: schema.DefaultTimeout(VMRetryTimeout),
			Delete: schema.DefaultTimeout(VMDeleteRetryTimeout),
		},

		Schema: map[string]*schema.Schema{
			"displayName": {
				Type:     schema.TypeString,
				Required: true,
			},
			"members": {
				Type:     schema.TypeList,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ExactlyOneOf: []string{"IPSET", "VIRTUAL_MACHINE", "LOGICAL_PORT"},
						},
					},
				},
			},
		},
	}
}

func resourceServiceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	service := meta.(*client.Ochk).SecurityGroupController

	securityGroup := &models.SecurityGroup{
		DisplayName: d.Get("display_name").(string),
		Members:     nil,
	}

	params := controller.NewCreateUsingPUTParams().WithSecurityGroup(securityGroup)

	put, _, err := service.CreateUsingPUT(params)
	if err != nil {
		return fmt.Errorf("error while creating security group: %+v", err)
	}

	if !put.Payload.Success {
		return fmt.Errorf("creating security group failed: %s", put.Payload.Messages)
	}

	d.SetId(put.Payload.SecurityGroup.ID)

	return nil
}

func resourceServiceGroupRead(d *schema.ResourceData, meta interface{}) error {
	service := meta.(*client.Ochk).SecurityGroupController

	params := controller.NewGetUsingGET3Params().WithGroupID(d.Id())

	response, err := service.GetUsingGET3(params)
	if err != nil {
		return fmt.Errorf("error while reading security group: %+v", err)
	}

	securityGroup := response.Payload.SecurityGroup

	err = d.Set("members", flattenSecurityGroupMember(securityGroup.Members))
	if err != nil {

	}

	return nil
}

func resourceServiceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceServiceGroupDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
