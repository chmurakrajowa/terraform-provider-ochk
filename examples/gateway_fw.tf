
data "ochk_service" "http_gw" {
  display_name = "HTTP"
}


resource "ochk_firewall_sn_rule" "fw-ew3" {
  display_name = "${var.test-data-prefix}-tf-fw-sn-http"
  router_id = ochk_router.subtenant-vpc.id
  services = [data.ochk_service.http_gw.id]
  custom_services = [ochk_custom_service.web-servers.id]

  source = [ochk_security_group.subtenant-sg1-src.id]
  destination = [ochk_security_group.subtenant-sg1-dst.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  priority = 1000
}
