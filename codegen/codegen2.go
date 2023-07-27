package codegen

import (
	"fmt"
	"os"

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

	builder, err := templates.NewBuilder(proto)
	if err != nil {
		return err
	}

	if err := builder.Build(os.Stdout, templates.NewDirectories().Cmd, templates.NewFiles().Main); err != nil {
		return err
	}

	/* if err := os.Mkdir(proto.ServiceName, 0755); err != nil {
		return fmt.Errorf("error: Unable to create the root folder: %v", err)
	}

	if err := createRootFolder(proto.ServiceName, proto.GithubName, proto.FileName); err != nil {
		return fmt.Errorf("error: Unable to create the root folder: %v", err)
	} */

	/* 	if err := goModInit(githubName, serviceName); err != nil {
		return fmt.Errorf("error: Unable to create the go.mod file: %v", err)
	} */

	return nil
}

// GenerateAndReturn creates a directory, navigates to that directory,
// calls the generation function, and returns to the prior directory.
func GenerateAndReturn(dir string, generateFunc func() error) error {
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

	// Call the generation function
	if err := generateFunc(); err != nil {
		return err
	}

	// Return to the previous directory
	if err := os.Chdir(prevDir); err != nil {
		return err
	}

	return nil
}
