#!/bin/bash
rm -rf ./TesterService

go build -o pb2go

./pb2go -proto ./test.proto -github dietzy1


