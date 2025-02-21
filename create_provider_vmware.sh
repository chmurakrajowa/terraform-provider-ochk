#!/usr/bin/env bash

rm -r examples/VMWARE/.terraform
rm -r examples/VMWARE/terraform.d
rm examples/VMWARE/.terraform.lock.hcl

rm -r debug.log

make build_local
case $1 in
  "env" )
#    . test.sh
    cd env

    ./create_testdata.sh

    #copy development version of provider to env directory
    rm -r debug.log
    rm -r .terraform
    rm -r terraform.d
    rm .terraform.lock.hcl
    cp -r ../examples/terraform.d ./

    terraform init -var-file=var-test.tfvars -backend-config="path=test.tfstate"
    TF_DEBUG=debug terraform apply -var-file=var-test.tfvars

    cd ..
    export DEBUG=1
    go test -v ./...
    ;;

  *)
    cd examples/VMWARE
    rm -r debug.log
    terraform init -var-file=test.tfvars -backend-config="path=test.tfstate"
    TF_DEBUG=DEBUG TF_LOG=DEBUG terraform apply -var-file=test.tfvars  -no-color 2>&1 | tee plan.log
    ;;
esac
