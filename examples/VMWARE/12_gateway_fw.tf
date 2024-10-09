/*
data "ochk_service" "http_gw" {
  display_name = "HTTP"
}

data "ochk_firewall_sn_rules" "firewall_sn_rules" {
    vpc_id = ochk_vpc.project-vpc.id
}


resource "ochk_firewall_sn_rule" "fw-ew3" {
  display_name = "${var.test-data-prefix}-tf-fw-sn-http"
  vpc_id = ochk_vpc.project-vpc.id
  project_id = data.ochk_project.project_for_vm.id
  services = [data.ochk_service.http_gw.id]
  custom_services = [ochk_custom_service.web-servers.id]

  source = [ochk_security_group.project-sg1-src.id]
  destination = [ochk_security_group.project-sg1-dst.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  priority = 1000
}*/
