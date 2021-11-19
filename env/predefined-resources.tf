data "ochk_router" "default-vrf" {
  display_name = var.vrf_router
}

resource "ochk_ip_collection" "default" {
  display_name = "${var.test-data-prefix}-ipc"
  ip_addresses = [
    "1.1.1.1",
    "1.0.0.1",
    "8.8.8.8"
  ]
}

resource "ochk_router" "default-vpc" {
  display_name = "vpc-testy"
  parent_router_id = data.ochk_router.default-vrf.id
}

resource "ochk_virtual_network" "subtenant-network" {
  display_name = "${var.test-data-prefix}-vnet"
  subtenants = [
    ochk_subtenant.subtenant-1.id
  ]
  router = ochk_router.default-vpc.id
  ipam_enabled = false
}

resource "ochk_subtenant" "subtenant-1" {
  name = "${var.test-data-prefix}-subt1"
  email = "email1@example.com"
  description = "Business group description 1"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  lifecycle {
    ignore_changes = [
      description
    ]
  }
}

resource "ochk_subtenant" "subtenant-2" {
  name = "${var.test-data-prefix}-subt2"
  email = "email2@example.com"
  description = "Business group description 2"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  lifecycle {
    ignore_changes = [
      description
    ]
  }
}

resource "ochk_subtenant" "subtenant-3" {
  name = "${var.test-data-prefix}-subt3"
  email = "email3@example.com"
  description = "Business group description 3"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  lifecycle {
    ignore_changes = [
      description
    ]
  }
}

resource "ochk_subtenant" "subtenant-4" {
  name = "${var.test-data-prefix}-subt4"
  email = "email4@example.com"
  description = "Business group description 4"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400

  lifecycle {
    ignore_changes = [
      description
    ]
  }
}

resource "ochk_router" "default" {
  display_name = "${var.test-data-prefix}-router"
}

data "ochk_deployment" "centos" {
  display_name = "CentOS 7"
}

resource "ochk_virtual_network" "default" {
  display_name = "${var.test-data-prefix}-vnet3"
  subtenants = [
    ochk_subtenant.subtenant-1.id
  ]
}

resource "ochk_virtual_network" "vnet2" {
  display_name = "${var.test-data-prefix}-vnet4"
  ipam_enabled = true
  subtenants = [
    ochk_subtenant.subtenant-1.id
  ]

  router = ochk_router.default.id
  dns_search_suffix = "test.lcl,prod.lcl"
  dns_suffix = "test.lcl"
  primary_dns_address = "192.168.1.6"
  secondary_dns_address = "192.168.1.2"
  primary_wins_address = "192.168.1.3"
  secondary_wins_address = "192.168.1.3"
  subnet_network_cidr = "192.168.200.0/24"
}

locals {
  ssh_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCbMgMU2dxQYg+WLoim6ZuGsuMZ8QB9mrylNqpbWQrCNXZnajuhjff62E1yMPh7uh2nrLBAFhDu7jOOLPMY8uG7Z9FwutnRQbWsve2uo84FmeLgXcbxg/hD3b9CH5pjqZUjJCCN9DpFveKWVsw+4VvIbTS1m5JcNHXccY3mrUCtTPfUP/W3bRQTFyYtmzX4rV68eoBIUlgNia8DF9sUrgvNVEElaK6gXXjt2UW3aHe6VZ4DUl/MfarWwrY92XL9HwZ81S75Q7NBh75PtnR4ipk8QYNqxoOWsbJB9QnqeURdMgWxciaU3Z1eBTfzLmHXMv2EvqYBcHQ2lMhbFRn/2/an radoslawkubera@NB155.local"
}

resource "ochk_billing_tag" "bt-default" {
    display_name = "${var.test-data-prefix}-billing-tag"
}

resource "ochk_system_tag" "st-default" {
    display_name = "${var.test-data-prefix}-system-tag"
}

resource "ochk_virtual_machine" "default" {
  display_name = "${var.test-data-prefix}-vm"
  deployment_id = data.ochk_deployment.centos.id
  initial_password = "0c80aab0dc1ce6"

  power_state = "poweredOn"
  resource_profile = "SIZE_S"
  storage_policy = "STANDARD"
  subtenant_id = ochk_subtenant.subtenant-1.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.default.id
  }

  ssh_key = local.ssh_key

  billing_tags = [
    ochk_billing_tag.bt-default.id
  ]

  system_tags = [
    ochk_system_tag.st-default.id
  ]
}

resource "ochk_custom_service" "web_servers_https" {
  display_name = "${var.test-data-prefix}-https"

  ports {
    protocol = "TCP"
    source = ["443"]
    destination = ["1-65535"]
  }
}

resource "ochk_custom_service" "web_servers_http" {
  display_name = "${var.test-data-prefix}-http"

  ports {
    protocol = "TCP"
    source = ["80", "8080-8090"]
    destination = ["1-65535"]
  }
}

resource "ochk_kms_key" "aes_key" {
  display_name = "${var.test-data-prefix}-key"
  key_usage = [
    "ENCRYPT",
    "DECRYPT"
  ]
  algorithm = "AES"
  size = "256"
}

