package templating

import (
	"fmt"
	"html/template"
	"io"
)

type InterfaceData struct {
	InterfaceName string
	Methods       []MethodData
}

type MethodData struct {
	Name       string
	Params     string
	ReturnType string
}

func GenerateInterface(w io.Writer, data InterfaceData) error {

	tmplString := `
type {{.InterfaceName}} interface {
{{- range .Methods}}
	{{.Name}}({{.Params}}) {{.ReturnType}}
	{{- end}}
}
`

	tmpl, err := template.New("interface").Parse(tmplString)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}

	fmt.Println("Interface generated successfully.")
	return nil
}
