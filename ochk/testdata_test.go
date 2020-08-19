package ochk

import (
	"bytes"
	"text/template"
)

const (
	testDataVirtualMachine1DisplayName = "devel0000001157"
	testDataIPSet1DisplayName          = "googledns"
)

func executeTemplateToString(templateString string, data interface{}) string {
	parsedTemplate, err := template.New("template_name").Parse(templateString)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := parsedTemplate.Execute(&tpl, data); err != nil {
		panic(err)
	}

	return tpl.String()
}
