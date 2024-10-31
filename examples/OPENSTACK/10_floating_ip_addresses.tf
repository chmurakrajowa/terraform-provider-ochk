data "ochk_floating_ip_addresses" "default" {
}

data "ochk_floating_ip_address" "default" {
    display_name = "publicIp1"
}


resource "ochk_floating_ip_address" "tj1" {
    display_name = "XXXXXXXX"
    description = "XXXXXXXX"
    vm_port_id = "d3c6bda6-172d-408b-a3ce-4f025f467e31" // TODO dodać endpoint do proxy aby można było pobrać porty vm przez providera
}

