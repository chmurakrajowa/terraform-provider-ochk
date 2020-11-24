data "ochk_security_policy" "default" {
  display_name = "devel"
}

data "ochk_user" "default" {
  name = var.test_user
}

resource "ochk_ip_collection" "default" {
  display_name = "${var.test-data-prefix}-ipc-default"
  ip_addresses = [
    "1.1.1.1",
    "1.0.0.1",
    "8.8.8.8"
  ]
}

data "ochk_network" "subtenant-network" {
  name = var.subtenant_network_name
}

resource "ochk_subtenant" "subtenant-1" {
  name = "${var.test-data-prefix}-subtenant-1"
  email = "email1@example.com"
  description = "Business group description 1"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  users = [
    data.ochk_user.default.id
  ]
  networks = [
    data.ochk_network.subtenant-network.id
  ]

  lifecycle {
    ignore_changes = [
      description
    ]
  }
}

resource "ochk_subtenant" "subtenant-2" {
  name = "${var.test-data-prefix}-subtenant-2"
  email = "email2@example.com"
  description = "Business group description 2"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  users = [
    data.ochk_user.default.id
  ]
  networks = [
    data.ochk_network.subtenant-network.id
  ]

  lifecycle {
    ignore_changes = [
      description
    ]
  }
}

resource "ochk_subtenant" "subtenant-3" {
  name = "${var.test-data-prefix}-subtenant-3"
  email = "email3@example.com"
  description = "Business group description 3"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  users = [
    data.ochk_user.default.id
  ]
  networks = [
    data.ochk_network.subtenant-network.id
  ]

  lifecycle {
    ignore_changes = [
      description
    ]
  }
}

resource "ochk_subtenant" "subtenant-4" {
  name = "${var.test-data-prefix}-subtenant-4"
  email = "email4@example.com"
  description = "Business group description 4"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  users = [
    data.ochk_user.default.id
  ]
  networks = [
    data.ochk_network.subtenant-network.id
  ]

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

data "ochk_subtenant" "subtenant_for_vm" {
  name = var.subtenant_for_vm_name
}

resource "ochk_virtual_network" "default" {
  display_name = "${var.test-data-prefix}-vnet3"
  subtenants = [
    data.ochk_subtenant.subtenant_for_vm.id
  ]
}

resource "ochk_virtual_network" "vnet2" {
  display_name = "${var.test-data-prefix}-vnet4"
  ipam_enabled = true
  subtenants = [
    data.ochk_subtenant.subtenant_for_vm.id
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


resource "ochk_virtual_machine" "default" {
  display_name = "${var.test-data-prefix}-vm"
  deployment_id = data.ochk_deployment.centos.id
  initial_password = "0c80aab0dc1ce6"

  power_state = "poweredOn"
  resource_profile = "SIZE_S"
  storage_policy = "STANDARD"
  subtenant_id = data.ochk_subtenant.subtenant_for_vm.id

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

