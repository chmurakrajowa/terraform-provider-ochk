
resource "ochk_ip_collection" "default" {
  display_name = "${var.test-data-prefix}-ipc-default"
  ip_addresses = [
    "1.1.1.1",
    "1.0.0.1",
    "8.8.8.8"
  ]
}

resource "ochk_router" "default-vpc" {
  display_name = "${var.test-data-prefix}-vpc-testy"
  parent_router_id = data.ochk_router.default-vrf.id
}

resource "ochk_virtual_network" "subtenant-network" {
  display_name = "${var.test-data-prefix}-vnet-subt"
  subtenants = [
    ochk_subtenant.subtenant-1.id
  ]
  router = ochk_router.default-vpc.id
  subnet_network_cidr = "192.168.254.0/24"
  ipam_enabled = false
}

resource "ochk_subtenant" "subtenant-1" {
  name = "${var.test-data-prefix}-subt1"
  email = "email1@example.com"
  description = "Business group description 1"
  memory_reserved_size_mb = 255000
  storage_reserved_size_gb = 1000
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
  parent_router_id = data.ochk_router.default-vrf.id
}

resource "ochk_virtual_network" "default" {
  display_name = "${var.test-data-prefix}-vnet"
  subtenants = [
    ochk_subtenant.subtenant-1.id
  ]
}

resource "ochk_virtual_network" "vnet1" {
  display_name = "${var.test-data-prefix}-vnet1"
  ipam_enabled = false
  subtenants = [
    ochk_subtenant.subtenant-1.id
  ]

  router = ochk_router.default.id
  subnet_network_cidr = "192.168.201.0/24"
}

resource "ochk_virtual_network" "vnet2" {
  display_name = "${var.test-data-prefix}-vnet2"
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

resource "ochk_virtual_network" "vnet3" {
  display_name = "${var.test-data-prefix}-vnet3"
  ipam_enabled = false
  subtenants = [
    ochk_subtenant.subtenant-1.id
  ]

  router = ochk_router.default.id
  subnet_network_cidr = "192.168.202.0/24"
}

resource "ochk_virtual_network" "vnet4" {
  display_name = "${var.test-data-prefix}-vnet4"
  ipam_enabled = false
  subtenants = [
    ochk_subtenant.subtenant-1.id
  ]

  router = ochk_router.default.id
  subnet_network_cidr = "192.168.203.0/24"
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
  storage_policy = "STANDARD_W1"
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

resource "ochk_virtual_machine" "vm1" {
  display_name = "${var.test-data-prefix}-vm1"
  deployment_id = data.ochk_deployment.centos.id
  initial_password = "0c80aab0dc1ce6"

  power_state = "poweredOn"
  resource_profile = "SIZE_S"
  storage_policy = "STANDARD_W1"
  subtenant_id = ochk_subtenant.subtenant-1.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.default.id
  }
}

resource "ochk_virtual_machine" "vm2" {
  display_name = "${var.test-data-prefix}-vm2"
  deployment_id = data.ochk_deployment.centos.id
  initial_password = "0c80aab0dc1ce6"

  power_state = "poweredOn"
  resource_profile = "SIZE_S"
  storage_policy = "STANDARD_W1"
  subtenant_id = ochk_subtenant.subtenant-1.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.default.id
  }
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

/*
resource "ochk_user" "test-user1" {
    name = "${var.test-data-prefix}-user1"
    display_name="${var.test-data-prefix}-user1"
    password="zaq1@WSXzaq1@WSX"
    email_address="${var.test-data-prefix}-user1@ochk.pl"
    first_name="Firstname"  
    last_name="Lastname"
    enabled=true
}

resource "ochk_user" "test-user2" {
    name = "${var.test-data-prefix}-user2"
    display_name="${var.test-data-prefix}-user2"
    password="zaq1@WSXzaq1@WSX"
    email_address="${var.test-data-prefix}-user2@ochk.pl"
    first_name="Firstname"  
    last_name="Lastname"
    enabled=true
}
*/

resource "ochk_billing_tag" "res-bt1" {
    display_name = "${var.test-data-prefix}-billing-t1"
}

resource "ochk_system_tag" "res-st1" {
    display_name = "${var.test-data-prefix}-system-t1"
}

resource "ochk_auto_nat" "auto-nat" {
    display_name = "${var.test-data-prefix}-autonat"
    virtual_network_id = ochk_virtual_network.vnet3.id
}

resource "ochk_manual_nat" "dnat" {
    display_name = "${var.test-data-prefix}-dnat"
    action = "DNAT"
    enabled = true
    vrf_id =  data.ochk_router.default-vrf.id
    source_network = "10.10.0.10"
    translated_network = "10.12.12.12"
    destination_network = local.nat_public_ip_addr
    priority = 915
    service_id = data.ochk_service.http.id
    translated_ports = "80"
}

resource "ochk_manual_nat" "snat" {
    display_name = "${var.test-data-prefix}-snat2"
    action = "SNAT"
    enabled = true
    vrf_id = data.ochk_router.default-vrf.id
    source_network = "10.10.0.10"
    translated_network = local.nat_public_ip_addr
    destination_network = "10.12.12.12"
    priority = 918
}

resource "ochk_manual_nat" "no_snat" {
    display_name = "${var.test-data-prefix}-no-snat2"
    action = "NO_SNAT"
    enabled = true
    vrf_id = data.ochk_router.default-vrf.id
    source_network = "10.10.0.10"
    destination_network = "10.12.12.12"
    priority = 919
}

resource "ochk_firewall_ew_rule" "fw-ew1" {
  display_name = "${var.test-data-prefix}-tf-fw-ew-http"
  router_id = ochk_router.default.id
  services = [data.ochk_service.http.id]

  source = [ochk_security_group.subtenant-sg1-src.id]
  destination = [ochk_security_group.subtenant-sg1-dst.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  priority = 10
}

resource "ochk_security_group" "subtenant-sg1-src" {
  display_name = "${var.test-data-prefix}-sg1-src"

  members {
    id = ochk_ip_collection.default.id
    type = "IPCOLLECTION"
  }
}

resource "ochk_security_group" "subtenant-sg1-dst" {
  display_name = "${var.test-data-prefix}-sg1-dst"

   members {
    id = ochk_ip_collection.default.id
    type = "IPCOLLECTION"
  }
}

resource "ochk_firewall_sn_rule" "fw-sn1" {
  display_name = "${var.test-data-prefix}-tf-fw-sn-http"
  router_id = ochk_router.default.id
  services = [data.ochk_service.http.id]

  source = [ochk_security_group.subtenant-sg1-src.id]
  destination = [ochk_security_group.subtenant-sg1-dst.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  priority = 10
}

