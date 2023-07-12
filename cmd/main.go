package main

import (
	"errors"
	"flag"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/dietzy1/pb2go/codegen"
	"github.com/dietzy1/pb2go/templating"
	"github.com/dietzy1/pb2go/tui"
)

//List of issues that needs to be dealt with
//Should simply start by only supporting

//TODO: Figure out how many layers are needed

//Protobuf layer // should be generated
//protobuf server methods generated from the proto file

//Validation layer // should be generated
//Basicly just functions that checks if the input is valid

//Business layer // should be generated

//SQL layer // should be generated

func main() {

	start := time.Now()
	if err := run(); err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	log.Printf("It took: %s to generate the project", elapsed)

}

//First i need to prompt the user if its ok to install a bunch of shit
//yes / no

//Then the user should be asked what their github is

//Some prompt should show up with different configurations of the program (for now we simply stick to simple 1 way)

//I need to prompt the user for their github

func run() error {

	//Start the TUI application
	if err := tui.Run(); err != nil {
		return err
	}

	// Define a command-line flag for the protobuf file path
	protoFilePath := flag.String("proto", "", "Path to the protobuf file")
	githubName := flag.String("github", "", "Github name of the user")

	//I should probaly read in the file early so I can display to the user that the project will be
	//Generated with whatever is infront of the test.proto
	//At the end the user should be prompted if they are satisfied with the result and that the test.proto file should be deleted.

	//There is going to happen alot of code generation idk how long it is going to take so it might be
	//good to

	flag.Parse()

	//Validate the extension of the file
	if err := validate(*protoFilePath); err != nil {
		return err
	}
	if err := tempValidate(*githubName); err != nil {
		return err
	}

	// Process the protobuf file
	parsedProto, err := codegen.Parse(*protoFilePath)
	if err != nil {
		return err
	}

	/* 	input := templating.InterfaceData{
	   		InterfaceName: "TestInterface",
	   		Methods: []templating.MethodData{
	   			{
	   				Name:       "TestMethod",
	   				Params:     "input string",
	   				ReturnType: "string",
	   			},
	   			{
	   				Name:       "TestMethod2",
	   				Params:     "input string2",
	   				ReturnType: "string2",
	   			},
	   		},
	   	}

	   	file, err := os.Create("generated_interface.go")
	   	if err != nil {
	   		return err
	   	} */

	//Call in templating writer functions here
	if err := templating.GenerateInterface(file, input); err != nil {
		return err
	}

	_ = parsedProto

	/* if err := codegen.Run(parsedProto.ServiceName, *githubName, parsedProto.FileName); err != nil {
		return err
	} */

	return nil
}

func validate(path string) error {
	// Check if the protobuf file path is provided
	if path == "" {
		return errors.New("error: Please provide the path to the protobuf file using the -proto flag")
	}

	// Check if the file exists
	ex := strings.ToLower(filepath.Ext(path))

	if ex != ".proto" {
		return errors.New("error: Please provide a valid protobuf file")
	}

	return nil
}

func tempValidate(github string) error {
	if github == "" {
		return errors.New("error: Please provide your github name using the -github flag")
	}

	return nil
}
