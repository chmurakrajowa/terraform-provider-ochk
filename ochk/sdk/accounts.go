package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type AccountsProxy struct {
	httpClient *http.Client
	service    *openapi.AccountsAPIService
}

func (p *AccountsProxy) Read(ctx context.Context, accountID strfmt.UUID) (*openapi.AccountInstance, error) {

	action := p.service.BillingAccountsAccountIdGet(ctx, string(accountID))

	response, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while reading billing account: %w", err)
	}
	//
	isSuccess := *response.Success
	if !isSuccess {
		return nil, fmt.Errorf("retrieving billing account failed: %s", response.Messages)
	}

	return response.AccountInstance, nil
}

func (p *AccountsProxy) ListAccountByName(ctx context.Context, accountName string) ([]openapi.AccountInstance, error) {

	action := p.service.BillingAccountsGet(ctx).Name(accountName)
	response, _, err := action.Execute()
	if err != nil {
		return nil, fmt.Errorf("error while listing billing accounts: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("listing billing accounts failed: %s", response.Messages)
	}
	return response.AccountInstanceCollection, nil
}

func (p *AccountsProxy) ListAccounts(ctx context.Context) ([]openapi.AccountInstance, error) {

	action := p.service.BillingAccountsGet(ctx)
	response, _, err := action.Execute()
	if err != nil {
		return nil, fmt.Errorf("error while listing billing accounts: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("Listing billing accounts failed: %s", response.Messages)
	}

	return response.AccountInstanceCollection, nil
}

func (p *AccountsProxy) Create(ctx context.Context, account openapi.AccountInstance) (*openapi.AccountInstance, error) {
	action := p.service.BillingAccountsPut(ctx).AccountInstance(account)
	put, _, err := action.Execute()
	if err != nil {
		return nil, fmt.Errorf("error while creating account: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("creating account failed: %s", put.Messages)
	}

	return put.AccountInstance, nil
}

func (p *AccountsProxy) Delete(ctx context.Context, accountID strfmt.UUID) error {

	action := p.service.BillingAccountsAccountIdDelete(ctx, string(accountID))

	response, _, err := action.Execute()
	if err != nil {

		return fmt.Errorf("error while deleting accountt: %w", err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		return fmt.Errorf("deleting accountt failed: %s", response.Messages)
	}

	return nil
}

func (p *AccountsProxy) Update(ctx context.Context, account *openapi.AccountInstance) (*openapi.AccountInstance, error) {
	action := p.service.BillingAccountsAccountIdPut(ctx, account.GetAccountId())

	put, _, err := action.Execute()

	if err != nil {
		return nil, fmt.Errorf("error while modifying account: %w", err)
	}

	isSuccess := *put.Success

	if !isSuccess {
		return nil, fmt.Errorf("modifying account failed: %s", put.Messages)
	}
	return put.AccountInstance, nil
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
