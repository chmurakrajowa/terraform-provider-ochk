data "ochk_deployment" "centos" {
  display_name = "CentOS 7"
}

resource "ochk_virtual_machine" "default" {
  display_name = "${var.tenant}-VM-instance-${random_string.random}"
  deployment_id = data.ochk_deployment.centos.id
  initial_password = random_string.random

  power_state = "poweredOn"
  resource_profile = "SIZE_XS"
  storage_policy = "STANDARD"
  subtenant_id = data.ochk_subtenant.subtenant.id

  virtual_network_devices {
    virtual_network_id = ochk_virtual_network.default.id
  }
}
