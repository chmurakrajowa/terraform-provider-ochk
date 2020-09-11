package ochk

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"text/template"
)

const (
	testDataIPSet1DisplayName          = "ochk1"
	testDataLogicalPort1DisplayName    = "f766455e-22a7-a2de-35f5-b8599f064a08/devel0000000639.vmx@a596ebcb-a875-4188-8ec5-0a84bfbf1e11"
	testDataNetwork1Name               = "vtest8"
	testDataNetwork2Name               = "vtest7"
	testDataSubtenant1Name             = "acmpt_107"
	testDataUser1Name                  = "devel-jpuser"
	testDataVirtualMachine1DisplayName = "devel0000001157"
)

var templateFuncMap = map[string]interface{}{
	"StringsToTFList": stringsToTFList,
}

func executeTemplateToString(templateString string, data interface{}) string {
	parsedTemplate, err := template.New("template_name").Funcs(templateFuncMap).Parse(templateString)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := parsedTemplate.Execute(&tpl, data); err != nil {
		panic(err)
	}

	return tpl.String()
}

func stringsToTFList(list []string) string {
	builder := strings.Builder{}
	builder.WriteString("[")
	for i := 0; i < len(list); i++ {
		item := list[i]

		if isTerraformResourceName(item) {
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

func isTerraformResourceName(item string) bool {
	return strings.HasPrefix(item, "ochk_") || strings.HasPrefix(item, "data.ochk_")
}

func TestArrayOfStringToTFList(t *testing.T) {
	assert.Equal(t, `["a", "b", "c"]`, stringsToTFList([]string{"a", "b", "c"}))
	assert.Equal(t, `["a"]`, stringsToTFList([]string{"a"}))
	assert.Equal(t, `["123456780"]`, stringsToTFList([]string{"123456780"}))
	assert.Equal(t, `["", ""]`, stringsToTFList([]string{"", ""}))
	assert.Equal(t, `[]`, stringsToTFList(nil))
}

func TestExecuteTemplateWithFuncs(t *testing.T) {
	type tmplData struct {
		Strings []string
	}

	assert.Equal(t, `strings = ["a", "b", "c"]`, executeTemplateToString(`strings = {{ StringsToTFList .Strings }}`, tmplData{Strings: []string{"a", "b", "c"}}))
	assert.Equal(t, `strings = []`, executeTemplateToString(`strings = {{ StringsToTFList .Strings }}`, tmplData{Strings: nil}))
}
