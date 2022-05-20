
resource "ochk_ip_collection" "default" {
  display_name = "${var.test-data-prefix}-ipcol-default"
  ip_addresses = [
    "1.1.1.1",
    "1.0.0.1",
    "8.8.8.8"
  ]
}

