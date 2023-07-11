package templating

import (
	"fmt"

	"sigs.k8s.io/yaml"
)

// there is issues with the tab character in the string builder
func GenerateYaml() (string, error) {
	data := map[string]interface{}{
		"version": "v1",
		"breaking": map[string]interface{}{
			"use": []string{"FILE"},
		},
		"lint": map[string]interface{}{
			"use": []string{"DEFAULT"},
		},
	}

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error: Unable to marshal yaml: %v", err)
	}

	return string(yamlData), nil
}

func GenerateGenYaml() (string, error) {
	data := map[string]interface{}{
		"version": "v1",
		"plugins": []map[string]interface{}{
			{
				"name": "go",
				"out":  ".",
				"opt":  []string{"paths=source_relative"},
			},
			{
				"name": "go-grpc",
				"out":  ".",
				"opt": []string{
					"paths=source_relative",
					"require_unimplemented_servers=false",
				},
			},
		},
	}

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error: Unable to marshal yaml: %v", err)
	}
	return string(yamlData), err

}
