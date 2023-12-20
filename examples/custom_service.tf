/*
resource "ochk_custom_service" "web-servers" {
  display_name = "${var.test-data-prefix}-custom-serv"
  project_id = data.ochk_project.project_for_vm.id
  ports {
    protocol = "TCP"
    source = ["80", "443"]
    destination = ["80", "443"]
  }
}

data "ochk_custom_services" "custom_services" {
}*/

/*
data "ochk_custom_service" "custom_service" {
  display_name = "${var.test-data-prefix}-custom-serv"
}
*/