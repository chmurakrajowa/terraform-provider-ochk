## Compiling ##

Just use any OS with Golang installed and do ```make build_local```. Then follow the setup for launching examples below.

## How to use examples? ##

1. Enter ```examples``` directory in this code branch.

2. Create yourself a dev.tfvars file with following contents:

```
host = <your OCHK iaas API endpoint address>
tenant = <tenant name>
username = <your OCHK IAAS API username>
password = <your OCHK IAAS API password>
debug_log_file = "debug.log"
test-data-prefix = <your custom prefix for test data names>
test_user = <test user name>
subtenant_network_name = <subtenant network name for your vms>
subtenant_for_vm_name = <subtenant name for your vms>
```

3. Invoke ```terraform init -var-file=dev.tfvars``` to download all neede TF modules and init the .terraform directory.

4. Do normal Terraformy things:

```
terraform plan -var-file=dev.tfvars
terraform apply -var-file=dev.tfvars
terraform destroy -var-file=dev.tfvars
```

Enjoy!
