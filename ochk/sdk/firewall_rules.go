package sdk

import (
	"context"
	"fmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3"
	"github.com/go-openapi/strfmt"

	"net/http"
	"sync"
)

type FirewallRulesProxy struct {
	httpClient *http.Client
	service    *openapi.FirewallRuleAPIService
}

func (p *FirewallRulesProxy) Read(ctx context.Context, projectId strfmt.UUID, securityGroupId strfmt.UUID, ruleId strfmt.UUID) (*openapi.FirewallRule, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.ProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallRuleIdGet(ctx, string(ruleId), string(projectId), string(securityGroupId))
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while reading firwall openstack rule: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("retrieving firewall openstack rule: failed: %s", message)
	}
	if response.FirewallRule != nil {
		return response.FirewallRule, nil
	} else {
		return nil, nil
	}
}

func (p *FirewallRulesProxy) List(ctx context.Context, projectId strfmt.UUID, securityGroupId strfmt.UUID) ([]openapi.FirewallRule, error) {

	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.ProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallGet(ctx, string(projectId), string(securityGroupId))
	response, _, err := action.Execute()
	mutex.Unlock()

	if false {
		return nil, fmt.Errorf("openstack rule response.Payload.Success: %+v", response)
	}

	if err != nil {
		return nil, fmt.Errorf("error while listing firewall openstack rule: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing firewall openstack rule failed: %s", message)
	}

	return response.FirewallRuleCollection, nil
}

func (p *FirewallRulesProxy) ListByName(ctx context.Context, projectId strfmt.UUID, securityGroupId strfmt.UUID, name string) ([]openapi.FirewallRule, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.ProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallGet(ctx, string(projectId), string(securityGroupId)).Name(name)
	response, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while listing firewall rules: %w", err)
	}

	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("listing firewall rules failed: %s", message)
	}

	return response.FirewallRuleCollection, nil
}

func (p *FirewallRulesProxy) Create(ctx context.Context, projectId strfmt.UUID, securityGroupId strfmt.UUID, rule openapi.FirewallRule) (*openapi.FirewallRule, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.ProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallPut(ctx, string(projectId), string(securityGroupId)).FirewallRule(rule)
	put, _, err := action.Execute()

	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while creating firewall rule: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		message := "no error message"
		if value := put.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("creating firewall rule failed: %s", message)
	}

	return put.FirewallRule, nil
}

func (p *FirewallRulesProxy) Update(ctx context.Context, projectId strfmt.UUID, securityGroupId strfmt.UUID, rule openapi.FirewallRule) (*openapi.FirewallRule, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	action := p.service.ProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallRuleIdPut(ctx, rule.GetRuleId(), string(projectId), string(securityGroupId)).FirewallRule(rule)
	put, _, err := action.Execute()
	mutex.Unlock()

	if err != nil {
		return nil, fmt.Errorf("error while updating firewall rule: %w", err)
	}
	isSuccess := *put.Success

	if !isSuccess {
		message := "no error message"
		if value := put.Messages.Get(); value != nil {
			message = *value
		}
		return nil, fmt.Errorf("updating rule failed: %s", message)
	}

	return put.FirewallRule, nil
}

func (p *FirewallRulesProxy) Delete(ctx context.Context, projectId strfmt.UUID, securityGroupId strfmt.UUID, ruleID strfmt.UUID) error {

	action := p.service.ProjectsProjectIdOscSecurityGroupsSecurityGroupIdFirewallRuleIdDelete(ctx, string(ruleID), string(projectId), string(securityGroupId))
	response, _, err := action.Execute()

	if err != nil {
		return fmt.Errorf("error while deleting firewall rule: %w", err)
	}
	isSuccess := *response.Success

	if !isSuccess {
		message := "no error message"
		if value := response.Messages.Get(); value != nil {
			message = *value
		}
		return fmt.Errorf("deleting firewall rule failed: %s", message)
	}

	return nil
}
