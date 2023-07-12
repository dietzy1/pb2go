package templating

import (
	"fmt"
	"html/template"
	"io"
)

type StructData struct {
	StructName string
	Fields     []FieldData
}

type FieldData struct {
	Name string
	Type string
}

func GenerateStruct(w io.Writer, data StructData) error {

	tmplString := `
type {{.StructName}} struct {
{{- range .Fields}}
	{{.Name}} {{.Type}}
{{- end}}
}
`

	tmpl, err := template.New("struct").Parse(tmplString)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}

	fmt.Println("Struct generated successfully.")
	return nil
}
