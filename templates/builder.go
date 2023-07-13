package templates

import (
	"html/template"
	"io"
)

type BuilderData struct {
	PackageName string
	Imports     []string
	Templates   []string
}

func GenerateFile(w io.Writer, data BuilderData) error {
	tmplString := `
package {{.PackageName}}

import (
	{{- range .Imports}}
	"{{.}}"
	{{- end}}
)
	
	{{- range .Templates}}
	{{.}}
	{{- end}}
	`

	tmpl, err := template.New("file").Parse(tmplString)
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}

	return nil

}
