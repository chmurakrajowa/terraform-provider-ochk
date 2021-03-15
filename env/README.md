This terraform source is used for dogfooding predefined test resources on different environments.

After creating resource here, add it to terraform outputs and define in `testdata_predefined_test.go` file.
 
For running on specific env provide `var-<env>.tfvars` file with following variables:
```
host = "<host>"
tenant = "<tenant>"
username = "<username>"
password = "<password>"
debug_log_file = "<debug log file>"

# following variables are for resources which are unable to create using terraform yet
test-data-prefix = "tf-test"
test_user = "devel-tftest"
subtenant_network_name = "vtest7"
subtenant_for_vm_name = "auto-bg1"
```