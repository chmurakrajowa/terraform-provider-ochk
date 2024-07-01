/*data "ochk_vrf" "vrf" {
	display_name = var.vrf_router
}

data "ochk_firewall_ew_rules" "firewall_ew_rules" {
    vpc_id = ochk_vpc.project-vpc.id
}*/

/*
resource "ochk_firewall_ew_rule" "fw-ew2" {
  display_name = "${var.test-data-prefix}-tf-fw-ew-http"
  vpc_id = ochk_vpc.project-vpc.id
  services = [data.ochk_service.http.id]
  #custom_services = [data.ochk_custom_service.web-servers.id]

  source = [ochk_security_group.project-test.id]
  destination = [ochk_security_group.project-sg1-dst.id]

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
  vpc_id = data.ochk_vpc.project-vpc1234.id
  services = [data.ochk_service.http.id]
  #custom_services = [data.ochk_custom_service.web-servers.id]

  source = [ochk_security_group.project-sg1-src.id]
  destination = [ochk_security_group.project-sg1-dst.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  priority = 10
}

resource "ochk_firewall_ew_rule" "fw-ew3" {
  display_name = "${var.test-data-prefix}-tf-fw-ew-http3"
  vpc_id = data.ochk_vpc.project-vpc1234.id
  services = [data.ochk_service.http.id]
  #custom_services = [data.ochk_custom_service.web-servers.id]

  source = [ochk_security_group.project-sg1-src.id]
  destination = [ochk_security_group.project-sg1-dst.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  priority = 20
}

resource "ochk_firewall_ew_rule" "fw-ew4" {
  display_name = "${var.test-data-prefix}-tf-fw-ew-http4"
  vpc_id = data.ochk_vpc.project-vpc1234.id
  services = [data.ochk_service.http.id]
  #custom_services = [data.ochk_custom_service.web-servers.id]

  source = [ochk_security_group.project-sg1-src.id]
  destination = [ochk_security_group.project-sg1-dst.id]

  action = "ALLOW"
  ip_protocol = "IPV4_IPV6"

  priority = 30
}
*/