/*data "ochk_firewall_rule" "list_rules" {
    project_id = "9a23ad90-3d9f-48d8-b4e2-788db66d90f7"
    security_group_id = "e8a7947a-821a-4352-9e2f-a47306cc80b8"
    display_name = "2E38B351-2746-4FDE-B44B-5DD9AF5D585A"
}*/

/*
data "ochk_firewall_rule" "default2" {
    display_name = "a-acct-06j36"
    project_id = "3cda830c-b37f-46dc-be54-f649d31bec66"
    security_group_id = "ff765340-2c9e-4dc8-b565-bd1d84793f5d"
}
*/
/*
resource "ochk_firewall_rule" "default3" {
    display_name = "${var.test-data-prefix}-test05upd"
    project_id = "0f00844f-9da4-40f5-860e-98f0e24711a1"
    security_group_id = "1f59e0d4-94c2-4944-b1b0-62a378dabcb3"
    description = "TESTETST05upd"
     ether_type = "IPv4"
     direction = "EGRESS"
     protocol = "TCP"
     port_range_min = 20008
     port_range_max = 20020
 //    remote_ip_prefix = "200.200.200.1/32"
    dest_security_group = "8dabcc3a-f470-4314-826c-88b2a7c0d2d9"
}
*/
resource "ochk_firewall_rule" "default4" {
    display_name = "${var.test-data-prefix}-test02updtest"
    project_id = "9a23ad90-3d9f-48d8-b4e2-788db66d90f7"
    security_group_id = "8dabcc3a-f470-4314-826c-88b2a7c0d2d9"
    description = "TESTETST01"
     ether_type = "IPv4"
     direction = "EGRESS"
     protocol = "TCP"
     port_range_min = 20008
     port_range_max = 20020
     dest_security_group = "cad26944-7dd5-458d-8c30-3115f786ba95"
}