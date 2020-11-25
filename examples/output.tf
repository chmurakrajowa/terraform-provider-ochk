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
  value = data.ochk_network.subtenant-network
}

output "user1" {
  value = data.ochk_user.default
}

output "vnet" {
  value = ochk_virtual_network.default
}

output "vm" {
  value = ochk_virtual_machine.default
}

