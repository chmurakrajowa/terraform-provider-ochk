package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/client/requests"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"net/http"
	"time"
)

type RequestsProxy struct {
	httpClient *http.Client
	service    requests.ClientService
}

func (p *RequestsProxy) FetchResourceID(ctx context.Context, timeout time.Duration, request *models.RequestInstance) (error, string) {
	if err := verifyRequestStatusAndPhase(request); err != nil {
		return fmt.Errorf("request is not in valid state: %w", err), ""
	}

	resourceID := ""

	return resource.RetryContext(ctx, timeout, func() *resource.RetryError {
		requestState, err := p.Read(ctx, request.RequestID)
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("error reading request state: %w", err))
		}

		if err := verifyRequestStatusAndPhase(requestState); err != nil {
			return resource.NonRetryableError(fmt.Errorf("request is not in valid state: %w. %s", err, requestState.LastErrorMessage))
		}

		if requestState.RequestPhase != "FINISHED" {
			return resource.RetryableError(fmt.Errorf("expected request state FINISHED but was in state %s", requestState.RequestPhase))
		}

		resourceID = requestState.ResourceID

		return nil
	}), resourceID
}

func verifyRequestStatusAndPhase(request *models.RequestInstance) error {
	if request.RequestStatus == "FAILED" {
		return fmt.Errorf("request status is %s", request.RequestStatus)
	}

	if request.RequestPhase == "CANCELLED" || request.RequestPhase == "TIMEOUT" {
		return fmt.Errorf("request phase is %s", request.RequestPhase)
	}

	return nil
}

func (p *RequestsProxy) Read(ctx context.Context, requestID string) (*models.RequestInstance, error) {
	params := &requests.RequestGetUsingGETParams{
		RequestID:  requestID,
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	response, err := p.service.RequestGetUsingGET(params)

	if err != nil {
		var notFound *requests.RequestGetUsingGETNotFound
		if ok := errors.As(err, &notFound); ok {
			return nil, &NotFoundError{Err: err}
		}

		return nil, fmt.Errorf("error while reading request: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("retrieving request failed: %s", response.Payload.Messages)
	}

	return response.Payload.RequestInstance, nil
}
