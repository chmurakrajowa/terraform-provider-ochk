/*
resource "ochk_kms_key" "aes-generated" {
  display_name = "${var.test-data-prefix}-aes-key2"
  key_usage = ["ENCRYPT", "DECRYPT"]
  algorithm = "AES"
  size = 256
}
*/


/*
locals {
  # should not really be stored like that
  rsa_private_key = <<EOT
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEArUvueJtZtHhMtUVEHNI57cSVOOSlebpGl3XyYGW7MJQWJVSbgCppBjBN0CH/nuLwAWa8J8U1wfoy018GrwVj3TQSfhgqjq8KgJBcvh4FTjyNe5qHar04T32OTsQmNx5/Gl1FVcSZbFG7tgafpMkSEycSQbA2jioJzQUECIftP0DxXus8z18wPM6xI/3eeOjt1C7HIIxyZ1RZHxICRKL3+5ssMdw/1SCT21FyrqGrKwqJErc5MFvLIfzS5AsJkFQuA7kUpqO/BrVbKmAVoG5IJ4UQYmZuLq5bfaPfxKxFMg8cfzYxO+SE/7d3LJjZHwDRgPDidirO9z8i9eXS85d0IQIDAQABAoIBAQCi+IWLZo016uw4jKtBb3KK6BRtLJlxoHxCc+k8Pm+Lf7hT4v1Zyoh6CXMkCNUb1s9egf7wZ8XmZiTmIrdmYOtJW28IpjFffb2iqiCNqi3TJZ7oXqp69ve9yZXK/SvzKfreEFjkD7843qRoylHX5j/ZQyUoIU8s4tvUsqLYYvv6PsyIWMAcsIM0r6n0tQvXJUjNki2vnz4EkoAyJ4wBgfO1zxBsztONa/qsBzD8uGPGyhuopyUXuFMkUER0yDaoRXfK6hEyb5GLp/uA+Bo8tzGWv46A4P6WnTkg4xSXNI3g0Lo8PPSxPpBIt9yHORBzCzRQW+SW1nF4nZ7t5fmtKzMBAoGBANYfmhtCEzrA2aufJOjRGLxU6En2ktNoj5Qex5V8i86iSs3krIKl2hwLkfJ6ZyzIBqtaAUfGbOhAZaIR6YMgo4RaK1oFacC0CsH9L40121iBIMbjIsMpi5LGwDvDUT49SscfKtwYwDfGTqHsjuNEnwqD4S06V+oN+SkDRadqvptPAoGBAM8wRnZlwEM4OTbQN1Q+EoUGNTqyj6D5HhIS//+eS9eBcZJsKNNlxgkfwcXWN288z1p0HQe/rQv3ggUz5+yU7Fe3gJYXhTVe6/bYtIpS15hl2vYF2InTTxz+BB3IKbGItd2x9VSt9go0DxdydUJ/Fdm/lgBBKUGgGxQ3uP6HXV2PAoGATs1MktGLcCj/Mskb9U7WsqDw6B+Fw6YMiS+Wti4i6cYSj5mj32UYUw/zA3EGvZVT1wYFyJt/Ay6H+NBsxuRhKVlh8PxjAheAigoa0PZMJlNHc8qtxXNpuXfJ0XS9VpGKDqQqytuVCjHCSJddsshTW7RTT6jUKABCfVSF8uo3AcMCgYEAwCf30etAy5cjiTDHf/wp1PITACOHjjY9q2mvv/WiRXcqTDDAFH/5kZDAeKnas/JkJXI3Um51IfolvomLeiIMqLXqLkoWI4+rrT8DAp8ekZs13qNBV2KBp1P3oV95O8ydT8Wy0Ryeaowpywc+tbOEInXLuuS0TMcfFjsv+fC1qQUCgYA1PsGyy2a1uJpuPLBXbVgEj96Q2Mub6PlGgqIR8IMRI8mPKIAf2V6JYGH0LFdiWrNQScfEGeufYWRAe+W2gKkdGF26mXVYrWnUbrWCTaPE0hfzrs4A5z/YBMVNcThum4FCdJEhO7rItNE9WC1dhzBW3VYky+vHHuUzTTXDOvA+hg==
-----END RSA PRIVATE KEY-----
EOT
}


resource "ochk_kms_key" "rsa-to-unwrap" {
  display_name = "${var.test-data-prefix}-rsa-to-unwrap"
  key_usage = ["ENCRYPT", "DECRYPT"]
  algorithm = "RSA"
  size = 2048
  material = local.rsa_private_key
}

resource "ochk_kms_key" "aes-enc" {
  display_name = "${var.test-data-prefix}-aes-enc"
  key_usage = ["ENCRYPT", "DECRYPT"]
  algorithm = "AES"
  size = 256
  material = "lfdqqLLb+Wetq5IUex60VmISQ/qlElBettd+Cs8K7qqgwA0v7O8JdIw4bKGy1Fl2A033Fw8oHrtk0ZoUD99UQtWvsAiGZUbpvffdjiJDGoReOGANQD12lbeOwv9+911IuTX8Facys4yucJYpiQ75LcoDaj1M4zruVH5nwq3TMls4K1pQVjammmMFcTy30aDnTY13GmCjbSFGoDJ051FZmL9E23XXaI3sjD6oVsl7H5TN4w+13eHQDhFq84OBaffytqya+JihpROlh331rYDkILBM0WCkIV5aYjGPoGhT3HJNoPCLfwZhd7uqsBKhhFqYp9+DBNuMF9MH2DCzpEFjhA=="
  private_key_id_to_unwrap = ochk_kms_key.rsa-to-unwrap.id
}
*/