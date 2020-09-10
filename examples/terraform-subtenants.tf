resource "ochk_subtenant" "bg-1" {
  name = "tf-bg-${var.demo-id}"
  email = "email@example.com"
  description = "Business group description"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  users = [data.ochk_user.user1.id]
  networks = ["bd814070-18f3-4182-b2af-edaa72a50fee"]

  lifecycle {
    ignore_changes = [description]
  }
}