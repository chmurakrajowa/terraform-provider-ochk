## How to use examples? ##

Examples of TF code in this directory are using official OCHK terraform provider releases coming from ```https://registry.terraform.io/providers/chmurakrajowa/ochk/latest``` .

1. Enter ```examples``` directory in this code branch.

2. Create yourself a dev.tfvars file with following contents:

```
test-data-prefix = <prefix for created objects>
host = <your OCHK iaas API endpoint address>
tenant = <tenant name>
username = <your OCHK IAAS API username>
password = <your OCHK IAAS API password>

bg_manager_user =<test user name>
subtenant_network_name = <vnet name>
subtenant_for_vm_name = <bg name>
vrf_router = "VRF"

iso_image = <iso image name>
ovf_image = <ovf image name>
backup_plan = <backup plan name>
backup_list = "<backup list name>

billing_tag_cc = <billing tag name>
system_tag_os = <system tag name>

debug_log_file = "debug.log"
initial_password_for_vm=<initial password for vm>
```

3. Invoke ```terraform init -var-file=dev.tfvars``` to download all neede TF modules and init the .terraform directory.

4. Do normal Terraformy things:

```
terraform plan -var-file=dev.tfvars
terraform apply -var-file=dev.tfvars
terraform destroy -var-file=dev.tfvars
```

Enjoy!
