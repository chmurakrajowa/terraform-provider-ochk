OChK Terraform Provider
==================

![Test](https://github.com/chmurakrajowa/terraform-provider-ochk/actions/workflows/test.yml/badge.svg?event=push)

This repository contains Terraform Provider for managing cloud resources in OChK. 

- Documentation: https://registry.terraform.io/providers/chmurakrajowa/ochk/latest/docs
- Website: https://ochk.cloud

<img src="https://ochk.cloud/images/ochk_extended_logo.svg" width="200px">

Requirements
-------------------------

- [Terraform](https://www.terraform.io/downloads.html) 1.3.7+
- [Go](https://golang.org/doc/install) 1.18+ (to build the provider plugin)

Building The Provider
-------------------------

Clone the repository.

```sh
$ git clone git@github.com:chmurakrajowa/ochk-terraform-provider.git
```

Enter the provider directory and build the provider. Output binary will be placed examples directory.

```sh
$ cd ochk-terraform-provider
$ go build -o examples/ ./...
```

Using the provider
--------------------------

Provide connection details in `terraform.tfvars` file. 

```sh
$ cd examples
$ terraform init
$ terraform apply
```

Developing the Provider
---------------------------

---------------------------
Update/generate API client for Platforma OChK V3.

1. To generate new client from swagger, you need to download new swagger definition file in json format from https://pckproxy-at.ochk.pilot if you are working on AUTOTEST enviroment.

Swagger.json file is available at: https://pckproxy-at.ochk.pilot/swagger/v1/swagger.json

2. Place downloaded file in location ./api/v3/swagger.json
3. Run script:
```sh
./tools/swagger_gen.sh
```
4. New proxy client will be generated. Do not modify any file under /api/v3/ location

--------------------
If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.13+ is *required*).

To compile the provider, run `go build ./...`. This will build the provider and put the provider binary in the current directory.

```sh
$ go build ./...
```

In order to unit test the provider, you can simply run `go test ./...`.

```sh
$ go test ./...
```

In order to run the full suite of Acceptance tests you need to provide connection details in environment variables. Also enable running Acceptance tests by setting `TF_ACC=true`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ export TF_ACC=true
$ export TF_VAR_host=host
$ export TF_VAR_platform=platform
$ export TF_VAR_api_key=api_key
$ export TF_VAR_platform_type=VMWARE / OPENSTACK

$ go test ./...
```
