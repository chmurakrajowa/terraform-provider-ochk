package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/billing_tags"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/go-openapi/strfmt"
	"net/http"
)

type BillingTagsProxy struct {
	httpClient *http.Client
	service    billing_tags.ClientService
}

func (p *BillingTagsProxy) Read(ctx context.Context, billingTagID int32) (*models.BillingTag, error) {
	params := &billing_tags.BillingTagGetUsingGETParams{
		BillingTagID: billingTagID,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}
	response, err := p.service.BillingTagGetUsingGET(params)
	if err != nil {
		var notFound *billing_tags.BillingTagGetUsingGETNotFound

		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading billing tags: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving billing tags failed: %s", response.Payload.Messages)
	}

	return response.Payload.BillingTag, nil
}

func (p *BillingTagsProxy) ListBillingTagsByTagName(ctx context.Context, billingTagName string) ([]*models.BillingTag, error) {
	params := &billing_tags.BillingTagListUsingGETParams{
		TagValue:   &billingTagName,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.BillingTagListUsingGET(params)

	if err != nil {
		return nil, fmt.Errorf("error while listing billing tags: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("Listing billing tags failed: %s", response.Payload.Messages)
	}

	return response.Payload.BillingTagCollection, nil
}

func (p *BillingTagsProxy) Create(ctx context.Context, BillingTag *models.BillingTag) (*models.BillingTag, error) {
	if err := BillingTag.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating billing tag struct: %w", err)
	}

	params := &billing_tags.BillingTagCreateUsingPUTParams{
		BillingTag: BillingTag,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	_, put, err := p.service.BillingTagCreateUsingPUT(params)

	if err != nil {
		return nil, fmt.Errorf("error while creating billing tag: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("creating billing tag failed: %s", put.Payload.Messages)
	}

	return put.Payload.BillingTag, nil
}

func (p *BillingTagsProxy) Update(ctx context.Context, BillingTag *models.BillingTag) (*models.BillingTag, error) {
	if err := BillingTag.Validate(strfmt.Default); err != nil {
		return nil, fmt.Errorf("error while validating billing tag struct: %w", err)
	}

	params := &billing_tags.BillingTagUpdateUsingPUTParams{
		BillingTagID: BillingTag.BillingTagID,
		BillingTag:   BillingTag,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	put, err := p.service.BillingTagUpdateUsingPUT(params)

	if err != nil {
		return nil, fmt.Errorf("error while modifying billing tag: %w", err)
	}

	if !put.Payload.Success {
		return nil, fmt.Errorf("modifying billing tag failed: %s", put.Payload.Messages)
	}

	return put.Payload.BillingTag, nil
}

func (p *BillingTagsProxy) Exists(ctx context.Context, BillingTagID int32) (bool, error) {
	if _, err := p.Read(ctx, BillingTagID); err != nil {
		if IsNotFoundError(err) {
			return false, nil
		}

		return false, fmt.Errorf("error while reading billing tag: %w", err)
	}

	return true, nil
}

func (p *BillingTagsProxy) Delete(ctx context.Context, BillingTagID string) error {

	var billingTagIDInt32 int32
	fmt.Sscan(BillingTagID, &billingTagIDInt32)

	params := &billing_tags.BillingTagDeleteUsingDELETEParams{
		BillingTagID: billingTagIDInt32,
		Context:      ctx,
		HTTPClient:   p.httpClient,
	}

	response, err := p.service.BillingTagDeleteUsingDELETE(params)
	if err != nil {
		var badRequest *billing_tags.BillingTagDeleteUsingDELETEBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return &NotFoundError{Err: err}
		}

		return fmt.Errorf("error while deleting billing tag: %w", err)
	}

	if !response.Payload.Success {
		return fmt.Errorf("deleting billing tag failed: %s", response.Payload.Messages)
	}

	return nil
}
