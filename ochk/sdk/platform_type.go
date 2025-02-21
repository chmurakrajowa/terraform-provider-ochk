package sdk

import (
	"context"
	"errors"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/identification"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"net/http"
	"sync"
)

type PlatformTypeProxy struct {
	httpClient *http.Client
	service    identification.ClientService
}

func (p *PlatformTypeProxy) Read(ctx context.Context) (models.PlatformType, error) {
	params := &identification.GetIdentificationPlatformTypeParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetIdentificationPlatformType(params)
	mutex.Unlock()

	var unknown models.PlatformType = "UNKNOWN"

	if err != nil {
		var badRequest *identification.GetIdentificationPlatformTypeBadRequest
		if ok := errors.As(err, &badRequest); ok {
			return unknown, &NotFoundError{Err: err}
		}

		return unknown, fmt.Errorf("error while reading platform type: %w", err)
	}

	if !response.Payload.Success {
		return unknown, fmt.Errorf("retrieving platform type failed: %s", response.Payload.Messages)
	}

	return response.Payload.PlatformType, nil
}
