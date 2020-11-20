resource "ochk_ip_collection" "local24" {
  display_name = "tf-ipc-${var.demo-id}-local24"
  ip_addresses = ["192.168.0.1/22"]
}

resource "ochk_ip_collection" "ranges" {
  display_name = "tf-ipc-${var.demo-id}-ranges"
  ip_addresses = ["10.10.10.0-10.10.10.100"]
}

resource "ochk_ip_collection" "dns-servers" {
  display_name = "tf-ipc-${var.demo-id}-dns-servers"
  ip_addresses = ["1.1.1.1", "1.0.0.1", "8.8.8.8"]
}