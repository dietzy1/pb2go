package generator

import (
	"strings"
)

func GenerateStruct() string {

	var sb strings.Builder

	/* sb.WriteString(fmt.Sprintf("type %s struct {\n", input.ProtoBody[0].Name))
	for _, method := range input.ProtoBody[0].ServiceBody[0].RPCName {
		sb.WriteString(fmt.Sprintf("\t%s\n", method))
	}
	sb.WriteString("}") */

	return sb.String()
}
