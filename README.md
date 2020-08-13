OChK Terraform Provider
==================

![Build](https://github.com/chmurakrajowa/ochk-terraform-provider/workflows/Terraform%20Provider%20Checks/badge.svg?event=push)

This repository contains Terraform Provider for managing cloud resources in OChK. 

- Website: https://chmurakrajowa.pl

<img src="https://chmurakrajowa.pl/img/logotypes/chmura-krajowa-logo.svg" width="200px">

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.12+
- [Go](https://golang.org/doc/install) 1.14+ (to build the provider plugin)

Building The Provider
---------------------

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
----------------------

Provide connection details in `terraform.tfvars` file. Password should be set using `OCHK_PASSWORD` environment variable.

```sh
$ cd examples
$ export OCHK_PASSWORD=*******
$ terraform init
$ terraform apply
```

Developing the Provider
---------------------------

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
$ export OCHK_HOST=host
$ export OCHK_TENANT=tenant
$ export OCHK_USERNAME=username
$ export OCHK_PASSWORD=*******

$ go test ./...
```