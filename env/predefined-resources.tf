resource "ochk_project" "project-1" {
  display_name = "${var.test-data-prefix}-project-01"
  description = "Project description 1"
  memory_reserved_size_mb = 255000
  storage_reserved_size_gb = 10000
  vcpu_reserved_quantity = 100
  vrf_id = data.ochk_vrf.default-vrf.id
  lifecycle {
    ignore_changes = [
      description
    ]
  }
}

resource "ochk_project" "project-2" {
  display_name = "${var.test-data-prefix}-project-02"
  description = "Project description 2"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  vcpu_reserved_quantity = 100
  vrf_id = data.ochk_vrf.default-vrf.id
}

resource "ochk_project" "project-3" {
  display_name = "${var.test-data-prefix}-project-03"
  description = "Project description 3"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  vcpu_reserved_quantity = 100
  vrf_id = data.ochk_vrf.default-vrf.id
}

resource "ochk_project" "project-4" {
  display_name = "${var.test-data-prefix}-project-04"
  description = "Project description 4"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  vcpu_reserved_quantity = 100
  vrf_id = data.ochk_vrf.default-vrf.id
}

resource "ochk_ip_collection" "default" {
  display_name = "${var.test-data-prefix}-ipc-default"
  project_id = ochk_project.project-1.id
  ip_addresses = [
    "1.1.1.1",
    "1.0.0.1",
    "8.8.8.8"
  ]
}

resource "ochk_vpc" "default-vpc" {
  display_name = "${var.test-data-prefix}-vpc-test"
  vrf_id = data.ochk_vrf.default-vrf.id
  project_id = ochk_project.project-1.id
}

resource "ochk_virtual_network" "project-network" {
  display_name = "${var.test-data-prefix}-vnet-proj"
  vpc_id = ochk_vpc.default-vpc.id
  subnet_network_cidr = "192.168.254.0/24"
  ipam_enabled = false
  project_id = ochk_project.project-1.id
}

resource "ochk_vpc" "default" {
  display_name = "${var.test-data-prefix}-vpc"
  vrf_id = data.ochk_vrf.default-vrf.id
  project_id = ochk_project.project-1.id
}

resource "ochk_virtual_network" "default" {
  display_name = "${var.test-data-prefix}-vnet"
  project_id = ochk_project.project-1.id
}

resource "ochk_virtual_network" "vnet1" {
  display_name = "${var.test-data-prefix}-vnet1"
  ipam_enabled = false
  project_id = ochk_project.project-1.id

  vpc_id = ochk_vpc.default.id
  subnet_network_cidr = "192.168.201.0/24"
}

resource "ochk_virtual_network" "vnet2" {
  display_name = "${var.test-data-prefix}-vnet2"
  ipam_enabled = true
  project_id = ochk_project.project-1.id

  vpc_id = ochk_vpc.default.id
  subnet_network_cidr = "192.168.200.0/24"
}

resource "ochk_virtual_network" "vnet3" {
  display_name = "${var.test-data-prefix}-vnet3"
  ipam_enabled = false
  project_id = ochk_project.project-1.id

  vpc_id = ochk_vpc.default.id
  subnet_network_cidr = "192.168.202.0/24"
}

resource "ochk_virtual_network" "vnet4" {
  display_name = "${var.test-data-prefix}-vnet4"
  ipam_enabled = false
  project_id = ochk_project.project-1.id

  vpc_id = ochk_vpc.default.id
  subnet_network_cidr = "192.167.204.0/24"
}

locals {
  ssh_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCbMgMU2dxQYg+WLoim6ZuGsuMZ8QB9mrylNqpbWQrCNXZnajuhjff62E1yMPh7uh2nrLBAFhDu7jOOLPMY8uG7Z9FwutnRQbWsve2uo84FmeLgXcbxg/hD3b9CH5pjqZUjJCCN9DpFveKWVsw+4VvIbTS1m5JcNHXccY3mrUCtTPfUP/W3bRQTFyYtmzX4rV68eoBIUlgNia8DF9sUrgvNVEElaK6gXXjt2UW3aHe6VZ4DUl/MfarWwrY92XL9HwZ81S75Q7NBh75PtnR4ipk8QYNqxoOWsbJB9QnqeURdMgWxciaU3Z1eBTfzLmHXMv2EvqYBcHQ2lMhbFRn/2/an radoslawkubera@NB155.local"
}

resource "ochk_tag" "default" {
    display_name = "${var.test-data-prefix}-tag"
    project_id = ochk_project.project-1.id
}

resource "ochk_virtual_machine" "default" {
  display_name = "${var.test-data-prefix}-vm-default"
  deployment_id = data.ochk_deployment.centos.id
  initial_password = "0c80aab0dc1ce6"

  power_state = "poweredOn"
  cpu_count = 2
  memory_size_mb = 4096
  virtual_disk {
    size_mb = 40960
  }
  storage_policy = "STANDARD_W1"
  project_id = ochk_project.project-1.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.default.id
  }

  ssh_key = local.ssh_key

  tags = [
    ochk_tag.default.id
  ]
}

resource "ochk_virtual_machine" "vm1" {
  display_name = "${var.test-data-prefix}-vm1"
  deployment_id = data.ochk_deployment.centos.id
  initial_password = "0c80aab0dc1ce6"

  power_state = "poweredOn"
  cpu_count = 2
  memory_size_mb = 4096
  virtual_disk {
    size_mb = 40960
  }
  storage_policy = "STANDARD_W1"
  project_id = ochk_project.project-1.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.default.id
  }
}

resource "ochk_virtual_machine" "vm2" {
  display_name = "${var.test-data-prefix}-vm2"
  deployment_id = data.ochk_deployment.centos.id
  initial_password = "0c80aab0dc1ce6"

  power_state = "poweredOn"
  cpu_count = 2
  memory_size_mb = 4096
  virtual_disk {
    size_mb = 40960
  }
  storage_policy = "STANDARD_W1"
  project_id = ochk_project.project-1.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.default.id
  }
}

resource "ochk_virtual_machine" "vm3" {
  display_name = "${var.test-data-prefix}-vm3"
  deployment_id = data.ochk_deployment.centos.id
  initial_password = "0c80aab0dc1ce6"

  power_state = "poweredOn"
  cpu_count = 8
  memory_size_mb = 4096
  virtual_disk {
    size_mb = 40960
  }
  storage_policy = "STANDARD_W1"
  project_id = ochk_project.project-1.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.default.id
  }
}

resource "ochk_virtual_machine" "vm4" {
  display_name = "${var.test-data-prefix}-vm4"
  deployment_id = data.ochk_deployment.centos.id
  initial_password = "0c80aab0dc1ce6"

  power_state = "poweredOn"
  cpu_count = 8
  memory_size_mb = 4096
  virtual_disk {
    size_mb = 40960
  }

  storage_policy = "STANDARD_W1"
  project_id = ochk_project.project-1.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.default.id
  }
}

resource "ochk_custom_service" "web_servers_https" {
  display_name = "${var.test-data-prefix}-https"
  project_id = ochk_project.project-1.id
  ports {
    protocol = "TCP"
    source = ["443"]
    destination = ["1-65535"]
  }
}

resource "ochk_custom_service" "web_servers_http" {
  display_name = "${var.test-data-prefix}-http"
  project_id = ochk_project.project-1.id
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

resource "ochk_tag" "res-bt1" {
    display_name = "${var.test-data-prefix}-t1"
    project_id = ochk_project.project-1.id
}

resource "ochk_auto_nat" "auto-nat" {
    display_name = "${var.test-data-prefix}-autonat"
    virtual_network_id = ochk_virtual_network.vnet3.id
}

resource "ochk_manual_nat" "dnat" {
    display_name = "${var.test-data-prefix}-dnat"
    action = "DNAT"
    enabled = false
    vrf_id =  data.ochk_vrf.default-vrf.id
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
    vrf_id = data.ochk_vrf.default-vrf.id
    source_network = "10.10.0.10"
    translated_network = local.nat_public_ip_addr
    destination_network = "10.12.12.12"
    priority = 918
}

resource "ochk_manual_nat" "no_snat" {
    display_name = "${var.test-data-prefix}-no-snat2"
    action = "NO_SNAT"
    enabled = true
    vrf_id = data.ochk_vrf.default-vrf.id
    source_network = "10.10.0.10"
    destination_network = "10.12.12.12"
    priority = 919
}

resource "ochk_firewall_ew_rule" "fw-ew1" {
  display_name = "${var.test-data-prefix}-tf-fw-ew-http"
  vpc_id = ochk_vpc.default.id
  services = [data.ochk_service.http.id]
  project_id = ochk_project.project-1.id

  source = [ochk_security_group.subtenant-sg1-src.id]
  destination = [ochk_security_group.subtenant-sg1-dst.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  priority = 10
}

resource "ochk_security_group" "subtenant-sg1-src" {
  display_name = "${var.test-data-prefix}-sg1-src"
  project_id = ochk_project.project-1.id
  members {
    id = ochk_ip_collection.default.id
    type = "IPCOLLECTION"
  }
}

resource "ochk_security_group" "subtenant-sg1-dst" {
  display_name = "${var.test-data-prefix}-sg1-dst"
  project_id = ochk_project.project-1.id
  members {
    id = ochk_ip_collection.default.id
    type = "IPCOLLECTION"
  }
}

resource "ochk_security_group" "subtenant-sg-vm" {
  display_name = "${var.test-data-prefix}-sg-vm"
  project_id = ochk_project.project-1.id
  members {
    id = ochk_virtual_machine.default.id
    type = "VIRTUAL_MACHINE"
  }
}

resource "ochk_firewall_sn_rule" "fw-sn1" {
  display_name = "${var.test-data-prefix}-tf-fw-sn-http"
  vpc_id = ochk_vpc.default.id
  project_id = ochk_project.project-1.id
  services = [data.ochk_service.http.id]

  source = [ochk_security_group.subtenant-sg1-src.id]
  destination = [ochk_security_group.subtenant-sg1-dst.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  priority = 10
}
