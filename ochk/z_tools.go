package ochk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/api/v3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io"
	"reflect"
	"strings"
	"text/template"
)

type MemberType string

var pathToFile = "../env/predefined-resources.json"

const (
	IPCOLLECTION MemberType = "IPCOLLECTION"
	LOGICAL_PORT MemberType = "LOGICAL_PORT"
	IPSET        MemberType = "IPSET"
)

func mapResourceDataToSecurityGroup(d *schema.ResourceData, platformType models.PlatformType) (*models.SecurityGroup, diag.Diagnostics) {

	members, err, wrongMemberType := expandSecurityGroupMembers(d.Get("members").(*schema.Set).List(), platformType)
	if err != nil {
		return nil, diag.Errorf("mapResourceDataToSecurityGroup >>>> error while creating security group. Wrong type member: %+v", wrongMemberType)
	}
	return &models.SecurityGroup{
		DisplayName: d.Get("display_name").(string),
		ProjectID:   strfmt.UUID(d.Get("project_id").(string)),
		Members:     members,
	}, nil
}

//##################################################################################################################################
//#####												OPENSTACK TESTS															   #####
//#####													BEGIN																   #####
//##################################################################################################################################

func loadTestOpenstackData() {
	//jsonfile, err := os.ReadFile(pathToFile)

	jsonfile := "{\n\t\"Name\": \"NatPublicIpAddr\",\n\t\"Text\": \"45.130.209.203\"\n}\n{\n\t\"Name\": \"BackupPlanName\",\n\t\"Text\": \"Premium\"\n}\n{\n    \"Name\": \"BackupListName\",\n    \"Text\": \"Standard (1 month / 24h)\"\n}\n{\n\t\"Name\": \"VRF\",\n\t\"Text\": \"VRF-default\"\n}\n{\n\t\"Name\": \"Deployment1DisplayName\",\n\t\"Text\": \"CentOS 7\"\n}\n"
	//if err != nil {
	//	fmt.Printf("pathToFile err %v", err)
	//
	//}

	//if err == nil {

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
		val := reflect.ValueOf(&predefinedTestDataOpenstackDev).Elem().FieldByName(m.Name)

		if val.IsValid() {
			reflect.ValueOf(&predefinedTestDataOpenstackDev).Elem().FieldByName(m.Name).SetString(m.Text)
		}
	}
	//}
}

func getTestOpenstackData() predefinedTestOpenstackData {
	loadTestOpenstackData()
	return predefinedTestDataOpenstackDev
}

type predefinedTestOpenstackData struct {
	//LogicalPort1DisplayName    string
	//Network1Name               string
	//Network2Name               string
	Project1Name string
	//Project2Name               string
	//Project3Name               string
	//Project4Name               string
	//ProjectForVMName           string
	//VirtualMachine1DisplayName string
	//VirtualNetwork1DisplayName string
	//VirtualNetwork2DisplayName string
	//VirtualMachineDisplayName  string
	//VirtualMachine2DisplayName string
	//IPCollection1DisplayName   string
	//Deployment1DisplayName     string
	//CustomService1DisplayName  string
	//CustomService2DisplayName  string
	//KMSKeyDisplayName          string
	//BackupPlanName             string
	//BackupListName             string
	//TagName                    string
	//VPC                        string
	VRF string
	//AutoNatName                string
	//DnatName                   string
	//InfraAdminGroup            string
	//FirewallEWRuleName         string
	//FirewallSNRuleName         string
	//NatPublicIpAddr            string
	//SnapshotName               string
	//AccountName                string
}

var DevTestDataOpenstackPrefix = "tf-gojl"
var predefinedTestDataOpenstackDev = predefinedTestOpenstackData{
	//LogicalPort1DisplayName:    "",
	//BackupPlanName:             "",
	//BackupListName:             "",
	//Network1Name:               fmt.Sprintf("%s-vnet1", DevTestDataOpenstackPrefix),
	//Network2Name:               fmt.Sprintf("%s-vnet2", DevTestDataOpenstackPrefix),
	Project1Name: fmt.Sprintf("%s-project-01", DevTestDataOpenstackPrefix),
	//Project2Name:               fmt.Sprintf("%s-project-02", DevTestDataOpenstackPrefix),
	//Project3Name:               fmt.Sprintf("%s-project-03", DevTestDataOpenstackPrefix),
	//Project4Name:               fmt.Sprintf("%s-project-04", DevTestDataOpenstackPrefix),
	//ProjectForVMName:           fmt.Sprintf("%s-project-01", DevTestDataOpenstackPrefix),
	//VirtualMachine1DisplayName: fmt.Sprintf("%s-vm-default", DevTestDataOpenstackPrefix),
	//VirtualNetwork1DisplayName: fmt.Sprintf("%s-vnet3", DevTestDataOpenstackPrefix),
	//VirtualNetwork2DisplayName: fmt.Sprintf("%s-vnet4", DevTestDataOpenstackPrefix),
	//VirtualMachineDisplayName:  fmt.Sprintf("%s-vm1", DevTestDataOpenstackPrefix),
	//VirtualMachine2DisplayName: fmt.Sprintf("%s-vm2", DevTestDataOpenstackPrefix),
	//IPCollection1DisplayName:   fmt.Sprintf("%s-ipc-default", DevTestDataOpenstackPrefix),
	//Deployment1DisplayName:     "",
	//CustomService1DisplayName:  fmt.Sprintf("%s-https", DevTestDataOpenstackPrefix),
	//CustomService2DisplayName:  fmt.Sprintf("%s-http", DevTestDataOpenstackPrefix),
	//TagName:                    fmt.Sprintf("%s-t1", DevTestDataOpenstackPrefix),
	//KMSKeyDisplayName:          fmt.Sprintf("%s-key", DevTestDataOpenstackPrefix),
	//VPC:                        fmt.Sprintf("%s-vpc", DevTestDataOpenstackPrefix),
	//AutoNatName:                fmt.Sprintf("%s-autonat", DevTestDataOpenstackPrefix),
	//DnatName:                   fmt.Sprintf("%s-dnat", DevTestDataOpenstackPrefix),
	VRF: "",
	//InfraAdminGroup:            fmt.Sprintf("%s-subt1-InfraAdm", DevTestDataOpenstackPrefix),
	//FirewallEWRuleName:         fmt.Sprintf("%s-tf-fw-ew-http", DevTestDataOpenstackPrefix),
	//FirewallSNRuleName:         fmt.Sprintf("%s-tf-fw-sn-http", DevTestDataOpenstackPrefix),
	//NatPublicIpAddr:            "",
	//SnapshotName:               fmt.Sprintf("%s-snaps001", DevTestDataOpenstackPrefix),
	//AccountName:                fmt.Sprintf("%s-act1", DevTestDataOpenstackPrefix),
}

var TestOpenstackData = getTestOpenstackData()

//##################################################################################################################################
//#####												OPENSTACK TESTS															   #####
//#####													END																       #####
//##################################################################################################################################

var CastTemplateFuncMap = map[string]interface{}{
	"StringsToTFList": CastStringsToTFList,
	"StringTFValue":   CastStringTFValue,
	"UuidTFValue":     CastUuidTFValue,
}

type TestData interface {
	FullResourceName() string
}

func CastDataResourceID(td TestData) string {
	return td.FullResourceName() + ".id"
}

func CastTemplateToString(templateString string, data interface{}) string {
	parsedTemplate, err := template.New("template_name").Funcs(CastTemplateFuncMap).Parse(templateString)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := parsedTemplate.Execute(&tpl, data); err != nil {
		panic(err)
	}

	return tpl.String()
}

func CastStringTFValue(value string) string {
	if isTerraformResourceNameCheck(value) {
		return value
	}

	return fmt.Sprintf("%q", value)
}

func CastUuidTFValue(value strfmt.UUID) string {
	if isTerraformResourceNameCheck(value.String()) {
		return value.String()
	}

	return fmt.Sprintf("%q", value)
}

func CastStringsToTFList(list []string) string {
	builder := strings.Builder{}
	builder.WriteString("[")
	for i := 0; i < len(list); i++ {
		item := list[i]

		if isTerraformResourceNameCheck(item) {
			builder.WriteString(item)
		} else {
			builder.WriteString(fmt.Sprintf("%q", item))
		}

		if i < len(list)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("]")

	return builder.String()
}

func isTerraformResourceNameCheck(item string) bool {
	return strings.HasPrefix(item, "ochk_") || strings.HasPrefix(item, "data.ochk_")
}

/*
	func TestArrayOfStringToTFList(t *testing.T) {
		assert.Equal(t, `["a", "b", "c"]`, CastStringsToTFList([]string{"a", "b", "c"}))
		assert.Equal(t, `["a"]`, CastStringsToTFList([]string{"a"}))
		assert.Equal(t, `["123456780"]`, CastStringsToTFList([]string{"123456780"}))
		assert.Equal(t, `["", ""]`, CastStringsToTFList([]string{"", ""}))
		assert.Equal(t, `[]`, CastStringsToTFList(nil))
	}

	func TestExecuteTemplateWithFuncs(t *testing.T) {
		type tmplData struct {
			Strings []string
		}

		assert.Equal(t, `strings = ["a", "b", "c"]`, CastTemplateToString(`strings = {{ StringsToTFList .Strings }}`, tmplData{Strings: []string{"a", "b", "c"}}))
		assert.Equal(t, `strings = []`, CastTemplateToString(`strings = {{ StringsToTFList .Strings }}`, tmplData{Strings: nil}))
	}
*/
func GenerateRandName(devTestDataPrefix string) string {
	return generateShortRandName(devTestDataPrefix)
}

/*
	func GenerateShortRandName(devTestDataPrefix string) string {
		return fmt.Sprintf("%s-%s", devTestDataPrefix, acctest.RandStringFromCharSet(4, acctest.CharSetAlphaNum))
	}
*/
var TestAccProviderFactories map[string]func() (*schema.Provider, error)
var TestAccProviderDefinition *schema.Provider

func init() {
	TestAccProviderDefinition = Provider()
	TestAccProviderFactories = map[string]func() (*schema.Provider, error){
		//nolint:unparam
		"ochk": func() (*schema.Provider, error) {
			return TestAccProviderDefinition, nil
		},
	}
}

/*
func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}*/
