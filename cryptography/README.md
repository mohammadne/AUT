# Cryptography

## protocol-buffer tools

``` bash
# install protocol buffer compiler
sudo dnf install protobuf-compiler

# protocol compiler plugins for Go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# generate protocol buffer & grpc codes
protoc --proto_path=./contracts --go_out=. --go-grpc_out=. service.proto
```
