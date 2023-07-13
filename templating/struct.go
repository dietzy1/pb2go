package templating

import (
	"html/template"
	"log"
	"strings"
)

type StructData struct {
	StructName string
	Fields     []FieldData
}

type FieldData struct {
	Name string
	Type string
}

func GenerateStruct(data StructData) string {

	tmplString := `
type {{.StructName}} struct {
{{- range .Fields}}
	{{.Name}} {{.Type}}
{{- end}}
}
`

	tmpl, err := template.New("struct").Parse(tmplString)
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
