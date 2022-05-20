package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/active_directory_users"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type ADUsersProxy struct {
	httpClient *http.Client
	service    active_directory_users.ClientService
}

func (p *ADUsersProxy) GetUserByUserName(ctx context.Context, userName string) (*models.ADUserInstance, error) {
	params := &active_directory_users.GetADUserUsingGETParams{
		SamAccountName: userName,
		Context:        ctx,
		HTTPClient:     p.httpClient,
	}

	response, err := p.service.GetADUserUsingGET(params)
	if err != nil {
		return nil, fmt.Errorf("error while listing users: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing users failed: %s", response.Payload.Messages)
	}

	return response.Payload.UserInstance, nil
}

func (p *ADUsersProxy) Create(ctx context.Context, adUser *models.ADUserInstance) (*models.RequestInstance, error) {
	if err := adUser.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating user struct: %w", err)
	}

	params := &active_directory_users.CreateADUserUsingPUTParams{
		UserInstance: adUser,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	_, put, err := p.service.CreateADUserUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while creating user: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating user failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *ADUsersProxy) SetPassword(ctx context.Context, adUser *models.ADUserInstance, password string) (bool, error) {

	passwordInstance := models.SetUserInstancePassRequest{
		Password: password,
	}
	params := &active_directory_users.SetPasswordADUserUsingPOSTParams{
		SamAccountName:             adUser.SamAccountName,
		SetUserInstancePassRequest: &passwordInstance,
		Context:                    ctx,
		HTTPClient:                 p.httpClient,
	}

	put, err := p.service.SetPasswordADUserUsingPOST(params)

	if err != nil {
		return false, fmt.Errorf("error while set user password1: %w", err)
	}

	if !put.Payload.Success {
		return false, fmt.Errorf("set password user failed: %s", put.Payload.Messages)
	}

	return put.Payload.Success, nil
}

func (p *ADUsersProxy) Update(ctx context.Context, adUser *models.ADUserInstance) (*models.RequestInstance, error) {
	if err := adUser.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating user struct: %w", err)
	}

	params := &active_directory_users.UpdateADUserUsingPUTParams{
		SamAccountName: adUser.SamAccountName,
		UserInstance:   adUser,
		Context:        ctx,
		HTTPClient:     p.httpClient,
	}

	put, err := p.service.UpdateADUserUsingPUT(params)
	if err != nil {
		return nil, fmt.Errorf("error while modifying user: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying user failed: %s", put.Payload.Messages)
	}

	return put.Payload.RequestInstance, nil
}

func (p *ADUsersProxy) Delete(ctx context.Context, samAccountName string) (*models.RequestInstance, error) {
	params := &active_directory_users.DeleteADUserUsingDELETEParams{
		SamAccountName: samAccountName,
		Context:        ctx,
		HTTPClient:     p.httpClient,
	}

	response, err := p.service.DeleteADUserUsingDELETE(params)
	if err != nil {
		var badRequest *active_directory_users.DeleteADUserUsingDELETEBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while deleting user: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("deleting user failed: %s", response.Payload.Messages)
	}

	return response.Payload.RequestInstance, nil
}
