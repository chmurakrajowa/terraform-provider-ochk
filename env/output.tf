output "ip_collection" {
  value = ochk_ip_collection.default
}

output "subtenant_1" {
  value = ochk_subtenant.subtenant-1
}

output "subtenant_2" {
  value = ochk_subtenant.subtenant-2
}

output "subtenant_3" {
  value = ochk_subtenant.subtenant-3
}

output "subtenant_4" {
  value = ochk_subtenant.subtenant-4
}

output "network" {
  value = ochk_virtual_network.subtenant-network
}

output "vnet" {
  value = ochk_virtual_network.default
}

output "vm" {
  value = ochk_virtual_machine.default
}

output "custom-service-1" {
  value = ochk_custom_service.web_servers_https
}

output "custom-service-2" {
  value = ochk_custom_service.web_servers_http
}

output "billing-tag" {
  value = ochk_billing_tag.res-bt1
}

output "system-tag" {
  value = ochk_system_tag.res-st1
}

output "auto-nat" {
  value = ochk_auto_nat.auto-nat
}

output "dnat" {
  value = ochk_manual_nat.dnat
}

output "snat" {
  value = ochk_manual_nat.snat
}

output "no_snat" {
  value = ochk_manual_nat.no_snat
}

output "rule-ew" {
  value = ochk_firewall_ew_rule.fw-ew1
}

output "rule-sn" {
  value = ochk_firewall_sn_rule.fw-sn1
}