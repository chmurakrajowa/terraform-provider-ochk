package ochk

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"
)

var testData = getTestData()
var testDataSavedFileName = "../env/testdata.json"
var testDataSavedFields []string

func loadTestData() {

	jsonfile, err := os.ReadFile(testDataSavedFileName)
	if err == nil {
		type Message struct {
			Name, Text string
		}

		var m Message

		decoder := json.NewDecoder(strings.NewReader(string(jsonfile)))
		for {
			if err := decoder.Decode(&m); err == io.EOF {
				break
			} else if err != nil {
				fmt.Errorf(err.Error())
				break
			}
			val := reflect.ValueOf(&predefinedTestDataDev).Elem().FieldByName(m.Name)
			if val.IsValid() {
				reflect.ValueOf(&predefinedTestDataDev).Elem().FieldByName(m.Name).SetString(m.Text)
			} else {
				fmt.Errorf("TestData elemenet: " + m.Name + "not exists")
			}
		}
	}
}

func getTestData() predefinedTestData {
	loadTestData()
	return predefinedTestDataDev
}

type predefinedTestData struct {
	mu                         sync.Mutex
	LogicalPort1DisplayName    string
	Network1Name               string
	Network2Name               string
	Subtenant1Name             string
	Subtenant2Name             string
	Subtenant3Name             string
	Subtenant4Name             string
	SubtenantForVMName         string
	User1Name                  string
	User1Email                 string
	User2Name                  string
	User2Email                 string
	VirtualMachine1DisplayName string
	VirtualNetwork1DisplayName string
	VirtualNetwork2DisplayName string
	VirtualMachineDisplayName  string
	VirtualMachine2DisplayName string
	IPCollection1DisplayName   string
	Deployment1DisplayName     string
	CustomService1DisplayName  string
	CustomService2DisplayName  string
	KMSKeyDisplayName          string
	BackupPlanName             string
	BackupListName             string
	BillingTagName             string
	SystemTagName              string
	VPC                        string
	VRF                        string
	AutoNatName                string
	DnatName                   string
	InfraAdminGroup            string
	FirewallEWRuleName         string
	FirewallSNRuleName         string
}

var devTestDataPrefix = "tf-tst"
var predefinedTestDataDev = predefinedTestData{
	LogicalPort1DisplayName:    "",
	BackupPlanName:             "",
	BackupListName:             "",
	Network1Name:               fmt.Sprintf("%s-vnet1", devTestDataPrefix),
	Network2Name:               fmt.Sprintf("%s-vnet2", devTestDataPrefix),
	Subtenant1Name:             fmt.Sprintf("%s-subt1", devTestDataPrefix),
	Subtenant2Name:             fmt.Sprintf("%s-subt2", devTestDataPrefix),
	Subtenant3Name:             fmt.Sprintf("%s-subt3", devTestDataPrefix),
	Subtenant4Name:             fmt.Sprintf("%s-subt4", devTestDataPrefix),
	SubtenantForVMName:         fmt.Sprintf("%s-subt1", devTestDataPrefix),
	User1Name:                  fmt.Sprintf("%s-user1", devTestDataPrefix),
	User1Email:                 fmt.Sprintf("%s-user1@ochk.pl", devTestDataPrefix),
	User2Name:                  fmt.Sprintf("%s-user2", devTestDataPrefix),
	User2Email:                 fmt.Sprintf("%s-user2@ochk.pl", devTestDataPrefix),
	VirtualMachine1DisplayName: fmt.Sprintf("%s-vm", devTestDataPrefix),
	VirtualNetwork1DisplayName: fmt.Sprintf("%s-vnet3", devTestDataPrefix),
	VirtualNetwork2DisplayName: fmt.Sprintf("%s-vnet4", devTestDataPrefix),
	VirtualMachineDisplayName:  fmt.Sprintf("%s-vm1", devTestDataPrefix),
	VirtualMachine2DisplayName: fmt.Sprintf("%s-vm2", devTestDataPrefix),
	IPCollection1DisplayName:   fmt.Sprintf("%s-ipc-default", devTestDataPrefix),
	Deployment1DisplayName:     "",
	CustomService1DisplayName:  fmt.Sprintf("%s-https", devTestDataPrefix),
	CustomService2DisplayName:  fmt.Sprintf("%s-http", devTestDataPrefix),
	BillingTagName:             fmt.Sprintf("%s-billing-t1", devTestDataPrefix),
	SystemTagName:              fmt.Sprintf("%s-system-t1", devTestDataPrefix),
	KMSKeyDisplayName:          fmt.Sprintf("%s-key", devTestDataPrefix),
	VPC:                        fmt.Sprintf("%s-router", devTestDataPrefix),
	AutoNatName:                fmt.Sprintf("%s-autonat", devTestDataPrefix),
	DnatName:                   fmt.Sprintf("%s-dnat", devTestDataPrefix),
	VRF:                        "",
	InfraAdminGroup:            fmt.Sprintf("%s-subt1-InfraAdm", devTestDataPrefix),
	FirewallEWRuleName:         fmt.Sprintf("%s-tf-fw-ew-http", devTestDataPrefix),
	FirewallSNRuleName:         fmt.Sprintf("%s-tf-fw-sn-http", devTestDataPrefix),
}
