provider "ochk" {
  host = "iaas-api-proxy.ochk.pilot"
  tenant = "devel"
  username = "devel-admin"
  password = "zaq1@WSX"
}

resource "ochk_security_group" "demo" {
  display_name = "testM&L"

  members {
    id = "e1e2f617-014c-4119-bac8-49fa4a93db47"
    type = "VIRTUAL_MACHINE"
  }
}