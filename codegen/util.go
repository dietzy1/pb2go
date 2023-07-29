package codegen

import (
	"fmt"
	"io"

	"os"
	"strings"
)

func copyFile(sourcePath, destinationPath string) error {

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("error: Unable to open the source file: %v", err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return fmt.Errorf("error: Unable to create the destination file: %v", err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func insertSnippetInFile(filePath string, snippet string) error {
	// Read the Proto file
	protoFile, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error: Unable to read the Proto file: %v", err)
	}

	// Convert the file content to string
	protoContent := string(protoFile)

	// Find the position to insert the line after the package line
	packageLineIndex := strings.Index(protoContent, "package")
	if packageLineIndex == -1 {
		return fmt.Errorf("error: Unable to find the package line in the Proto file")
	}

	insertIndex := packageLineIndex + strings.Index(protoContent[packageLineIndex:], "\n") + 1

	// Insert the line into the content
	updatedProtoContent := protoContent[:insertIndex] + snippet + "\n" + protoContent[insertIndex:]

	// Write the updated content back to the Proto file
	err = os.WriteFile(filePath, []byte(updatedProtoContent), 0644)
	if err != nil {
		return fmt.Errorf("error: Unable to write to the Proto file: %v", err)
	}

	return nil
}

// Create a temporary file to write the modified
