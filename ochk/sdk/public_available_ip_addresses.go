package sdk

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
)

type AvailablePublicIpProxy struct {
	httpClient *http.Client
	service    *openapi.AvailablePublicIpAPIService
}

func (p *AvailablePublicIpProxy) Get(ctx context.Context) (*openapi.PublicIpAddress, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.IpamIpaddressPublicAvailableGet(ctx)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while get first available public ip address: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("getting first available public ip address failed: %s", message)
	}

	return response.PublicIpAddress, nil
}
