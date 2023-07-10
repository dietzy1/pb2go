package generator

import (
	"fmt"
	"strings"
)

func GenerateInterface(name string, methods []Method) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("type %s interface {\n", name))
	for _, method := range methods {
		sb.WriteString(fmt.Sprintf("\t%s\n", method))
	}
	sb.WriteString("}")

	return sb.String()
}

// Method represents a method definition.
type Method struct {
	Name       string
	Parameters []Parameter
	ReturnType string
}

// Parameter represents a parameter definition.
type Parameter struct {
	Name string
	Type string
}
