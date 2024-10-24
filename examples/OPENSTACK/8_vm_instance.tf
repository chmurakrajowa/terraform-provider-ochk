# data "ochk_virtual_machines" "virtual_machines_devel" {
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
   tags = [
     345
   ]
}

*/
/*
data "ochk_virtual_machines" "virtual_machines_openstack" {
}*/

# data "ochk_virtual_machine" "virtual_machines_openstack" {
#     display_name = "vmZielona"
# }

/*
data "ochk_deployment" "debian" {
  display_name = "Debian 12"
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
}

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
  folder_path = "/example"
  virtual_disk {
    size_mb = 40960
  }
  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.network_for_vm.id
  }
}
*/
