data "ochk_gateway_policy" "T1" {
  display_name = ochk_router.T1.display_name
  depends_on = [ochk_router.T1]
}

resource "ochk_firewall_sn_rule" "fw-sn-1" {
  display_name = "tf-${var.demo-id}-fw-sn-1-ssh-drop"
  gateway_policy_id = data.ochk_gateway_policy.T1.id
  scope = [ochk_router.T1.id]

  services = [data.ochk_service.ssh.id]

  source = [ochk_security_group.vm_ipc.id]
  destination = [ochk_security_group.vm.id]

  action = "DROP"
  ip_protocol = "IPV4"
}