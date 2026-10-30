#!/usr/bin/env bash


echo "----- REMOVE examples start -----"
rm -r examples/VMWARE/.terraform

rm -r examples/VMWARE/terraform.d
rm examples/VMWARE/.terraform.lock.hcl
echo "----- REMOVE examples end -----"

echo "----- REMOVE debug.log start -----"

rm -r debug.log
echo "----- REMOVE debug.log end -----"


echo "----- make build_local start -----"

make build_local

echo "----- make build_local end -----"


echo "PARAMS: " $1
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
    echo "RUN ONLY BUILD AND RUN PROVIDER"
    cd examples/VMWARE
    rm -r debug.log
#    terraform init -var-file=test.tfvars -backend-config="path=test.tfstate"
#    TF_DEBUG=TRACE TF_LOG=TRACE TF_LOG_PROVIDER=TRACE terraform apply -parallelism=1 -var-file=test.tfvars  -no-color 2>&1 | tee plan.log

    terraform init -var-file=test.tfvars -backend-config="path=test.tfstate"
    TF_DEBUG=TRACE TF_LOG=TRACE TF_LOG_PROVIDER=TRACE terraform apply -var-file=test.tfvars  -no-color 2>&1 | tee plan.log
    ;;
esac
