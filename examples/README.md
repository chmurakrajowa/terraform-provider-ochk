## How to use examples? ##

Examples of TF code in this directory are using official OCHK terraform provider releases coming from ```https://registry.terraform.io/providers/chmurakrajowa/ochk/latest``` .

1. Enter ```examples``` directory in this code branch.

2. Create yourself a dev.tfvars file with following contents:

```
host = "" #OCHK API endpoint
tenant = "" # obtain this account name from OCHK support
username = "" # OCHK API username
password = "" # OCHK API password
debug_log_file = "debug.log"
subtenant_name = "" # This is your new project name you create using foundation code

seed_network_name = "initial" #first subnetwork name for an account
seed_subtenant_name = "admin" #first subaccount name for an account
seed_ochk_user = "terra-tenantadm" #first seed user name

```

3. Invoke ```terraform init -var-file=dev.tfvars``` to download all neede TF modules and init the .terraform directory.

4. Do normal Terraformy things:

```
terraform plan -var-file=dev.tfvars
terraform apply -var-file=dev.tfvars
terraform destroy -var-file=dev.tfvars
```

### Setting up your OCHK accounts/projects

1. Use code from ```examples/foundation``` to create a subtenant/project.
2. Take subtenant/project name generated in step 1 and use the code from ```examples/vm-network-creation``` to create networks and VMs. 

Enjoy!
