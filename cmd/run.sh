#!/bin/bash
rm -rf ./TesterService
#rm generated_interface.go

go build -o pb2go

./pb2go -proto ./test.proto -github dietzy1


