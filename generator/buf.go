package generator

func generateYaml() string {
	return `version: v1
	breaking:
	  use:
		- FILE
	lint:
	  use:
		- DEFAULT`
}

func generateGenYaml() string {
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