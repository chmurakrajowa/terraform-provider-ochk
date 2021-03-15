package ochk

import (
	"fmt"
	"os"
)

var testData = getTestData()

func getTestData() predefinedTestData {
	switch os.Getenv("TEST_ENV") {
	case "AT":
		return predefinedTestDataAT
	default:
		return predefinedTestDataDev
	}
}

type predefinedTestData struct {
	LogicalPort1DisplayName          string
	Network1Name                     string
	Network2Name                     string
	SubtenantNetworkName             string
	Subtenant1Name                   string
	Subtenant2Name                   string
	Subtenant3Name                   string
	Subtenant4Name                   string
	SubtenantForVMName               string
	User1Name                        string
	User2Name                        string
	VirtualMachine1DisplayName       string
	VirtualNetwork1DisplayName       string
	VirtualNetwork2DisplayName       string
	LegacyVirtualMachineDisplayName  string
	LegacyVirtualMachine2DisplayName string
	IPCollection1DisplayName         string
	Deployment1DisplayName           string
	CustomService1DisplayName        string
	CustomService2DisplayName        string
	KMSKeyDisplayName                string
	SecurityPolicyDisplayName        string
}

var devTestDataPrefix = "tf-test2"
var predefinedTestDataDev = predefinedTestData{
	LogicalPort1DisplayName:          "d0ac165f-cec0-db4a-2a82-b8599f064900/devel0000001256.vmx@93f28b5d-ec29-4ba2-b753-160ce290b7fd",
	Network1Name:                     "vtest8",
	Network2Name:                     "vtest7",
	SubtenantNetworkName:             "vtest7",
	Subtenant1Name:                   fmt.Sprintf("%s-subtenant-1", devTestDataPrefix),
	Subtenant2Name:                   fmt.Sprintf("%s-subtenant-2", devTestDataPrefix),
	Subtenant3Name:                   fmt.Sprintf("%s-subtenant-3", devTestDataPrefix),
	Subtenant4Name:                   fmt.Sprintf("%s-subtenant-4", devTestDataPrefix),
	SubtenantForVMName:               "auto-bg1",
	User1Name:                        "devel-tftest",
	User2Name:                        "devel-bdgn",
	VirtualMachine1DisplayName:       fmt.Sprintf("%s-vm", devTestDataPrefix),
	VirtualNetwork1DisplayName:       fmt.Sprintf("%s-vnet3", devTestDataPrefix),
	VirtualNetwork2DisplayName:       fmt.Sprintf("%s-vnet4", devTestDataPrefix),
	LegacyVirtualMachineDisplayName:  "devel0000000098",
	LegacyVirtualMachine2DisplayName: "devel0000000344",
	IPCollection1DisplayName:         fmt.Sprintf("%s-ipc-default", devTestDataPrefix),
	Deployment1DisplayName:           "CentOS 7",
	CustomService1DisplayName:        fmt.Sprintf("%s-https", devTestDataPrefix),
	CustomService2DisplayName:        fmt.Sprintf("%s-http", devTestDataPrefix),
}

var testTestDataPrefix = "tf-test1"
var predefinedTestDataAT = predefinedTestData{
	LogicalPort1DisplayName:          "196eae5f-824c-3156-aa0b-98039b5c9998/Vm12345678-.vmx@8325dc52-dddc-4b43-89a1-a20271a40928",
	Network1Name:                     "CMP_VNET_7",
	Network2Name:                     "CMP_VNET_71",
	SubtenantNetworkName:             "CMP_VNET_1",
	Subtenant1Name:                   fmt.Sprintf("%s-subt1", testTestDataPrefix),
	Subtenant2Name:                   fmt.Sprintf("%s-subt2", testTestDataPrefix),
	Subtenant3Name:                   fmt.Sprintf("%s-subt3", testTestDataPrefix),
	Subtenant4Name:                   fmt.Sprintf("%s-subt4", testTestDataPrefix),
	SubtenantForVMName:               "bg_001",
	User1Name:                        "testy-tftest",
	User2Name:                        "testy-tt",
	VirtualMachine1DisplayName:       fmt.Sprintf("%s-vm", testTestDataPrefix),
	VirtualNetwork1DisplayName:       fmt.Sprintf("%s-vnet3", testTestDataPrefix),
	VirtualNetwork2DisplayName:       fmt.Sprintf("%s-vnet4", testTestDataPrefix),
	LegacyVirtualMachineDisplayName:  "testy0000000133",
	LegacyVirtualMachine2DisplayName: "testy0000000083",
	IPCollection1DisplayName:         fmt.Sprintf("%s-ipc", testTestDataPrefix),
	Deployment1DisplayName:           "CentOS 7",
	CustomService1DisplayName:        fmt.Sprintf("%s-https", testTestDataPrefix),
	CustomService2DisplayName:        fmt.Sprintf("%s-http", testTestDataPrefix),
	SecurityPolicyDisplayName:        "testy",
	KMSKeyDisplayName:                fmt.Sprintf("%s-key", testTestDataPrefix),
}
