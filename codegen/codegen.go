package codegen

import (
	"fmt"
	"os"

	"github.com/dietzy1/pb2go/generator"
)

//The field serviceName should be used to generate the file names for the generated code

//this file is responsible for creating the structure of the project

//Root folder should be called the name of the service

//Adapers folder
// -- server
// ....server.go
// ....handlers.go
// ....handlers_test.go
// ....middleware.go

// -- repository
//.... repository.go
//.... serviceName.go

//Domain folder
//-- validation
//....validation.go
//....validation_test.go
//.... domain.go

//Proto folder
// --servicename
//....servicename.proto
//buf.yaml
//buf.gen.yaml

const (
	serverFolderName     = "server"
	repositoryFolderName = "repository"
	domainFolderName     = "domain"
	protoFolderName      = "proto"

	bufFileName    = "buf.yaml"
	bufGenFileName = "buf.gen.yaml"
)

func CreateRootFolder(serviceName string, fileName string) error {
	err := os.Mkdir(serviceName, 0755)
	if err != nil {
		return fmt.Errorf("error: Unable to create the root folder: %v", err)
	}

	//Go into the root folder
	if err = os.Chdir(serviceName); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	//Create the server folder
	if err = createServerFolder(); err != nil {
		return fmt.Errorf("error: Unable to create the %s folder: %v", serverFolderName, err)
	}

	//Create the repository folder
	if err = createRepositoryFolder(); err != nil {
		return fmt.Errorf("error: Unable to create the %s folder: %v", repositoryFolderName, err)
	}

	//Create the domain folder
	if err = createDomainFolder(); err != nil {
		return fmt.Errorf("error: Unable to create the %s folder: %v", domainFolderName, err)
	}

	//Create the proto folder
	if err = createProtoFolder(fileName); err != nil {
		return fmt.Errorf("error: Unable to create the %s folder: %v", protoFolderName, err)
	}

	return nil
}

func createServerFolder() error {

	if err := os.Mkdir(serverFolderName, 0755); err != nil {
		return fmt.Errorf("error: Unable to create the %s folder: %v", serverFolderName, err)
	}

	//Go into the server folder
	if err := os.Chdir(serverFolderName); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	//Do some stuff here

	//Go back to the root folder
	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	return nil
}

func createRepositoryFolder() error {

	if err := os.Mkdir(repositoryFolderName, 0755); err != nil {
		return fmt.Errorf("error: Unable to create the %s folder: %v", repositoryFolderName, err)
	}

	//Go into the repository folder
	if err := os.Chdir(repositoryFolderName); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	//Do some stuff here

	//Go back to the root folder
	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	return nil
}

func createDomainFolder() error {

	if err := os.Mkdir(domainFolderName, 0755); err != nil {
		return fmt.Errorf("error: Unable to create the %s folder: %v", domainFolderName, err)
	}

	//Go into the domain folder
	if err := os.Chdir(domainFolderName); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	//Do some stuff here

	//Go back to the root folder
	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	return nil
}

func createProtoFolder(serviceName string) error {

	if err := os.Mkdir(protoFolderName, 0755); err != nil {
		return fmt.Errorf("error: Unable to create the %s folder: %v", protoFolderName, err)
	}

	//Go into the proto folder
	if err := os.Chdir(protoFolderName); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}
	//Call generator functions to create the 2 yaml files

	//Create file for the buf.gen.yaml
	file, err := os.Create("buf.gen.yaml")
	if err != nil {
		return fmt.Errorf("error: Unable to create the %s file: %v", bufGenFileName, err)
	}

	_, err = file.WriteString(generator.GenerateGenYaml())
	if err != nil {
		return fmt.Errorf("error: Unable to write to the %s file: %v", bufGenFileName, err)
	}
	file.Close()

	//Create file for the buf.yaml
	file, err = os.Create("buf.yaml")
	if err != nil {
		return fmt.Errorf("error: Unable to create the %s file: %v", bufFileName, err)
	}
	_, err = file.WriteString(generator.GenerateYaml())
	if err != nil {
		return fmt.Errorf("error: Unable to write to the %s file: %v", bufFileName, err)
	}
	defer file.Close()

	//Create 2 directories v1 and v2
	if err = os.MkdirAll("v1/"+serviceName, 0755); err != nil {
		return fmt.Errorf("error: Unable to create the v1 folder: %v", err)
	}

	//Go back to the root folder
	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	//Create

	return nil
}
