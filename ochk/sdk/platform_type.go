package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"net/http"
	"sync"
)

type PlatformTypeProxy struct {
	httpClient *http.Client
	service    *openapi.IdentificationAPIService
}

func (p *PlatformTypeProxy) Read(ctx context.Context) (openapi.PlatformType, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IdentificationPlatformTypeGet(ctx)
	response, _, err := action.Execute()
	mutex.Unlock()

	var unknown openapi.PlatformType = "UNKNOWN"

	if err != nil {
		//var badRequest *identification.GetIdentificationPlatformTypeBadRequest
		//if ok := errors.As(err, &badRequest); ok {
		//	return unknown, &NotFoundError{Err: err}
		//}
		return unknown, fmt.Errorf("error while reading platform type: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		return unknown, fmt.Errorf("retrieving platform type failed: %s", response.Messages)
	}

	return *response.PlatformType, nil
}
