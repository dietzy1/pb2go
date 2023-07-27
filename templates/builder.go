package templates

import (
	"embed"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/dietzy1/pb2go/parser"
)

//go:embed cmd/*.tmpl datastore/*.tmpl domain/*.tmpl proto/*.tmpl server/*.tmpl
var templates embed.FS

type directory string

type directories struct {
	Cmd       directory
	Datastore directory
	Domain    directory
	Proto     directory
	Server    directory
}

func NewDirectories() directories {
	return directories{
		Cmd:       "cmd",
		Datastore: "datastore",
		Domain:    "domain",
		Proto:     "proto",
		Server:    "server",
	}
}

type file string

type files struct {
	Main            file
	Datastore       file
	Service         file
	Domain          file
	Validator       file
	BufGenerateYaml file
	BufYaml         file
	Handlers        file
	Middleware      file
	Server          file
}

func NewFiles() files {
	return files{
		Main:            "main.go.tmpl",
		Datastore:       "datastore.go.tmpl",
		Service:         "service.go.tmpl",
		Domain:          "domain.go.tmpl",
		Validator:       "validator.go.tmpl",
		BufGenerateYaml: "buf.gen.yaml.go.tmpl",
		BufYaml:         "buf.yaml.go.tmpl",
		Handlers:        "handlers.go.tmpl",
		Middleware:      "middleware.go.tmpl",
		Server:          "server.go.tmpl",
	}
}

func declaredDirectories() []string {
	return []string{
		"cmd",
		"datastore",
		"domain",
		"proto",
		"server",
	}
}

type builder struct {
	//First key is the directory, second key is the template name
	templates map[string]map[string]*template.Template
	service   parser.Service
}

func NewBuilder(parser parser.Service) (*builder, error) {

	b := &builder{
		templates: make(map[string]map[string]*template.Template),
		service:   parser,
	}

	if err := b.parse(); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *builder) Build(w io.Writer, directory directory, file file) error {

	fmt.Println(string(directory))
	fmt.Println(string(file))

	//use directory and template to get the template from the map
	t := b.templates[string(directory)][string(file)]

	// print temp to stdout
	if err := t.Execute(os.Stdout, nil); err != nil {
		return fmt.Errorf("error: Unable to build template using the data: %v", err)
	}

	return nil
}

// Parse parses declared templates.
func (b *builder) parse() error {

	for _, dir := range declaredDirectories() {
		files, err := templates.ReadDir(dir)
		if err != nil {
			return fmt.Errorf("internal error: Unable to read the directory: %v", err)
		}

		// Initialize the inner map for this directory
		if b.templates[dir] == nil {
			b.templates[dir] = make(map[string]*template.Template)
		}

		for _, file := range files {
			fmt.Println(file.Name())

			t := template.New(file.Name())

			template, err := t.ParseFS(templates, fmt.Sprintf("%s/%s", dir, file.Name()))
			if err != nil {
				return fmt.Errorf("internal error: Unable to parse the template: %v", err)
			}
			b.templates[dir][file.Name()] = template

		}

	}

	return nil
}
