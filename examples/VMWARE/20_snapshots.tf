# data "ochk_snapshots" "snapshots" {
#   virtual_machine_id = "5e322d55-4bb9-4860-997f-6f73bc61947c"
# }

/*
data "ochk_snapshot" "snapshot" {
  virtual_machine_id = "bf1a5478-bd0b-448d-9970-c0384331e7d7"
  snapshot_name = "nowa_mig2"
}

*/
/*resource "ochk_snapshot" "snapshot_vm" {
    virtual_machine_id = "5e322d55-4bb9-4860-997f-6f73bc61947c"
    display_name = "${var.test-data-prefix}-proxySnapNoRam043ja"
    ram = true
}*/
/*
resource "ochk_snapshot" "snapshot_vm1" {
    virtual_machine_id = "bf1a5478-bd0b-448d-9970-c0384331e7d7"
    snapshot_name = "${var.test-data-prefix}-proxySnapNoRam000001"
    ram = false
}

resource "ochk_snapshot" "snapshot_vm2" {
    virtual_machine_id = "36b8b6e9-6fd2-4ad6-b26f-a93db34a9178"
    snapshot_name = "${var.test-data-prefix}-proxySnapNoRam04i"
    ram = false
}*/

