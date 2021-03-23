resource "ochk_kms_key" "aes-generated" {
  display_name = "${var.test-data-prefix}-aes-key2"
  key_usage = ["ENCRYPT", "DECRYPT"]
  algorithm = "AES"
  size = 256
}

locals {
  # should not really be stored like that
  rsa_private_key = <<EOT
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA8vXmbi9IAPnkIQTwTijtX3tdNRokMJ7WA2pwMOswEihck5n3
f+tbtBTGwWuat2UiPUgWfDtNiqHeK8CTWZuGFyITxF8n2iiQM9CTcR8AJUkp08Fn
ftNcopnEjYN92wV6ehuPo8k88R9OYA6N5SgxvqFraGCTNE6h+02M+rLeetVVMatw
+ZvCz3I8bHFchI9zR+jE3h02z6DvHG9TcaDG2qBn/azWAVY+6sO/6xQgMN1k+jMo
PMH6cS62mBdcK+xNq9NhIDWlptnNFuUw97n7a7VVvsTGIS/rjYZuO/ym1s9gXr6B
+OoiJmuurxSGer8nYlpYq2VpYlXNAcw1UGMhHQIDAQABAoIBAQCGabYm5S9/osAr
6FCN3SSdu2EwfJri7yzVTPBuj97TXNMCsZ50faAJO6lN3psEtQXBQ311E1XtyWlh
aTPb0ifX6nlnHYGttt04XT8EyTLKbfSe+xOn3YUVS96qr8FUB27f2RmZcj6t4zT3
/XVQ/vCuVx1V7H/j41DH9/pzw7tD1mhL0rCjt9NKLXyrFT4ii63q9iLE2GPqfeKV
qbG9A6XGhRcumvsIkVxBTh5ofLK2bxLXdHyRk5HEDq0SmKXEeXPULR5vnfZ/lskP
HNx2rA3yXS8sEBeigCYtQOngByCWSfXy2R5HVKp2t27PWgWYR9gx58T8WOGOhooC
xgWNZmMdAoGBAPsU3VKRzPE2BCrO+hrpYofmZIniHTJGMhVG4mo9WvT8K4eta97Q
qE+4I2YrumdFlQ3/ZttRR0KPV0PkmxloOe18oGsp2qQEmUBeVvguFrkwJ7faXacS
X/WXLyI+mX6p1GPWx62MalYW9bHRn5Fh1TrTVMDXs84zpcxp90HT7PpLAoGBAPe4
T2sPienL3bwl9srRfqo/moQR4xsW7lMIzVgYOZIqhcDf0zPRChW6+uYqn36spDnQ
syZq2LGhCQqhputEbt29jrBsXYTo23FgIojxrrFUcsSnWqWuHhUswAqfSEBVodCr
3g2I4jXu11ZIimweFD/nrvHiqdVA+Vv4dRXrEzE3AoGAFaVc10t+kaUIgvBJG5zX
Q8QXEtQNlFH861yrFGGpv8klr5LB4/m1KPpFAv/uGA0lpolIQswlCpX0/gNtY6la
pSDDa7m2AcHrvGLluIuwpdKC+hS3UjoBT9jy1U70SLk3eEwB9vJOEJH2KJhb21rF
2UZy3hU6iSJmvtK74E729TUCgYEAtDgr1yjL3gLKQ3qfLHjzHOr+//x/bBLnuhMa
SW/+Wl+DRYnQ/s6i9qI8rLzvoln5dHJoE5gCJGCS0mA+rsTvn3Sr3aBI/UvncnlN
RrIFtM9KW9WhNg4RprgS0ueEygFCoyyWdORUJoantQc7ZWMQullUxndvtUz63TVK
AXMvWEsCgYEAhQ6Ylqpuco/60ybjRzUNhGGsyVMZsi03uyYfjzOsu7UIcdurOZTO
A2UhIbAMpBvNLgqqtBwVcwxJkdMW7Z7TGnfqiyh11Pr8yZQcbfMP3FMLgnsTmnrC
0yG2J8e9jUhdd70nTWsAQikQaaszrZLdJ/Hk1zm0VIU+ohBaWmg02m4=
-----END RSA PRIVATE KEY-----
EOT
}

resource "ochk_kms_key" "rsa-to-unwrap" {
  display_name = "rsa-to-unwrap2"
  key_usage = ["ENCRYPT", "DECRYPT"]
  algorithm = "RSA"
  size = 2048
  material = local.rsa_private_key
}

resource "ochk_kms_key" "aes-enc" {
  display_name = "${var.test-data-prefix}-aes-enc2"
  key_usage = ["ENCRYPT", "DECRYPT"]
  algorithm = "AES"
  size = 256
  material = "IgQEzbAU/oCTkXyrRZ6zKEna3aNQv+G9OBdwwkDmT3rLE9TbP4BeVlBm7drqt7yc/xUS/QJeHcGxjpVH4BNbLN1iJ91DTdV7ombHVOra3wpXX03zxPgZvBEBnoWzSvRtZjOr+XtfUSILXE9lHCYVnPCDjxi+bZw7Ls095fd/V1QoIG3RShKHtzFr8am7+TG9oTh2NjIYfTRy36jng1WTYAIM10LGf2SvgN7h7XQQrZxlEqwO5bB1psCTNc9WliLCL7jJK6XSzc+AfYEe231vbQ+KDWBAuq5SgsgQnhA7B7xcyIUpaSm0g8Ie5MM9+XGvUxRfRKxAXdhrXbX/yA6JIw=="
  private_key_id_to_unwrap = ochk_kms_key.rsa-to-unwrap.id
}