/*
data "ochk_subtenant" "subtenant_for_vm" {
  name = var.subtenant_for_vm_name
}
*/

/*
data "ochk_user" "bg_manager_user1" {
  name = var.bg_manager_user
}

data "ochk_user" "bg_manager_user2" {
  name = var.bg_manager_user2
}
*/

/*
resource "ochk_subtenant" "subtenant-1" {
  name = "${var.test-data-prefix}-bg-01"
  #name = "testy-test"
  email = "email2@example.com"
  description = "qq"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
}
*/

/*
resource "ochk_subtenant" "subtenant-1" {
  name = "${var.test-data-prefix}-subtenant-1"
  email = "email1@example.com"
  description = "Business group description 1"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400

  lifecycle {
    ignore_changes = [
      description
    ]
  }
}
*/

/*
resource "ochk_subtenant" "subtenant-2" {
  name = "${var.test-data-prefix}-subtenant-2"
  email = "email1@example.com"
  description = "Business group description 2"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400

  lifecycle {
    ignore_changes = [
      description
    ]
  }
}
*/
/*
resource "ochk_subtenant" "subtenant-2" {
  name = "${var.test-data-prefix}-subtenant-2"
  email = "email2@example.com"
  description = "Business group description 2"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400

  lifecycle {
    ignore_changes = [
      description
    ]
  }
}

resource "ochk_subtenant" "subtenant-3" {
  name = "${var.test-data-prefix}-subtenant-3"
  email = "email3@example.com"
  description = "Business group description 3"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400

  lifecycle {
    ignore_changes = [
      description
    ]
  }
}

resource "ochk_subtenant" "subtenant-4" {
  name = "${var.test-data-prefix}-subtenant-4"
  email = "email4@example.com"
  description = "Business group description 4"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400

  lifecycle {
    ignore_changes = [
      description
    ]
  }
}
*/
