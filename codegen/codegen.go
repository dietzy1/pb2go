package codegen

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/dietzy1/pb2go/parser"
	"github.com/dietzy1/pb2go/templates"
)

func Run(proto parser.Service) error {

	//Create new builder object should only be instantiated once
	builder, err := templates.NewBuilder(proto)
	if err != nil {
		return err
	}

	//Create root folder
	if err := os.MkdirAll(proto.ServiceName, 0755); err != nil {
		return err
	}

	//Change directory to the root folder
	if err := os.Chdir(proto.ServiceName); err != nil {
		return err
	}

	//Generate folders and files
	if err := generateAndReturn(builder); err != nil {
		return err
	}

	//Generate files
	if err := generateProtobuffers(proto.FileName, proto.GithubName, proto.ServiceName); err != nil {
		return err
	}

	//instantiate go mod file
	if err := goModInit(proto.GithubName, proto.ServiceName); err != nil {
		return err
	}

	/* if err := installDependencies(); err != nil {
		return fmt.Errorf("error: Unable to install the dependencies: %v", err)
	} */

	return nil
}

func generateProtobuffers(fileName, githubName, serviceName string) error {
	//Go into the proto folder
	if err := os.Chdir("proto"); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	//Create 2 directories v1 and v2
	if err := os.MkdirAll(fileName+"/v1", 0755); err != nil {
		return fmt.Errorf("error: Unable to create the v1 folder: %v", err)
	}

	if err := os.Chdir(fileName + "/v1"); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	//Copy the proto file into the v1 folder
	if err := copyFile("../../../../"+fileName+".proto", fileName+".proto"); err != nil {
		return fmt.Errorf("error: Unable to copy the proto file: %v", err)
	}

	//Reference
	//option go_package = "github.com/dietzy1/chatapp/services/chatroom/proto/chatroom/v1;chatroomv1";
	optionPackage := fmt.Sprintf("option go_package = \"github.com/%s/%s/proto/%s/v1;%sv1\";", githubName, fileName, fileName, fileName)

	if err := insertSnippetInFile(fileName+".proto", optionPackage); err != nil {
		return fmt.Errorf("error: Unable to insert the snippet in the proto file: %v", err)
	}

	//I need to go back to directories
	if err := os.Chdir("../../"); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
	}

	//Here we need to generate the go files from the proto file
	//We need to call the protoc command
	command := exec.Command("buf", "generate")
	_, err := command.Output()
	if err != nil {
		return fmt.Errorf("error: Unable to generate the go files: %v", err)
	}

	if err := os.Chdir("../"); err != nil {
		return fmt.Errorf("error: Unable to change directory: %v", err)
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
