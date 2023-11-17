/*data "ochk_projects" "projects" {

}*/
/*
data "ochk_project" "project_for_vm" {
  //name = var.project_for_vm_name
  display_name = "admin"
}

resource "ochk_project" "proj-1" {
  display_name = "${var.test-data-prefix}-bg-01"
  vrf_id = data.ochk_vrf.project-vrf.id
  #display_name = "testy-test"
  description = "qq"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  vcpu_reserved_quantity = 10
}

resource "ochk_project" "project1" {
  display_name = "${var.test-data-prefix}-proj01"
  vrf_id = data.ochk_vrf.project-vrf.id
  description = "qq"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400
  vcpu_reserved_quantity = 10
}
*/
/*
resource "ochk_project" "project-1" {
  display_name = "${var.test-data-prefix}-project-1"
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
resource "ochk_project" "project-2" {
  name = "${var.test-data-prefix}-project-2"
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
resource "ochk_project" "project-2" {
  name = "${var.test-data-prefix}-project-2"
  description = "Business group description 2"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400

  lifecycle {
    ignore_changes = [
      description
    ]
  }
}

resource "ochk_project" "project-3" {
  name = "${var.test-data-prefix}-project-3"
  description = "Business group description 3"
  memory_reserved_size_mb = 30000
  storage_reserved_size_gb = 400

  lifecycle {
    ignore_changes = [
      description
    ]
  }
}

resource "ochk_project" "project-4" {
  name = "${var.test-data-prefix}-project-4"
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
