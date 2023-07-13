package templates

import (
	"html/template"
	"log"
	"strings"
)

type InterfaceData struct {
	InterfaceName string
	Methods       []MethodData
}

func GenerateInterface(data InterfaceData) string {

	tmplString := `
type {{.InterfaceName}} interface {
{{- range .Methods}}
	{{.Name}}({{.Params}}) {{.ReturnType}}
	{{- end}}
}
`
	tmpl, err := template.New("interface").Parse(tmplString)
	if err != nil {
		log.Fatalf("Error while parsing template: %v", err)
	}

	var builder strings.Builder
	err = tmpl.Execute(&builder, data)
	if err != nil {
		log.Fatalf("Error while executing template: %v", err)
	}

	return builder.String()
}
