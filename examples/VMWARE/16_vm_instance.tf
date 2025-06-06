# data "ochk_virtual_machines" "virtual_machines_list" {
# }

# data "ochk_virtual_machine" "virtual_machine" {
#     display_name = "vmZielona"
# }


# data "ochk_deployment" "debian" {
#       display_name = "Debian 12"
# }

/*
resource "ochk_virtual_machine" "default-devel" {
  display_name = "${var.test-data-prefix}-vm17up"
  deployment_id = data.ochk_deployment.debian.id
  initial_password = var.initial_password_for_vm
  os_type = "LINUX"
  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  virtual_disk {
      size_mb = 40960
   }
  project_id = "566CF9F6-6633-443F-9F3B-24B41C0689F8"

   virtual_network_devices {
      virtual_network_id = "5C83E86F-8D9A-4C7E-BEA1-10D69EE14883"
    }

  additional_virtual_disks {
    lun_id=1
    size_mb=1000
  }

  additional_virtual_disks {
    lun_id=2
    size_mb=2000
  }

#   tags = [
#     345
#   ]

#   backup_lists = [
#     "5403439c-38a5-4f98-a58b-134072260bfb"
#   ]
}


resource "ochk_virtual_machine" "default" {
  display_name = "${var.test-data-prefix}-vm-tf"
  deployment_id = data.ochk_deployment.debian.id
  initial_password = var.initial_password_for_vm

  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  virtual_disk {
      size_mb = 40960
   }
  project_id = "3cda830c-b37f-46dc-be54-f649d31bec66"

   virtual_network_devices {
      virtual_network_id = "0b6fa626-c969-41a9-83bd-101e915ddfea"
    }

  additional_virtual_disks {
    lun_id=1
    size_mb=1024
  }

#   additional_virtual_disks {
#     lun_id=2
#     size_mb=2048
#   }


  tags = [
    345
  ]

#   backup_lists = [
#     "5403439c-38a5-4f98-a58b-134072260bfb"
#   ]
 # encryption_recrypt = "NONOE"
  encryption = true
  encryption_key_id = "e8a2f47c-d4a7-4eca-99c5-309fa02a67dc"
}
*/

/*
resource "ochk_virtual_machine" "default-oscd2" {
  display_name = "${var.test-data-prefix}-vm-tf1up"
  deployment_id = data.ochk_deployment.debian.id
  initial_password = var.initial_password_for_vm

  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  virtual_disk {
      size_mb = 40960
   }
  project_id = "3cda830c-b37f-46dc-be54-f649d31bec66"

   virtual_network_devices {
      virtual_network_id = "0b6fa626-c969-41a9-83bd-101e915ddfea"
    }

  additional_virtual_disks {
    lun_id=1
    size_mb=1024
  }

  additional_virtual_disks {
    lun_id=2
    size_mb=2048
  }

#   tags = [
#     345
#   ]
}

*/

/*
data "ochk_deployment_types" "dep_types" {
}

data "ochk_deployment_params_types" "dep_types" {
}

data "ochk_deployments" "ovf" {
  deployment_type = "ISO"
}*/

/*
data "ochk_virtual_machine" "test" {
  display_name = data.ochk_virtual_machines.virtual_machines.virtual_machines[0].display_name
}
*/
/*
data "ochk_deployment" "debian11" {
  display_name = "Debian 11"
}*/

/*
resource "ochk_virtual_machine" "default" {
  display_name = "${var.test-data-prefix}-vm"
  deployment_id = data.ochk_deployment.debian11.id
  initial_password = var.initial_password_for_vm

  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  project_id = data.ochk_project.project_for_vm.id
  folder_path = "/path"
  virtual_disk {
    size_mb = 40960
  }
  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.network_for_vm.id
  }
}
*/

/*
resource "ochk_virtual_machine" "default" {
  display_name = "${var.test-data-prefix}-vm"
  deployment_id = "6bbd31ee-9144-4f14-af36-549065112cb9"
  initial_password = var.initial_password_for_vm

  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  project_id = data.ochk_project.project_for_vm.id

  virtual_network_devices {
    virtual_network_id = "b6529fce-d950-4a73-9915-7dc031293196"
  }

  deployment_params {
    param_name="hostname"
    param_type="HOST_NAME"
    param_value="myshost1"
  }

   deployment_params {
     param_name="hostname"
     param_type="NET_DNS_PRIMARY"
     param_value="192.168.15.100"
   }

  additional_virtual_disks {
    lun_id=1
    size_mb=1000
  }

  additional_virtual_disks {
    lun_id=2
    size_mb=2000
  }

#   tags = [
#     345
#   ]

  backup_lists = [
    "5403439c-38a5-4f98-a58b-134072260bfb"
  ]
}

/*
data "ochk_deployment" "debian11" {
  display_name = "Debian 11"
}
*/
/*
locals {
  # should not really be stored like that
  ssh_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCbMgMU2dxQYg+WLoim6ZuGsuMZ8QB9mrylNqpbWQrCNXZnajuhjff62E1yMPh7uh2nrLBAFhDu7jOOLPMY8uG7Z9FwutnRQbWsve2uo84FmeLgXcbxg/hD3b9CH5pjqZUjJCCN9DpFveKWVsw+4VvIbTS1m5JcNHXccY3mrUCtTPfUP/W3bRQTFyYtmzX4rV68eoBIUlgNia8DF9sUrgvNVEElaK6gXXjt2UW3aHe6VZ4DUl/MfarWwrY92XL9HwZ81S75Q7NBh75PtnR4ipk8QYNqxoOWsbJB9QnqeURdMgWxciaU3Z1eBTfzLmHXMv2EvqYBcHQ2lMhbFRn/2/an radoslawkubera@NB155.local"
}

resource "ochk_virtual_machine" "ssh-key" {
  display_name = "${var.test-data-prefix}-vm-sk"
  deployment_id = data.ochk_deployment.debian11.id
  initial_password = var.initial_password_for_vm

  ssh_key = local.ssh_key

  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  project_id = data.ochk_project.project_for_vm.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.network_for_vm.id
  }

  backup_lists = [
    data.ochk_backup_list.backup_list.id
  ]
}
*/

/*
resource "ochk_virtual_machine" "tags" {
  display_name = "${var.test-data-prefix}-vm-tag"
  deployment_id = data.ochk_deployment.debian11.id
  initial_password = var.initial_password_for_vm

  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  project_id = data.ochk_project.project_for_vm.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.network_for_vm.id
  }

#   tags = [
#     345
#   ]
}
*/

/*
resource "ochk_virtual_machine" "backup-list" {
  display_name = "${var.test-data-prefix}-vm-bl"
  deployment_id = data.ochk_deployment.debian11.id
  initial_password = var.initial_password_for_vm

  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  project_id = data.ochk_project.project_for_vm.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.network_for_vm.id
  }

  backup_lists = [
    data.ochk_backup_list.backup_list.id
  ]
}

*/

/*
resource "ochk_virtual_machine" "default" {
  display_name = "${var.test-data-prefix}-vm"
  deployment_id = data.ochk_deployment.debian11.id
  initial_password = var.initial_password_for_vm

  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  project_id = ochk_project.project-1.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.network_for_vm.id
  }

  additional_virtual_disks {
    lun_id=1
    size_mb=1000
  }

    additional_virtual_disks {
    lun_id=2
    size_mb=2000
  }
}
*/

/*
data "ochk_deployment" "archlinux" {
  display_name = var.iso_image
}

resource "ochk_virtual_machine" "iso" {
  display_name = "${var.test-data-prefix}-vm-iso"
  deployment_id = data.ochk_deployment.archlinux.id
  os_type = "WINDOWS"

  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  project_id = data.ochk_project.project_for_vm.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.network_for_vm.id
  }
}

data "ochk_deployment" "ovf" {
  display_name = var.ovf_image
}

resource "ochk_virtual_machine" "ovf" {
  display_name = "${var.test-data-prefix}-vm-ovf"
  deployment_id = data.ochk_deployment.ovf.id

  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  project_id = data.ochk_project.project_for_vm.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.network_for_vm.id
  }
  initial_user_name = "root"
  initial_password = var.initial_password_for_vm
  ovf_ip_configuration = false
  os_type = "LINUX"
}
*/

/*
resource "ochk_virtual_machine" "encrypted" {
  display_name = "${var.test-data-prefix}-vm-en"
  deployment_id = data.ochk_deployment.debian11.id
  initial_password = var.initial_password_for_vm

  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  project_id = data.ochk_project.project_for_vm.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.network_for_vm.id
  }

  encryption = true
  encryption_key_id = ochk_kms_key.aes-generated.id
}
*/

/*
resource "ochk_virtual_machine" "encrypted-with-own-encrypted-key" {
  display_name = "${var.test-data-prefix}-vm-eb"
  deployment_id = data.ochk_deployment.debian11.id
  initial_password = var.initial_password_for_vm

  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  project_id = data.ochk_project.project_for_vm.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.network_for_vm.id
  }

  encryption = true
  encryption_key_id = ochk_kms_key.aes-enc.id
}


resource "ochk_virtual_machine" "encrypted-managed2" {
  display_name = "${var.test-data-prefix}-vm-em"
  deployment_id = data.ochk_deployment.debian11.id
  initial_password = var.initial_password_for_vm

  power_state = "poweredOn"
  memory_size_mb = 4096
  cpu_count = 2
  storage_policy = "STANDARD_W1"
  project_id = data.ochk_project.project_for_vm.id

  virtual_network_devices {
    virtual_network_id =ochk_virtual_network.network_for_vm.id
  }

  encryption = true
}
*/