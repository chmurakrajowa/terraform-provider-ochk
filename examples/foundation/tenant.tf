data "ochk_subtenant" "seed_subtenant" {
  name = var.seed_subtenant_name
}

data "ochk_user" "seed_user" {
  name = var.seed_ochk_user
}

resource "random_string" "random" {
  length = 5
  special = false
}

resource "ochk_subtenant" "my_subtenant" {
  name = "${var.tenant}-subtenant-${random_string.random}"
  email = "email@example.com"
  description = "Subtenant desciptions"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  users = [
    data.ochk_user.seed_user.id
  ]
  networks = [
    data.ochk_network.subtenant_seed_network.id
  ]

  lifecycle {
    ignore_changes = [
      description
    ]
  }
}
