package main

import (
	"errors"
	"flag"
	"log"
	"path/filepath"
	"strings"

	"github.com/dietzy1/pb2go/codegen"
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
	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {

	// Define a command-line flag for the protobuf file path
	protoFilePath := flag.String("proto", "", "Path to the protobuf file")
	flag.Parse()

	//Validate the extension of the file
	if err := validate(*protoFilePath); err != nil {
		return err
	}

	// Process the protobuf file
	parsedProto, err := codegen.Parse(*protoFilePath)
	if err != nil {
		return err
	}

	if err := codegen.CreateRootFolder(parsedProto.ServiceName, parsedProto.FileName); err != nil {
		return err
	}

	//Generate nessesary files
	/* 	if err := generator.Generator(parsedProto); err != nil {
	   		return err
	   	}
	*/
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
