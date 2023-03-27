
resource "ochk_ip_collection" "default" {
  display_name = "${var.test-data-prefix}-ipcol-default"
  project_id = data.ochk_project.project_for_vm.id
  ip_addresses = [
    "1.1.1.1",
    "1.0.0.1",
    "8.8.8.8"
  ]
}

data "ochk_ip_collections" "ip_collections" {

}

