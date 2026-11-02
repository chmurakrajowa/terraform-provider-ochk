package sdk

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
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
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return unknown, fmt.Errorf("retrieving platform type failed: %s", message)
	}

	return *response.PlatformType, nil
}
