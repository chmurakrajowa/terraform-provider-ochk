package ochk

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/strfmt"
	"io"
	"os"
	"reflect"
	"strings"
)

var testData = getTestData()
var testDataSavedFileName = "../env/predefined-resources.json"

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
				break
			}
			val := reflect.ValueOf(&predefinedTestDataDev).Elem().FieldByName(m.Name)
			if val.IsValid() {
				reflect.ValueOf(&predefinedTestDataDev).Elem().FieldByName(m.Name).SetString(m.Text)
			}
		}
	}
}

func getTestData() predefinedTestData {
	loadTestData()
	return predefinedTestDataDev
}

type predefinedTestData struct {
	LogicalPort1DisplayName    string
	Network1Name               string
	Network2Name               string
	Project1Name               string
	Project1Id                 strfmt.UUID
	Project2Name               string
	Project3Name               string
	Project4Name               string
	ProjectForVMName           string
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
	TagName                    string
	VPC                        string
	VRF                        string
	VRF_OPENSTACK              string
	AutoNatName                string
	DnatName                   string
	InfraAdminGroup            string
	FirewallEWRuleName         string
	FirewallSNRuleName         string
	FirewallRuleName           string
	NatPublicIpAddr            string
	SnapshotName               string
	AccountName                string
	SecurityGroupName          string
	FloatingIpAddressName      string
	FloatingIpAddress          string
	PortForwardingName         string
}

var devTestDataPrefix = "tf-gojl"
var predefinedTestDataDev = predefinedTestData{
	LogicalPort1DisplayName:    "",
	BackupPlanName:             "",
	BackupListName:             "",
	Network1Name:               fmt.Sprintf("%s-vnet1", devTestDataPrefix),
	Network2Name:               fmt.Sprintf("%s-vnet2", devTestDataPrefix),
	Project1Name:               fmt.Sprintf("%s-project-01", devTestDataPrefix),
	Project1Id:                 strfmt.UUID("3cda830c-b37f-46dc-be54-f649d31bec66"),
	Project2Name:               fmt.Sprintf("%s-project-02", devTestDataPrefix),
	Project3Name:               fmt.Sprintf("%s-project-03", devTestDataPrefix),
	Project4Name:               fmt.Sprintf("%s-project-04", devTestDataPrefix),
	ProjectForVMName:           fmt.Sprintf("%s-project-01", devTestDataPrefix),
	VirtualMachine1DisplayName: fmt.Sprintf("%s-vm-default", devTestDataPrefix),
	VirtualNetwork1DisplayName: fmt.Sprintf("%s-vnet3", devTestDataPrefix),
	VirtualNetwork2DisplayName: fmt.Sprintf("%s-vnet4", devTestDataPrefix),
	VirtualMachineDisplayName:  fmt.Sprintf("%s-vm1", devTestDataPrefix),
	VirtualMachine2DisplayName: fmt.Sprintf("%s-vm2", devTestDataPrefix),
	IPCollection1DisplayName:   fmt.Sprintf("%s-ipc-default", devTestDataPrefix),
	Deployment1DisplayName:     "",
	CustomService1DisplayName:  fmt.Sprintf("%s-https", devTestDataPrefix),
	CustomService2DisplayName:  fmt.Sprintf("%s-http", devTestDataPrefix),
	TagName:                    fmt.Sprintf("%s-t1", devTestDataPrefix),
	KMSKeyDisplayName:          fmt.Sprintf("%s-key", devTestDataPrefix),
	VPC:                        fmt.Sprintf("%s-vpc", devTestDataPrefix),
	AutoNatName:                fmt.Sprintf("%s-autonat", devTestDataPrefix),
	DnatName:                   fmt.Sprintf("%s-dnat", devTestDataPrefix),
	VRF:                        "",
	VRF_OPENSTACK:              "",
	InfraAdminGroup:            fmt.Sprintf("%s-subt1-InfraAdm", devTestDataPrefix),
	FirewallEWRuleName:         fmt.Sprintf("%s-tf-fw-ew-http", devTestDataPrefix),
	FirewallSNRuleName:         fmt.Sprintf("%s-tf-fw-sn-http", devTestDataPrefix),
	FirewallRuleName:           fmt.Sprintf("%s-tf-fw", devTestDataPrefix),
	NatPublicIpAddr:            "",
	SnapshotName:               fmt.Sprintf("%s-snaps001", devTestDataPrefix),
	AccountName:                fmt.Sprintf("%s-act1", devTestDataPrefix),
	SecurityGroupName:          fmt.Sprintf("%s-sg1", devTestDataPrefix),
	FloatingIpAddressName:      "jltest01",
	FloatingIpAddress:          "203.0.113.15",
	PortForwardingName:         "pbtest1", // fmt.Sprintf("%s-port-fwd", devTestDataPrefix),
}
