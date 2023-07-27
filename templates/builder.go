package templates

import (
	"embed"
	"fmt"
	"io"
	"text/template"

	"github.com/dietzy1/pb2go/parser"
)

//go:embed cmd/*.tmpl datastore/*.tmpl domain/*.tmpl proto/*.tmpl server/*.tmpl
var templates embed.FS

type Builder struct {
	//First key is the directory, second key is the template name
	Templates map[string]map[string]*template.Template
	service   parser.Service
}

func NewBuilder(parser parser.Service) (*Builder, error) {

	b := &Builder{
		Templates: make(map[string]map[string]*template.Template),
		service:   parser,
	}

	if err := b.parse(); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *Builder) Build(w io.Writer, directory string, file string) error {
	//use directory and template to get the template from the map
	t := b.Templates[directory][file]

	// print temp to stdout
	if err := t.Execute(w, b.service); err != nil {
		return fmt.Errorf("error: Unable to build template using the data: %v", err)
	}

	return nil
}

func DeclaredDirectories() []string {
	return []string{
		"cmd",
		"datastore",
		"domain",
		"proto",
		"server",
	}
}

// Parse parses declared templates.
func (b *Builder) parse() error {

	for _, dir := range DeclaredDirectories() {
		files, err := templates.ReadDir(dir)
		if err != nil {
			return fmt.Errorf("internal error: Unable to read the directory: %v", err)
		}

		// Initialize the inner map for this directory
		if b.Templates[dir] == nil {
			b.Templates[dir] = make(map[string]*template.Template)
		}

		for _, file := range files {
			fmt.Println(file.Name())

			t := template.New(file.Name())

			template, err := t.ParseFS(templates, fmt.Sprintf("%s/%s", dir, file.Name()))
			if err != nil {
				return fmt.Errorf("internal error: Unable to parse the template: %v", err)
			}
			b.Templates[dir][file.Name()] = template

		}

	}

	return nil
}
