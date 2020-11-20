package ochk

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"text/template"
)

var templateFuncMap = map[string]interface{}{
	"StringsToTFList": stringsToTFList,
	"StringTFValue":   stringTFValue,
}

type TestData interface {
	FullResourceName() string
}

func testDataResourceID(td TestData) string {
	return td.FullResourceName() + ".id"
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

func stringTFValue(value string) string {
	if isTerraformResourceName(value) {
		return value
	}

	return fmt.Sprintf("%q", value)
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
