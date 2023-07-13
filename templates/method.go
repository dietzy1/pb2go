package templates

import (
	"html/template"
	"log"
	"strings"
)

type MethodData struct {
	TypeName   string
	Name       string
	Params     string
	ReturnType string
}

func GenerateMethod(data MethodData) string {
	tmplString := `

func (t *{{.TypeName}}) {{.Name}}({{.Params}}) {{.ReturnType}} {
	// Method implementation goes here
}

`

	tmpl, err := template.New("file").Parse(tmplString)
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
