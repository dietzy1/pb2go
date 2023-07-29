package codegen

import (
	"fmt"
	"os"
	"os/exec"
)

func goModInit(githubName, serviceName string) error {

	//Create the go.mod file
	command := exec.Command("go", "mod", "init", fmt.Sprintf("github.com/%s/%s", githubName, serviceName))
	if err := command.Run(); err != nil {
		return fmt.Errorf("error: Unable to create the go.mod file: %v", err)
	}

	return nil
}

func installDependencies() error {

	//print the current working directory
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	command := exec.Command("go", "get", "-d", "./...")

	if err = command.Run(); err != nil {
		return fmt.Errorf("error: Unable to install the dependencies: %v", err)
	}

	return nil
}
