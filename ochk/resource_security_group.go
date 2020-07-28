package ochk

import (
	"context"
	"errors"
	"fmt"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk"
	controller "github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/client/security_groups"
	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	SecurityGroupRetryTimeout = 1 * time.Minute
)

func resourceSecurityGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceGroupCreate,
		Read:   resourceServiceGroupRead,
		Update: resourceServiceGroupUpdate,
		Delete: resourceServiceGroupDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(SecurityGroupRetryTimeout),
			Update: schema.DefaultTimeout(SecurityGroupRetryTimeout),
			Delete: schema.DefaultTimeout(SecurityGroupRetryTimeout),
		},

		Schema: map[string]*schema.Schema{
			"display_name": {
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
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceServiceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	service := client.GetOchk().SecurityGroups

	securityGroup := &models.SecurityGroup{
		DisplayName: d.Get("display_name").(string),
		Members:     expandSecurityGroupMembers(d.Get("members").([]interface{})),
	}

	if err := securityGroup.Validate(nil); err != nil {
		return fmt.Errorf("error while validating security group structure: %+v", err)
	}

	params := &controller.SecurityGroupCreateUsingPUTParams{
		SecurityGroup: securityGroup,
		Context:       context.Background(),
		HTTPClient:    client.GetHTTPClient(),
	}

	_, put, err := service.SecurityGroupCreateUsingPUT(params)
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
	client := meta.(*sdk.Client)
	service := client.GetOchk().SecurityGroups

	params := &controller.SecurityGroupGetUsingGETParams{
		GroupID:    d.Id(),
		Context:    context.Background(),
		HTTPClient: client.GetHTTPClient(),
	}

	response, err := service.SecurityGroupGetUsingGET(params)
	if err != nil {
		return fmt.Errorf("error while reading security group: %+v", err)
	}

	var notFound *controller.SecurityGroupGetUsingGETNotFound
	if ok := errors.As(err, &notFound); ok {
		d.SetId("")
		return nil
	}

	if !response.Payload.Success {
		return fmt.Errorf("retrieving security group failed: %s", response.Payload.Messages)
	}

	securityGroup := response.Payload.SecurityGroup

	if err := d.Set("display_name", securityGroup.DisplayName); err != nil {
		return fmt.Errorf("error setting displayName: %+v", err)
	}

	if err := d.Set("members", flattenSecurityGroupMembers(securityGroup.Members)); err != nil {
		return fmt.Errorf("error setting members: %+v", err)
	}

	return nil
}

func resourceServiceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("updating SecurityGroup is not implemented")
}

func resourceServiceGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*sdk.Client)
	service := client.GetOchk().SecurityGroups

	params := &controller.SecurityGroupDeleteUsingDELETEParams{
		GroupID:    d.Id(),
		Context:    context.Background(),
		HTTPClient: client.GetHTTPClient(),
	}

	response, err := service.SecurityGroupDeleteUsingDELETE(params)
	if err != nil {
		return fmt.Errorf("error while deleting security group: %+v", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("retrieving security group failed: %s", response.Payload.Messages)
	}

	//TODO po co nam security group w odpowiedzi? securityGroup := response.Payload.SecurityGroup

	return nil
}
