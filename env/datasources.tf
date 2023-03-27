data "ochk_vrf" "default-vrf" {
  display_name = var.vrf_router
}

data "ochk_deployment" "centos" {
  display_name = "CentOS 7"
}

data "ochk_service" "http" {
  display_name = "HTTP"
}
