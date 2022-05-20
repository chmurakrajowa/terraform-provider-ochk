
/*
data "ochk_manual_nat" "manual-nat1" {
  display_name = "${var.test-data-prefix}-dnat1"
}

data "ochk_manual_nat" "manual-nat" {
  display_name = "${var.test-data-prefix}-dnat"
}
*/
data "ochk_service" "http" {
  display_name = "HTTP"
}

resource "ochk_manual_nat" "dnat" {
    display_name = "${var.test-data-prefix}-dnat"
    action = "DNAT"
    enabled = true
    vrf_id = data.ochk_router.subtenant-vrf.id
    source_network = "10.10.0.10"
    translated_network = "10.12.12.12"
    destination_network = "45.130.209.206"
    priority = 115
    service_id = data.ochk_service.http.id
    translated_ports = "80"
}

resource "ochk_manual_nat" "dnat2" {
    display_name = "${var.test-data-prefix}-dnat2"
    action = "DNAT"
    enabled = true
    vrf_id = data.ochk_router.subtenant-vrf.id
    source_network = "10.10.0.10"
    translated_network = "10.12.12.12"
    destination_network = "45.130.209.206"
    priority = 126
}

resource "ochk_manual_nat" "snat" {
    display_name = "${var.test-data-prefix}-snat2"
    action = "SNAT"
    enabled = true
    vrf_id = data.ochk_router.subtenant-vrf.id
    source_network = "10.10.0.10"
    translated_network = "45.130.209.206"
    destination_network = "10.12.12.12"
    priority = 118
}

resource "ochk_manual_nat" "no_snat" {
    display_name = "${var.test-data-prefix}-no-snat2"
    action = "NO_SNAT"
    enabled = true
    vrf_id = data.ochk_router.subtenant-vrf.id
    source_network = "10.10.0.10"
    destination_network = "10.12.12.12"
    priority = 119
}
