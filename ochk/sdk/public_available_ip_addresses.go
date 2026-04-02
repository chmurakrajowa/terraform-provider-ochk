package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"net/http"
	"sync"
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
		return nil, fmt.Errorf("getting first available public ip address failed: %s", response.Messages)
	}

	return response.PublicIpAddress, nil
}
