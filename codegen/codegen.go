package codegen

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dietzy1/pb2go/templates"
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

func createRootFolder(serviceName, githubName, fileName string) error {

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
	if err = createProtoFolder(fileName, githubName, serviceName); err != nil {
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

func createProtoFolder(fileName, githubName, serviceName string) error {

	if err := os.Mkdir(protoFolderName, 0755); err != nil {
		return fmt.Errorf("error: Unable to create the %s folder: %v", protoFolderName, err)
	}

	//Go into the proto folder
	if err := os.Chdir(protoFolderName); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}
	//Call generator functions to create the 2 yaml files

	//Create file for the buf.gen.yaml
	file, err := os.Create(bufGenFileName)
	if err != nil {
		return fmt.Errorf("error: Unable to create the %s file: %v", bufGenFileName, err)
	}

	err = templates.GenerateGenYaml(file)
	if err != nil {
		return fmt.Errorf("error: Unable to generate the %s file: %v", bufGenFileName, err)
	}

	/* 	_, err = file.WriteString(yaml)
	   	if err != nil {
	   		return fmt.Errorf("error: Unable to write to the %s file: %v", bufGenFileName, err)
	   	} */
	file.Close()

	//Create file for the buf.yaml
	file, err = os.Create(bufFileName)
	if err != nil {
		return fmt.Errorf("error: Unable to create the %s file: %v", bufFileName, err)
	}

	err = templates.GenerateYaml(file)
	if err != nil {
		return fmt.Errorf("error: Unable to generate the %s file: %v", bufFileName, err)
	}

	/* _, err = file.WriteString(yaml)
	if err != nil {
		return fmt.Errorf("error: Unable to write to the %s file: %v", bufFileName, err)
	} */
	defer file.Close()

	//Create 2 directories v1 and v2
	if err = os.MkdirAll(fileName+"/v1", 0755); err != nil {
		return fmt.Errorf("error: Unable to create the v1 folder: %v", err)
	}

	if err = os.Chdir(fileName + "/v1"); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	//Copy the proto file into the v1 folder
	if err = copyFile("../../../../"+fileName+".proto", fileName+".proto"); err != nil {
		return fmt.Errorf("error: Unable to copy the proto file: %v", err)
	}

	//Reference
	//option go_package = "github.com/dietzy1/chatapp/services/chatroom/proto/chatroom/v1;chatroomv1";
	optionPackage := fmt.Sprintf("option go_package = \"github.com/%s/%s/proto/%s/v1;%sv1\";", githubName, fileName, fileName, fileName)

	if err = insertSnippetInFile(fileName+".proto", optionPackage); err != nil {
		return fmt.Errorf("error: Unable to insert the snippet in the proto file: %v", err)
	}

	//I need to go back to directories
	if err = os.Chdir("../.."); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	//Here we need to generate the go files from the proto file
	//We need to call the protoc command
	command := exec.Command("buf", "generate")
	_, err = command.Output()
	if err != nil {
		return fmt.Errorf("error: Unable to generate the go files: %v", err)
	}

	//Go back to the root folder
	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	//Create

	//templating.GenerateDomain()

	return nil
}

// Entry function for the codegen package
func Run(serviceName, githubName, fileName string) error {

	if err := createRootFolder(serviceName, githubName, fileName); err != nil {
		return fmt.Errorf("error: Unable to create the root folder: %v", err)
	}

	if err := goModInit(githubName, serviceName); err != nil {
		return fmt.Errorf("error: Unable to create the go.mod file: %v", err)
	}

	return nil
}
