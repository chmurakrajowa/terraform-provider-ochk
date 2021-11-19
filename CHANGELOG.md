## 1.2.1 (2021-11-19)

## Resources
* Changed resources
  * `resource_virtual_machine`
  * `resource_subtenant`
  * `resource_virtual_nmetwork`

## Data sources
* Changed data sources:
  * `data_source_group`
  * `data_source_subtenant`
  * `data_source_virtual_nmetwork`
  * `data_source_virtual_machine`



## Bugfix
* `Fix for recreation vNet`
  * `Change for vnet  parameters for IPAM fields (dns_search_suffix, dns_suffix, primary_dns_address, second_dns_address, primary_wins_address, second_wins_address). Changing any of those parameter will not recrate network resources.`
* `Fix for recreation subtenant`
  * `Change for subtenant (Business Group)  parameters. Parameter "users" will not  be required any more. This change eliminates recreation Business Group at access management to  resources`
* `Fix for security group`

### Features
* `Instead of value STANDARD put STANDARD_W1, STANDARD_W2 for storage_policy in virtual machine object`
  * `Configuration change for availability policy from "STANDARD" to "STANDARD W1" and "STANDARD_W2" describes where virtual machine is running (which data center)`

## 1.2 (2021-08-20)

## Resources
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

## Data sources
* New data sources:
  * `data_source_backup_list`
  * `data_source_backup_plan`
  * `data_source_billing_tag`
  * `data_source_system_tag`
    
Changed data sources:
  * `data_source_router`


## Misc
  * `Support for Terraform 1.0`

## Bugfix
  * `Fix for creating machine and setting virtual machine password - inital password is set correctly`  
  * `Fix for KMS encrypted key - every run apply command has caused to recreation KMS key`  
  * `Change for creating S-N firewall rule - putting ID router VPC instead of security_policy_id`

## Features
  * `Added ssh-key upload when crating virtual machine`
  * `Added system and billing tags management and the ability to assign them to virtual machines`
  * `Added virtual machine backup plan management`
  * `Added deploymnet virtual machine from OVF file`
  * `Added to terrafrom state IP addreess virtual mnachine`
  



## 1.1 (2021-03-23)

## Resources
* New resources:
  * `ochk_kms_key`
* Changed resources:
  * `ochk_virtual_machine` - added encryption settings

## Data sources
* New data sources:
  * `ochk_kms_key`
* Removed data sources:
  * `ochk_ip_set`

## Misc
* Added dumping request/response body to terraform log when TF_LOG=debug

## 1.0 (2020-11-26)

## Resources
* New resources:
  * `ochk_custom_service`
  * `ochk_ip_collection`
  * `ochk_subtenant`
  * `ochk_virtual_machine`
  * `ochk_virtual_network`

## Data sources
* New data sources:
  * `ochk_custom_service`
  * `ochk_deployment`
  * `ochk_ip_collection`
  * `ochk_network`
  * `ochk_subtenant`
  * `ochk_user`
  * `ochk_virtual_network`

## Misc
* Updated `examples` directory with ready-to-use terraform code.

## 0.1.1-pre (2020-08-27)

## Resources
* New resource for managing routers: `ochk_router`.

## Data sources 
* New data source for logical ports: `ochk_logical_port`.

## Misc 
* Added caching credentials for improved performance
* Updated markdown docs

## 0.1.0-pre (2020-08-20)

## Resources
* New resource for managing routers: `ochk_router`.

## Data sources 
* `gateway_policy`
* `ip_set`
* `logical_port`
* `router`
* `security_group`
* `security_policy`
* `service`
* `virtual_machine`

## Misc 
* Added caching credentials for improved performance.
* Updated markdown docs


 