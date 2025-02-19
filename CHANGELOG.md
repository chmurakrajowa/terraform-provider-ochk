## 3.0.0 (2025-02-17)

### Resources

* Added new resources for OpenStack
  * `resource_floating_ip` - reservation of public IP addresses
  * `resource_port_forwarding` - network address translation
  * `resource_firewall_rule`

### Data Sources

* Added new data sources for OpenStack
  * `data_source_floating_ip`
  * `data_source_floating_ips`
  * `data_source_port_forwarding`
  * `data_source_ports_forwarding`
  * `data_source_firewall_rule`
  * `data_source_firewall_rules`
  * `data_source_floating_ip_vms`

### Misc
* Added support for two types of platforms: VMWARE and OPENSTACK. Changed login parameters: to set the type of platform that we currently use,`platform_type` param was added.
* Added `autonat_enabled` to `resource_vpc` for OPENSTACK
* Added platform_type checking validation to `resource_security_group` to verify members_type. Only `VIRTUAL_MACHINE` member_type is allowed for OPENSTACK.
* Changed the resource id type from string to uuid
* Added deployment_category to data_source_deployment
* Changed MB to GB project size in `resource_project`


## 2.1.0 (2024-07-01)

### Resources
* Added data sources
  * `data_source_billing_account`
  * `data_source_billing_accounts`
  * `data_source_snapshot`
  * `data_source_snapshots`

* Added resources
  * `resource_billing_account`
  * `resource_snapshot`

### Security
* Updated external libraries versions in go.mod for security reasons
* Changed basic auth method to x_api_key method

### Misc
* Changed import functions - id is case-insensitive now

## 2.0.3 (2023-03-27)

### Security
* Updated external libraries versions in go.mod for security reasons

## 2.0.2 (2023-03-27)

### Resources
* Changed resources
  * `resource_custom_service`
  * `resource_virtual_machine`
  * `resource_firewall_sn_rule`
  * `resource_firewall_ew_rule`
  * `resource_virtual_network`
  * `resource_ip_collection`
  * `resource_security_group`

* Added resources
  * `resource_tag`
  * `resource_project`
  * `resource_vpc`

* Removed resources
  * `resource_subtenant`
  * `resource_router`
  * `resource_system_tag`
  * `resource_billing_tag`

### Data sources
* Changed data sources
  * `data_source_custom_service`
  * `data_source_service`
  * `data_source_security_group`
  * `data_source_deployment`
  * `data_source_firewall_ew_rule`
  * `data_source_firewall_sn_rule`
  * `data_source_virtual_machine`
  * `data_source_virtual_network`
  * `data_source_ip_collection`

* Added data sources
  * `data_source_public_ip_addresses`
  * `data_source_backup_lists`
  * `data_source_nats`
  * `data_source_project`
  * `data_source_backup_plans`
  * `data_source_projects`
  * `data_source_custom_services`
  * `data_source_security_groups`
  * `data_source_deployment_params_types`
  * `data_source_deployment_types`
  * `data_source_services`
  * `data_source_deployments`
  * `data_source_tag`
  * `data_source_tags`
  * `data_source_firewall_ew_rules`
  * `data_source_virtual_machines`
  * `data_source_firewall_sn_rules`
  * `data_source_folders`
  * `data_source_virtual_networks`
  * `data_source_vpc`
  * `data_source_ip_collections`
  * `data_source_vpcs`
  * `data_source_vrf`
  * `data_source_kms_keys`
  * `data_source_vrfs`

* Removed data sources
  * `data_source_subtenant`
  * `data_source_router`
  * `data_source_gateway_policy`
  * `data_source_security_policy`
  * `data_source_system_tag`
  * `data_source_billing_tag`

### Misc
* Changed host url form `waw1.api.ochk.pl` to `api.portal.ochk.pl`
* Changed login parameters, `tenant` param was replaced with `platform` param
* Subtenant resource replaced with `project` resource
* Billing_tag and system_tag resources replaced with `tag` resource
* Router resource replaced with `vpc` and `vrf` resources
  * `router_id` param replaced with `vpc_id`
  * `parent_router_id` param replaced with `vrf_id`
* Added `folder_path` in resources: `virtual_machines`, `vpc`, `virtual_network`
* Added `project_id` in same resources

## 1.2.3 (2022-05-20)

### Resources
* Changed resources
  * `resource_firewall_ew_rule`
  * `resource_firewall_sn_rule`
  * `router`
  * `securioty_group`
  * `virtual_machine`

* Add resources
  * `manual_nat`
  * `auto_nat`

### Data sources
* Changed data sources
  * `virtual_machine`

* Remove data sources
  * `gateway_policy`
  * `security_policy`
  * `networks` Legacy Networks - instead by `virtual_networks`

* Add data sources
  * `firewall_ew_rule`
  * `firewall_sn_rule`
  * `auto_nat`
  * `manual_nat`

### Misc
* Changed return params by Virtual Machine. Returned id virtual machine is the same
  as id in virtual machine resource
* Data source network was definitely replaced by virtual_networks
* Fix for importer rules mechanism (firewall_ew_rule, firewall_sn_rule)
* Changed resource router, if parent_router_id is changed, then router is recreated
* Changed resource security_group. An id virtual machine has to the same as virtual machine
  id in data source virtual machine
* Resource manual_nat let to translate internal network addresses to a single public address
* Resource auto_nat let to make configuration NAT rule with output to internet

## 1.2.2 (2021-11-25)

### Resources
* Changed resources
  * `resource_firewall_ew_rule`
  * `resource_firewall_sn_rule`

### Structure
* Changed structures
  * `structure_users`
  * `structure_router_instance`

### Bugfix
* `Fix for modify S-N rule, adding 'scope' field to object S-N rule`
* `Fix for modify custom services in S-N and E-W rules`

## 1.2.1 (2021-11-19)

### Resources
* Changed resources
  * `resource_virtual_machine`
  * `resource_subtenant`
  * `resource_virtual_nmetwork`

### Data sources
* Changed data sources
  * `data_source_group`
  * `data_source_subtenant`
  * `data_source_virtual_nmetwork`
  * `data_source_virtual_machine`

### Bugfix
* `Fix for recreation vNet`
  * `Change for vnet  parameters for IPAM fields (dns_search_suffix, dns_suffix, primary_dns_address, second_dns_address, primary_wins_address, second_wins_address). Changing any of those parameter will not recrate network resources.`
* `Fix for recreation subtenant`
  * `Change for subtenant (Business Group)  parameters. Parameter "users" will not  be required any more. This change eliminates recreation Business Group at access management to  resources`
* `Fix for security group`

### Features
* `Instead of value STANDARD put STANDARD_W1, STANDARD_W2 for storage_policy in virtual machine object`
  * `Configuration change for availability policy from "STANDARD" to "STANDARD W1" and "STANDARD_W2" describes where virtual machine is running (which data center)`

## 1.2 (2021-08-20)

### Resources
* New resources
  * `resource_billing_tag`
  * `resource_system_tag`

* Changed resources
  * `resource_custom_service`
  * `resource_firewall_ew_rule`
  * `resource_firewall_sn_rule`
  * `resource_ip_collection`
  * `resource_kms_key`
  * `resource_router`
  * `resource_security_group`
  * `resource_subtenant`
  * `resource_virtual_machine`
  * `resource_virtual_network`

### Data sources
* New data sources
  * `data_source_backup_list`
  * `data_source_backup_plan`
  * `data_source_billing_tag`
  * `data_source_system_tag`

* Changed data sources
  * `data_source_router`

### Misc
* `Support for Terraform 1.0`

### Bugfix
* `Fix for creating machine and setting virtual machine password - inital password is set correctly`
* `Fix for KMS encrypted key - every run apply command has caused to recreation KMS key`
* `Change for creating S-N firewall rule - putting ID router VPC instead of security_policy_id`

### Features
* `Added ssh-key upload when crating virtual machine`
* `Added system and billing tags management and the ability to assign them to virtual machines`
* `Added virtual machine backup plan management`
* `Added deploymnet virtual machine from OVF file`
* `Added to terrafrom state IP addreess virtual mnachine`

## 1.1 (2021-03-23)

### Resources
* New resources
  * `ochk_kms_key`
* Changed resources
  * `ochk_virtual_machine` - added encryption settings

### Data sources
* New data sources
  * `ochk_kms_key`
* Removed data sources
  * `ochk_ip_set`

### Misc
* Added dumping request/response body to terraform log when TF_LOG=debug

## 1.0 (2020-11-26)

### Resources
* New resources
  * `ochk_custom_service`
  * `ochk_ip_collection`
  * `ochk_subtenant`
  * `ochk_virtual_machine`
  * `ochk_virtual_network`

### Data sources
* New data sources
  * `ochk_custom_service`
  * `ochk_deployment`
  * `ochk_ip_collection`
  * `ochk_network`
  * `ochk_subtenant`
  * `ochk_user`
  * `ochk_virtual_network`

### Misc
* Updated `examples` directory with ready-to-use terraform code.

## 0.1.1-pre (2020-08-27)

### Resources
* New resource for managing routers `ochk_router`.

### Data sources
* New data source for logical ports `ochk_logical_port`.

### Misc
* Added caching credentials for improved performance
* Updated markdown docs

## 0.1.0-pre (2020-08-20)

### Resources
* New resource for managing routers `ochk_router`.

### Data sources
* `gateway_policy`
* `ip_set`
* `logical_port`
* `router`
* `security_group`
* `security_policy`
* `service`
* `virtual_machine`

### Misc
* Added caching credentials for improved performance.
* Updated markdown docs
