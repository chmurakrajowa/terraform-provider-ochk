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

