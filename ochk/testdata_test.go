package ochk

import (
	"bytes"
	"text/template"
)

const (
	testDataVirtualMachine1DisplayName = "devel0000001157"
	testDataIPSet1DisplayName          = "googledns"
)

type ConfigProvider interface {
	ToString() string
	FullResourceName() string
}

func createNewTemplate(name string, templateString string) *template.Template {
	parsedTemplate, err := template.New(name).Parse(templateString)
	if err != nil {
		panic(err)
	}

	return parsedTemplate
}

func executeTemplateToString(template *template.Template, data interface{}) string {
	var tpl bytes.Buffer
	if err := template.Execute(&tpl, data); err != nil {
		panic(err)
	}

	return tpl.String()
}
