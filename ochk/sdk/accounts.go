package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/billing_accounts"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type AccountsProxy struct {
	httpClient *http.Client
	service    billing_accounts.ClientService
}

func (p *AccountsProxy) Read(ctx context.Context, accountID string) (*models.AccountInstance, error) {
	params := &billing_accounts.AccountGetUsingGETParams{
		AccountID:  accountID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.AccountGetUsingGET(params)

	if err != nil {
		var notFound *billing_accounts.AccountGetUsingGETNotFound

		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading billing account: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving billing account failed: %s", response.Payload.Messages)
	}

	return response.Payload.AccountInstance, nil
}

func (p *AccountsProxy) ListAccountByName(ctx context.Context, accountName string) ([]*models.AccountInstance, error) {
	params := &billing_accounts.AccountListUsingGETParams{
		Name:       &accountName,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.AccountListUsingGET(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing billing accounts: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing billing accounts failed: %s", response.Payload.Messages)
	}
	return response.Payload.AccountInstanceCollection, nil
}

func (p *AccountsProxy) ListAccounts(ctx context.Context) ([]*models.AccountInstance, error) {
	params := &billing_accounts.AccountListUsingGETParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.AccountListUsingGET(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing billing accounts: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("Listing billing accounts failed: %s", response.Payload.Messages)
	}

	return response.Payload.AccountInstanceCollection, nil
}

func (p *AccountsProxy) Create(ctx context.Context, account *models.AccountInstance) (*models.AccountInstance, error) {
	if err := account.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating account struct: %w", err)
	}

	params := &billing_accounts.AccountCreateUsingPUTParams{
		AccountInstance: account,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	put, _, err := p.service.AccountCreateUsingPUT(params)

	if err != nil {
		return nil, fmt.Errorf("error while creating account: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating account failed: %s", put.Payload.Messages)
	}

	return put.Payload.AccountInstance, nil
}

func (p *AccountsProxy) Delete(ctx context.Context, accountID string) error {

	params := &billing_accounts.AccountDeleteUsingDELETEParams{
		AccountID:  accountID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, _, err := p.service.AccountDeleteUsingDELETE(params)
	if err != nil {
		var badRequest *billing_accounts.AccountGetUsingGETBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting accountt: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting accountt failed: %s", response.Payload.Messages)
	}

	return nil
}

func (p *AccountsProxy) Update(ctx context.Context, account *models.AccountInstance) (*models.AccountInstance, error) {
	if err := account.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating account struct: %w", err)
	}

	params := &billing_accounts.AccountUpdateUsingPUTParams{
		AccountID:       account.AccountID,
		AccountInstance: account,
		Context:         ctx,
		HTTPClient:      p.httpClient,
	}

	put, _, err := p.service.AccountUpdateUsingPUT(params)

	if err != nil {
		return nil, fmt.Errorf("error while modifying account: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying account failed: %s", put.Payload.Messages)
	}
	return put.Payload.AccountInstance, nil
}

func (p *AccountsProxy) Exists(ctx context.Context, accountID string) (bool, error) {
	if _, err := p.Read(ctx, accountID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}
		return false, fmt.Errorf("error while reading account: %w", err)
	}
	return true, nil
}
