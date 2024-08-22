package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/accounts"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type AccountsProxy struct {
	httpClient *http.Client
	service    accounts.ClientService
}

func (p *AccountsProxy) Read(ctx context.Context, accountID strfmt.UUID) (*models.AccountInstance, error) {
	params := &accounts.GetBillingAccountsAccountIDParams{
		AccountID:  accountID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.GetBillingAccountsAccountID(params)

	if err != nil {
		var notFound *accounts.GetBillingAccountsAccountIDNotFound

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
	params := &accounts.GetBillingAccountsParams{
		Name:       &accountName,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.GetBillingAccounts(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing billing accounts: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing billing accounts failed: %s", response.Payload.Messages)
	}
	return response.Payload.AccountInstanceCollection, nil
}

func (p *AccountsProxy) ListAccounts(ctx context.Context) ([]*models.AccountInstance, error) {
	params := &accounts.GetBillingAccountsParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.GetBillingAccounts(params)

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

	params := &accounts.PutBillingAccountsParams{
		Body:       account,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	put, err := p.service.PutBillingAccounts(params)

	if err != nil {
		return nil, fmt.Errorf("error while creating account: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating account failed: %s", put.Payload.Messages)
	}

	return put.Payload.AccountInstance, nil
}

func (p *AccountsProxy) Delete(ctx context.Context, accountID strfmt.UUID) error {

	params := &accounts.DeleteBillingAccountsAccountIDParams{
		AccountID:  accountID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.DeleteBillingAccountsAccountID(params)
	if err != nil {
		var badRequest *accounts.DeleteBillingAccountsAccountIDBadRequest
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

	params := &accounts.PutBillingAccountsAccountIDParams{
		AccountID:  account.AccountID,
		Body:       account,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	put, err := p.service.PutBillingAccountsAccountID(params)

	if err != nil {
		return nil, fmt.Errorf("error while modifying account: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying account failed: %s", put.Payload.Messages)
	}
	return put.Payload.AccountInstance, nil
}

func (p *AccountsProxy) Exists(ctx context.Context, accountID strfmt.UUID) (bool, error) {
	if _, err := p.Read(ctx, accountID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}
		return false, fmt.Errorf("error while reading account: %w", err)
	}
	return true, nil
}
