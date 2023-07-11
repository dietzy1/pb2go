package templating

// there is issues with the tab character in the string builder
func GenerateYaml() string {
	return `version: v1
	breaking:
	  use:
		- FILE
	lint:
	  use:
		- DEFAULT`
}

func GenerateGenYaml() string {
	return `version: v1
	plugins:
	  - name: go
		out: .
		opt:
		  - paths=source_relative
	  - name: go-grpc
		out: .
		opt:
		  - paths=source_relative
		  - require_unimplemented_servers=false`
}
