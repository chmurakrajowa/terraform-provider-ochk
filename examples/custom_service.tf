
resource "ochk_custom_service" "web-servers" {
  display_name = "${var.test-data-prefix}-custom-serv"
  ports {
    protocol = "TCP"
    source = ["80", "443"]
    destination = ["80", "443"]
  }
}
