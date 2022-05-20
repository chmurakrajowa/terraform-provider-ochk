#!/bin/bash
DEPLOYMENT_NAME="CentOS 7"


TEST_DATA_FILE="testdata.json"
rm ./$TEST_DATA_FILE
rm ./"$0.log"

get_token() {
  echo "Reading Bearer Token"
  echo "--- Reading Bearer Token" >> "$0.log"
  BODY="{\"username\":\"$USERNAME\",\"password\":\"$PASSWORD\",\"tenant\":\"$TENANT\"}"
  TOKEN_RES=$(curl --silent -X POST --header 'Content-Type: application/json' --data $BODY "https://$HOST/vidm/token" --insecure)
  TOKEN=$(echo $TOKEN_RES| jq -r '.token')
  if [ -z "$TOKEN" ] || [ "$TOKEN" = "null" ]; then
    echo "Wrong login. Token is empty. Check configuration file for properly login parameters and network connection"
    exit 1
  fi

  echo "curl POST https://$HOST/$ENDPOINT" >> "$0.log"
  echo "BODY: $BODY" >> "$0.log"
  echo "RESPONSE: $TOKEN_RES" >> "$0.log"
  echo "" >> "$0.log"
}

get_logical_ports() {
  BACKUP_PLAN_ID=$1
  HEADER="token:\"$TOKEN\""
  ENDPOINT="network/logical-ports"

  local LIST=$(curl --silent -X GET --header $HEADER https://$HOST/$ENDPOINT --insecure)
  echo $LIST |jq  -r '.logicalPortCollection[] | "\(.id)|\(.displayName)" '

  echo "curl GET https://$HOST/$ENDPOINT" >> "$0.log"
  echo "HEAD: $HEADER" >> "$0.log"
  echo "RESPONSE: $LIST" >> "$0.log"
  echo "" >> "$0.log"
}

get_nat_public_ip_addr() {
  HEADER="token:\"$TOKEN\""
  ENDPOINT="ipam/ipaddress/public/allocation"
  local LIST=$(curl --silent -X GET --header $HEADER https://$HOST/$ENDPOINT --insecure)
  #echo $LIST |jq -r  '.publicIpAllocationCollection[] | "\(.publicIpAddress.ipAddress)|\(.serviceList[].name)" ' | grep "NAT"
  echo $LIST |jq -r  '.publicIpAllocationCollection[] | "\(.publicIpAddress.ipAddress)|\(.publicIpAddress.ipAddressId)" '

  echo "curl GET https://$HOST/$ENDPOINT" >> "$0.log"
  echo "HEAD: $HEADER" >> "$0.log"
  echo "RESPONSE: $LIST" >> "$0.log"
  echo "" >> "$0.log"
}

get_deployment() {
  HEADER="token:\"$TOKEN\""
  ENDPOINT="deployments"
  local LIST=$(curl --silent -X GET --header $HEADER https://$HOST/$ENDPOINT --insecure)
  echo $LIST |jq -r '.deploymentInstanceCollection[] | "\(.deploymentId)|\(.displayName)" ' | grep "|$DEPLOYMENT_NAME"

  echo "curl GET https://$HOST/$ENDPOINT" >> "$0.log"
  echo "HEAD: $HEADER" >> "$0.log"
  echo "RESPONSE: $LIST" >> "$0.log"
  echo "" >> "$0.log"
}

get_vrf() {
  HEADER="token:\"$TOKEN\""
  ENDPOINT="network/routers"
  local LIST=$(curl --silent -X GET --header $HEADER https://$HOST/$ENDPOINT --insecure)
  echo $LIST |jq -r  '.routerCollection[] | select(.routerType=="TIER0") | "\(.routerId)|\(.displayName)" '

  echo "curl GET https://$HOST/$ENDPOINT" >> "$0.log"
  echo "HEAD: $HEADER" >> "$0.log"
  echo "RESPONSE: $LIST" >> "$0.log"
  echo "" >> "$0.log"
}

get_backup_list() {
  BACKUP_PLAN_ID=$1
  HEADER="token:\"$TOKEN\""
  ENDPOINT="backups/plans/$BACKUP_PLAN_ID/lists"

  local LIST=$(curl --silent -X GET --header $HEADER https://$HOST/$ENDPOINT --insecure)
  echo $LIST |jq  -r '.backupListCollection[] | "\(.backupListId)|\(.backupListName)" '

  echo "curl GET https://$HOST/$ENDPOINT" >> "$0.log"
  echo "HEAD: $HEADER" >> "$0.log"
  echo "RESPONSE: $LIST" >> "$0.log"
  echo "" >> "$0.log"
}

get_backup_plans() {
  HEADER="token:\"$TOKEN\""
  ENDPOINT="backups/plans"

  local LIST=$(curl --silent -X GET --header $HEADER https://$HOST/$ENDPOINT --insecure)
  echo $LIST |jq  -r '.backupPlanCollection[] | "\(.backupPlanId)|\(.backupPlanName)" '

  echo "curl GET https://$HOST/$ENDPOINT" >> "$0.log"
  echo "HEAD: $HEADER" >> "$0.log"
  echo "RESPONSE: $LIST" >> "$0.log"
  echo "" >> "$0.log"
}

add_backup_plan() {
 get_backup_plans | while read BackupPlan ; do
    BACKUP_PLAN_ID=$(echo `expr "$BackupPlan" : '\(.*\)|'`)
    NAME=$(echo `expr "$BackupPlan" : '.*|\(.*\)'`)
    echo "=== Found BackupPlan: $NAME ($BACKUP_PLAN_ID)"
    echo "{
	\"Name\": \"BackupPlanName\",
	\"Text\": \"$NAME\"
}" >>$TEST_DATA_FILE
    get_backup_list $BACKUP_PLAN_ID | while read BackupList ; do
      BACKUP_LIST_ID=$(echo `expr "$BackupList" : '\(.*\)|'`)
      NAME=$(echo `expr "$BackupList" : '.*|\(.*\)'`)
      echo "=== Found BackupList: $NAME ($BACKUP_LIST_ID)"
      echo "{
    \"Name\": \"BackupListName\",
    \"Text\": \"$NAME\"
}" >>$TEST_DATA_FILE
      break
    done
   break
  done
}

add_logical_port() {
 get_logical_ports | while read LogicalPort ; do
    UUID=$(echo `expr "$LogicalPort" : '\(.*\)|'`)
    NAME=$(echo `expr "$LogicalPort" : '.*|\(.*\)'`)
    echo "=== Found LogicalPort: $NAME ($UUID)"
    echo "{
	\"Name\": \"LogicalPort1DisplayName\",
	\"Text\": \"$NAME\"
}" >>$TEST_DATA_FILE
    break
  done
}

add_vrf() {
 get_vrf | while read LogicalPort ; do
    UUID=$(echo `expr "$LogicalPort" : '\(.*\)|'`)
    NAME=$(echo `expr "$LogicalPort" : '.*|\(.*\)'`)
    echo "=== Found VRF: $NAME ($UUID)"
    echo "{
	\"Name\": \"VRF\",
	\"Text\": \"$NAME\"
}" >>$TEST_DATA_FILE
    break
  done
}

add_deployment() {
 get_deployment | while read Deployment ; do
    UUID=$(echo `expr "$Deployment" : '\(.*\)|'`)
    NAME=$(echo `expr "$Deployment" : '.*|\(.*\)'`)
    echo "=== Found Deployment: $NAME ($UUID)"
    echo "{
	\"Name\": \"Deployment1DisplayName\",
	\"Text\": \"$NAME\"
}" >>$TEST_DATA_FILE
    break
  done
}

add_nat_public_ip_addr() {
  get_nat_public_ip_addr | while read NatAddr ; do
    ADDR=$(echo `expr "$NatAddr" : '\(.*\)|'`)
    NAME=$(echo `expr "$NatAddr" : '.*|\(.*\)'`)
    echo "=== Found Nat Public IP Addr: $ADDR ($NAME)"
    echo "{
	\"Name\": \"NatPublicIpAddr\",
	\"Text\": \"$ADDR\"
}" >>$TEST_DATA_FILE


    echo "locals {
  nat_public_ip_addr = \"$ADDR\"
}">locals.tf
    break
  done
}

. create_testdata.cfg

get_token

## Nat Public IP Addr
VAL=$(add_nat_public_ip_addr)
if [ "$VAL" = "" ]; then
  echo "!!! Nat Public IP Addr not found"
else
  echo "$VAL"
fi

## Backup plan
VAL=$(add_backup_plan)
if [ "$VAL" = "" ]; then
  echo "!!! Backup plan not found"
else
  echo "$VAL"
fi

## Logical port
VAL=$(add_logical_port)
if [ "$VAL" = "" ]; then
  echo "!!! Logical port not found"
else
  echo "$VAL"
fi

## VRF
VAL=$(add_vrf)
if [ "$VAL" = "" ]; then
  echo "!!! VRF not found"
else
  echo "$VAL"
fi

## Deployment
VAL=$(add_deployment)
if [ "$VAL" = "" ]; then
  echo "!!! Deployment not found"
else
  echo "$VAL"
fi
