/*data "ochk_firewall_rules" "list_rules" {
    project_id = "3cda830c-b37f-46dc-be54-f649d31bec66"
    security_group_id = "ff765340-2c9e-4dc8-b565-bd1d84793f5d"
}*/

/*
data "ochk_firewall_rule" "default2" {
    display_name = "a-acct-06j36"
    project_id = "3cda830c-b37f-46dc-be54-f649d31bec66"
    security_group_id = "ff765340-2c9e-4dc8-b565-bd1d84793f5d"
}
*/
/*resource "ochk_firewall_rule" "default3" {
    display_name = "${var.test-data-prefix}-test05"
    project_id = "3cda830c-b37f-46dc-be54-f649d31bec66"
    security_group_id = "ff765340-2c9e-4dc8-b565-bd1d84793f5d"
    description = "TESTETST05upd"
     ether_type = "IPv4"
     direction = "EGRESS"
     protocol = "TCP"
     port_range_min = 20008
     port_range_max = 20020
     remote_ip_prefix = ""
}*/