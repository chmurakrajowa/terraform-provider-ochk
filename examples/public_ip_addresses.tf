data "ochk_public_ip_addresses" "public_ip_addresses" {
}

output "public_ip_addresses" {
  value = data.ochk_public_ip_addresses.public_ip_addresses
}