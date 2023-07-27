package codegen

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dietzy1/pb2go/parser"
	"github.com/dietzy1/pb2go/templates"
)

func Run(proto parser.Service) error {

	// print working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wd)

	//Create new builder object should only be instantiated once
	builder, err := templates.NewBuilder(proto)
	if err != nil {
		return err
	}

	if err := generateAndReturn(builder); err != nil {
		return err
	}

	return nil
}

// GenerateAndReturn creates a directory, navigates to that directory,
// calls the generation function, and returns to the prior directory.
func generateAndReturn(builder *templates.Builder) error {

	//Get directories we are generating
	directories := templates.DeclaredDirectories()

	for _, dir := range directories {

		// Get the current working directory
		prevDir, err := os.Getwd()
		if err != nil {
			return err
		}

		// Create the target directory
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}

		// Change to the target directory
		if err := os.Chdir(dir); err != nil {
			return err
		}

		for _, v := range builder.Templates[dir] {
			fmt.Println(v.Name())

			fileName := strings.TrimSuffix(v.Name(), ".tmpl")

			// create a file object
			file, err := os.Create(fileName)
			if err != nil {
				return err
			}
			defer file.Close()

			//remove 5 characters from the end of the string

			if err := builder.Build(file, dir, v.Name()); err != nil {
				return err
			}
			log.Println("Generated", v.Name())
		}
		if err := os.Chdir(prevDir); err != nil {
			fmt.Printf("error: Unable to change directory: %v", err)
		}
		log.Println("Returned to", prevDir)

	}

	return nil
}
