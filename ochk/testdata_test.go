package ochk

import (
	"bytes"
	"text/template"
)

const (
	testDataVirtualMachine1DisplayName = "devel0000001157"
	testDataIPSet1DisplayName          = "googledns"
	testDataLogicalPort1DisplayName    = "f766455e-22a7-a2de-35f5-b8599f064a08/devel0000000639.vmx@a596ebcb-a875-4188-8ec5-0a84bfbf1e11"
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
