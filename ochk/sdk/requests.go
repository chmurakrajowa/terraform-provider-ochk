package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"net/http"
	"time"
)

type RequestsProxy struct {
	httpClient *http.Client
	service    *openapi.RequestsAPIService
}

func (p *RequestsProxy) FetchResourceID(ctx context.Context, timeout time.Duration, request *openapi.RequestInstance) (error, strfmt.UUID) {
	if err := verifyRequestStatusAndPhase(request); err != nil {
		return fmt.Errorf("request is not in valid state: %w", err), ""
	}

	var resourceID strfmt.UUID

	return resource.RetryContext(ctx, timeout, func() *resource.RetryError {
		requestState, err := p.ReadByStringValue(ctx, request.GetRequestId())
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("error reading request state: %w", err))
		}

		if err := verifyRequestStatusAndPhase(requestState); err != nil {
			return resource.NonRetryableError(fmt.Errorf("Request is not in valid state. %w. %s", err, requestState.LastErrorMessage))
		}

		if requestState.GetRequestPhase() != "FINISHED" {
			return resource.RetryableError(fmt.Errorf("expected request state FINISHED but was in state %s", requestState.RequestPhase))
		}

		if requestState.ResourceId.IsSet() {
			resourceID = strfmt.UUID(requestState.GetResourceId())
		} else {
			resourceID = ""
		}
		return nil
	}), resourceID
}

func verifyRequestStatusAndPhase(request *openapi.RequestInstance) error {
	if request.GetRequestPhase() == "FAILED" {
		return fmt.Errorf("Request status is %s", request.RequestStatus)
	}

	if request.GetRequestPhase() == "CANCELLED" || request.GetRequestPhase() == "TIMEOUT" {
		return fmt.Errorf("Request phase is %s", request.RequestPhase)
	}

	return nil
}

func (p *RequestsProxy) Read(ctx context.Context, requestID strfmt.UUID) (*openapi.RequestInstance, error) {
	action := p.service.RequestRequestIdGet(ctx, string(requestID))
	response, _, err := action.Execute()
	if err != nil {
		return nil, fmt.Errorf("error while reading request: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving request failed: %s", response.Messages)
	}

	return response.RequestInstance, nil
}

func (p *RequestsProxy) ReadByStringValue(ctx context.Context, requestID string) (*openapi.RequestInstance, error) {
	action := p.service.RequestRequestIdGet(ctx, requestID)
	response, _, err := action.Execute()
	if err != nil {
		return nil, fmt.Errorf("error while reading request: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return nil, fmt.Errorf("retrieving request failed: %s", response.Messages)
	}

	return response.RequestInstance, nil
}
