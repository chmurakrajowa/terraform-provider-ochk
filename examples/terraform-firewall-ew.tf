resource "ochk_firewall_ew_rule" "fw-ew1" {
  display_name = "tf-${var.demo-id}-fw-ew1-updated"
  security_policy_id = data.ochk_security_policy.default.id
  services = [data.ochk_service.http.id]

  source = [ochk_security_group.vm.id]
  destination = [ochk_security_group.ipset.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"
}

resource "ochk_firewall_ew_rule" "fw-ew2" {
  display_name = "tf-${var.demo-id}-fw-ew2"
  security_policy_id = data.ochk_security_policy.default.id
  services = [data.ochk_service.http.id]

  source = [ochk_security_group.vm_ipset.id]
  destination = [ochk_security_group.vm.id]

  action = "ALLOW" // default
  ip_protocol = "IPV4_IPV6" // default

  position {
    rule_id = ochk_firewall_ew_rule.fw-ew1.id
    revise_operation = "BEFORE"
  }

  lifecycle {
    ignore_changes = [position]
  }
}




