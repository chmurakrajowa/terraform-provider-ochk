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

var predefinedTestDataAT = predefinedTestData{
	LogicalPort1DisplayName:          "f766455e-22a7-a2de-35f5-b8599f064a08/devel0000000639.vmx@a596ebcb-a875-4188-8ec5-0a84bfbf1e11",
	Network1Name:                     "vtest8",
	Network2Name:                     "vtest7",
	SubtenantNetworkName:             "vtest7",
	Subtenant1Name:                   "tf-acc_test-9j97besn",
	Subtenant2Name:                   "tf-acc_test-0yawdijz",
	Subtenant3Name:                   "tf-test-name-tf-acc_test-amwyi09b",
	Subtenant4Name:                   "tf-test-name-tf-acc_test-03w1wt34",
	SubtenantForVMName:               "admin",
	User1Name:                        "devel-jpuser",
	User2Name:                        "devel-jpuser",
	VirtualMachine1DisplayName:       "devel0000001157",
	VirtualNetwork1DisplayName:       "tf-test-vnet",
	VirtualNetwork2DisplayName:       "tf-test-vnet2",
	LegacyVirtualMachineDisplayName:  "devel0000000098",
	LegacyVirtualMachine2DisplayName: "devel0000000098",
	IPCollection1DisplayName:         "tf-test-ipcollection",
	Deployment1DisplayName:           "CentOS 7",
	CustomService1DisplayName:        fmt.Sprintf("%s-https", devTestDataPrefix),
	CustomService2DisplayName:        fmt.Sprintf("%s-http", devTestDataPrefix),
}
