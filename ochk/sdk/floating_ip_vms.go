package sdk

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
)

type FloatingIPVmsProxy struct {
	httpClient *http.Client
	service    *openapi.FloatingIpVmsAPIService
}

func (p *FloatingIPVmsProxy) List(ctx context.Context) ([]openapi.PortFwdVm, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.NetworkFloatingIpsVmsGet(ctx)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing floating ip vms: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing floating ip vms failed: %s", message)
	}

	return response.PortFwdVmCollection, nil
}
