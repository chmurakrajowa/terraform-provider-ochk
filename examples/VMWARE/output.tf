/*
output "vrfs" {
  value = data.ochk_vrfs.vrfs
}

output "folders" {
  value = data.ochk_folders.folders
}

output "vpcs" {
  value = data.ochk_vpcs.vpcs
}

output "dep_types" {
  value = data.ochk_deployment_params_types.dep_types
}

output "projects" {
  value = data.ochk_projects.projects
}

output "fw_rule" {
  value = ochk_firewall_ew_rule.def
}

output "backup_plans" {
  value = data.ochk_backup_plans.backup_plans
}

output "backup_lists" {
  value = data.ochk_backup_lists.backup_lists
}

output "tags" {
  value = data.ochk_tags.tags
}


output "custom_services" {
  value = data.ochk_custom_services.custom_services
}

output "services" {
  value = data.ochk_services.services
}

output "ochk_firewall_ew_rules" {
  value = data.ochk_firewall_ew_rules.firewall_ew_rules
}

output "ochk_firewall_sn_rules" {
  value = data.ochk_firewall_sn_rules.firewall_sn_rules
}

output "ochk_ip_collections" {
  value = data.ochk_ip_collections.ip_collections
}

output "ochk_kms_keys" {
  value = data.ochk_kms_keys.kms_keys
}

output "ochk_auto_nats" {
  value = data.ochk_auto_nats.auto_nats
}

output "ochk_manual_nats" {
  value = data.ochk_manual_nats.manual_nats
}


output "ochk_security_groups" {
  value = data.ochk_security_groups.security_groups
}

output "ochk_virtual_machines" {
  value = data.ochk_virtual_machines.virtual_machines
}


output "ochk_virtual_machines" {
  value = data.ochk_virtual_machines.virtual_machines
}

output "ochk_virtual_networks" {
  value = data.ochk_virtual_networks.virtual_networks
}

output "ip_collection" {
  value = ochk_ip_collection.default
}

output "project_1" {
  value = ochk_project.project-1
}

output "project_2" {
  value = ochk_project.project-2
}

output "project_3" {
  value = ochk_project.project-3
}

output "project_4" {
  value = ochk_project.project-4
}

output "user" {
  value = ochk_user.testy-radek
}

output "test_vm" {
  value = data.ochk_virtual_machine.test
}

output "project-vpc" {
  value = ochk_router.project-vpc
}

output  "auto-nat" {
  value = data.ochk_auto_nat.auto-nat
}

output  "manual-nat" {
  value = ochk_manual_nat.dnat
}

output  "manual-nat" {
  value = data.ochk_manual_nat.manual-nat
}

output  "manual-nat1" {
  value = data.ochk_manual_nat.manual-nat1
}

output "network" {
  value = data.ochk_network.project-network
}

output "user1" {
  value = data.ochk_user.bg_manager_user
}

output "vnet" {
  value = ochk_virtual_network.network_for_vm
}

output "tag" {
  value = ochk_tag.res-cc2
}


output "vm" {
  value = ochk_virtual_machine.default
}

output "aes-generated" {
  value = ochk_kms_key.aes-generated
  sensitive = true
}

output "rsa-to-unwrap" {
  value = ochk_kms_key.rsa-to-unwrap
  sensitive = true
}

output "aes-enc" {
  value = ochk_kms_key.aes-enc
  sensitive = true
}

output "encrypted-with-own-encrypted-key" {
  value = ochk_virtual_machine.encrypted-with-own-encrypted-key
}

output "encrypted-own-key" {
  value = ochk_virtual_machine.encrypted
}
*/


