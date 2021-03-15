data "ochk_deployment" "centos" {
  display_name = "CentOS 7"
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

resource "ochk_virtual_machine" "encrypted" {
  display_name = "${var.test-data-prefix}-vme1"
  deployment_id = data.ochk_deployment.centos.id
  initial_password = "0c80aab0dc1ce6"

  power_state = "poweredOn"
  resource_profile = "SIZE_S"
  storage_policy = "STANDARD"
  subtenant_id = data.ochk_subtenant.subtenant_for_vm.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.default.id
  }

  encryption = true
  encryption_key_id = ochk_kms_key.aes-generated.id
}


/*
resource "ochk_virtual_machine" "encrypted-with-own-encrypted-key" {
  display_name = "${var.test-data-prefix}-vm-e2"
  deployment_id = data.ochk_deployment.centos.id
  initial_password = "0c80aab0dc1ce6"

  power_state = "poweredOn"
  resource_profile = "SIZE_S"
  storage_policy = "STANDARD"
  subtenant_id = data.ochk_subtenant.subtenant_for_vm.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.default.id
  }

  encryption = true
  encryption_key_id = ochk_kms_key.aes-enc.id
//  encryption_key_id = ochk_kms_key.aes-generated.id
//  encryption_recrypt = "DEEP"
}*/
//
//
//resource "ochk_virtual_machine" "encrypted-managed2" {
//  display_name = "${var.test-data-prefix}-vme4"
//  deployment_id = data.ochk_deployment.centos.id
//  initial_password = "0c80aab0dc1ce6"
//
//  power_state = "poweredOn"
//  resource_profile = "SIZE_S"
//  storage_policy = "STANDARD"
//  subtenant_id = data.ochk_subtenant.subtenant_for_vm.id
//
//  virtual_network_devices {
//    virtual_network_id = ochk_virtual_network.default.id
//  }
//
//  encryption = true
//}




