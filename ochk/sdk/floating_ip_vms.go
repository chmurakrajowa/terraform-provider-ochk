package sdk

import (
	"context"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/client/floating_ip_vms"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"net/http"
	"sync"
)

type FloatingIPVmsProxy struct {
	httpClient *http.Client
	service    floating_ip_vms.ClientService
}

func (p *FloatingIPVmsProxy) List(ctx context.Context) ([]*models.PortFwdVM, error) {

	params := &floating_ip_vms.GetNetworkFloatingIpsVmsParams{
		Context:    ctx,
		HTTPClient: p.httpClient,
	}

	mutex := sync.Mutex{}
	mutex.Lock()
	response, err := p.service.GetNetworkFloatingIpsVms(params)
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing floating ip vms: %w", err)
	}

	if !response.Payload.Success {
		return nil, fmt.Errorf("listing floating ip vms failed: %s", response.Payload.Messages)
	}

	return response.Payload.PortFwdVMCollection, nil
}
