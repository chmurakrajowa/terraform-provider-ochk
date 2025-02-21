// Code generated by go-swagger; DO NOT EDIT.

package firewall_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
)

// GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDReader is a Reader for the GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleID structure.
type GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /projects/{projectId}/osc/security-groups/{securityGroupId}/firewall/{ruleId}] GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleID", response, response.Code())
	}
}

// NewGetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK creates a GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK with default headers values
func NewGetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK() *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK {
	return &GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK{}
}

/*
GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK describes a response with status code 200, with default header values.

OK
*/
type GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK struct {
	Payload *models.GetFirewallRuleResponse
}

// IsSuccess returns true when this get projects project Id osc security groups security group Id firewall rule Id o k response has a 2xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get projects project Id osc security groups security group Id firewall rule Id o k response has a 3xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get projects project Id osc security groups security group Id firewall rule Id o k response has a 4xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get projects project Id osc security groups security group Id firewall rule Id o k response has a 5xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get projects project Id osc security groups security group Id firewall rule Id o k response a status code equal to that given
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get projects project Id osc security groups security group Id firewall rule Id o k response
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK) Code() int {
	return 200
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}/osc/security-groups/{securityGroupId}/firewall/{ruleId}][%d] getProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallRuleIdOK %s", 200, payload)
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}/osc/security-groups/{securityGroupId}/firewall/{ruleId}][%d] getProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallRuleIdOK %s", 200, payload)
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK) GetPayload() *models.GetFirewallRuleResponse {
	return o.Payload
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetFirewallRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest creates a GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest with default headers values
func NewGetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest() *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest {
	return &GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest{}
}

/*
GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get projects project Id osc security groups security group Id firewall rule Id bad request response has a 2xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get projects project Id osc security groups security group Id firewall rule Id bad request response has a 3xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get projects project Id osc security groups security group Id firewall rule Id bad request response has a 4xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get projects project Id osc security groups security group Id firewall rule Id bad request response has a 5xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get projects project Id osc security groups security group Id firewall rule Id bad request response a status code equal to that given
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get projects project Id osc security groups security group Id firewall rule Id bad request response
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest) Code() int {
	return 400
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}/osc/security-groups/{securityGroupId}/firewall/{ruleId}][%d] getProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallRuleIdBadRequest %s", 400, payload)
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}/osc/security-groups/{securityGroupId}/firewall/{ruleId}][%d] getProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallRuleIdBadRequest %s", 400, payload)
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized creates a GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized with default headers values
func NewGetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized() *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized {
	return &GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized{}
}

/*
GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get projects project Id osc security groups security group Id firewall rule Id unauthorized response has a 2xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get projects project Id osc security groups security group Id firewall rule Id unauthorized response has a 3xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get projects project Id osc security groups security group Id firewall rule Id unauthorized response has a 4xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get projects project Id osc security groups security group Id firewall rule Id unauthorized response has a 5xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get projects project Id osc security groups security group Id firewall rule Id unauthorized response a status code equal to that given
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get projects project Id osc security groups security group Id firewall rule Id unauthorized response
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized) Code() int {
	return 401
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}/osc/security-groups/{securityGroupId}/firewall/{ruleId}][%d] getProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallRuleIdUnauthorized %s", 401, payload)
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}/osc/security-groups/{securityGroupId}/firewall/{ruleId}][%d] getProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallRuleIdUnauthorized %s", 401, payload)
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden creates a GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden with default headers values
func NewGetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden() *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden {
	return &GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden{}
}

/*
GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden struct {
	Payload *models.ProblemDetails
}

// IsSuccess returns true when this get projects project Id osc security groups security group Id firewall rule Id forbidden response has a 2xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get projects project Id osc security groups security group Id firewall rule Id forbidden response has a 3xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get projects project Id osc security groups security group Id firewall rule Id forbidden response has a 4xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get projects project Id osc security groups security group Id firewall rule Id forbidden response has a 5xx status code
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get projects project Id osc security groups security group Id firewall rule Id forbidden response a status code equal to that given
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get projects project Id osc security groups security group Id firewall rule Id forbidden response
func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden) Code() int {
	return 403
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}/osc/security-groups/{securityGroupId}/firewall/{ruleId}][%d] getProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallRuleIdForbidden %s", 403, payload)
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /projects/{projectId}/osc/security-groups/{securityGroupId}/firewall/{ruleId}][%d] getProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallRuleIdForbidden %s", 403, payload)
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden) GetPayload() *models.ProblemDetails {
	return o.Payload
}

func (o *GetProjectsProjectIDOscSecurityGroupsSecurityGroupIDFirewallRuleIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProblemDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
