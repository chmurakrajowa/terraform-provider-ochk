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
