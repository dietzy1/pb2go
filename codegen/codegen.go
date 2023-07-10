package codegen

import (
	"fmt"
	"os"
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

func CreateRootFolder(serviceName string) error {
	err := os.Mkdir(serviceName, 0755)
	if err != nil {
		return fmt.Errorf("error: Unable to create the root folder: %v", err)
	}

	//Go into the root folder
	if err = os.Chdir("myfolder"); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	//Create the server folder
	if err = createServerFolder(); err != nil {
		return fmt.Errorf("error: Unable to create the server folder: %v", err)
	}

	//Create the repository folder
	if err = createRepositoryFolder(); err != nil {
		return fmt.Errorf("error: Unable to create the repository folder: %v", err)
	}

	//Create the domain folder
	if err = createDomainFolder(); err != nil {
		return fmt.Errorf("error: Unable to create the domain folder: %v", err)
	}

	//Create the proto folder
	if err = createProtoFolder(); err != nil {
		return fmt.Errorf("error: Unable to create the proto folder: %v", err)
	}

	return nil
}

func createServerFolder() error {
	return nil
}

func createRepositoryFolder() error {
	return nil
}

func createDomainFolder() error {
	return nil
}

func createProtoFolder() error {
	return nil
}
