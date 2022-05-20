data "ochk_router" "vrf" {
	display_name = "T0"
}

data "ochk_router" "vpc" {
  display_name = "tf-tst-router"
  parent_router_id = data.ochk_router.vrf.id
}

resource "ochk_firewall_ew_rule" "def" {
  display_name = "tf-tst-tf-fw-ew-http"
  router_id = data.ochk_router.vpc.id
  priority = 10
}

/*
resource "ochk_firewall_ew_rule" "fw-ew2" {
  display_name = "${var.test-data-prefix}-tf-fw-ew-http"
  router_id = ochk_router.subtenant-vpc.id
  services = [data.ochk_service.http.id]
  #custom_services = [data.ochk_custom_service.web-servers.id]

  source = [ochk_security_group.subtenant-test.id]
  destination = [ochk_security_group.subtenant-sg1-dst.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  priority = 10
}
*/

/*
data "ochk_service" "http" {
  display_name = "HTTP"
}

#data "ochk_custom_service" "web-servers" {
#  display_name = "web-servers"
#}



resource "ochk_firewall_ew_rule" "fw-ew2" {
  display_name = "${var.test-data-prefix}-tf-fw-ew-http"
  router_id = data.ochk_router.subtenant-vpc1234.id
  services = [data.ochk_service.http.id]
  #custom_services = [data.ochk_custom_service.web-servers.id]

  source = [ochk_security_group.subtenant-sg1-src.id]
  destination = [ochk_security_group.subtenant-sg1-dst.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  priority = 10
}

resource "ochk_firewall_ew_rule" "fw-ew3" {
  display_name = "${var.test-data-prefix}-tf-fw-ew-http3"
  router_id = data.ochk_router.subtenant-vpc1234.id
  services = [data.ochk_service.http.id]
  #custom_services = [data.ochk_custom_service.web-servers.id]

  source = [ochk_security_group.subtenant-sg1-src.id]
  destination = [ochk_security_group.subtenant-sg1-dst.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  priority = 20
}

resource "ochk_firewall_ew_rule" "fw-ew4" {
  display_name = "${var.test-data-prefix}-tf-fw-ew-http4"
  router_id = data.ochk_router.subtenant-vpc1234.id
  services = [data.ochk_service.http.id]
  #custom_services = [data.ochk_custom_service.web-servers.id]

  source = [ochk_security_group.subtenant-sg1-src.id]
  destination = [ochk_security_group.subtenant-sg1-dst.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  priority = 30
}
*/