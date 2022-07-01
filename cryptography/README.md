# Cryptography

## How it works

> here we are using 3 gRPC servers responsible for `Alice`, `Bob` and `Eve`.
>
> the `Eve` server is also a moderator server which sees and check all the messages in transient.
>
> in this iplementation we try to handle this scenario interactively by using `cobra` package and setting up 3 different serves to communicate, the `Alice` and `Bob` are gRPC servers and `Eve` is a gRPC client server to monitor the conversation between `Alice` and `Bob`.
>
> we use 2 RSA key-pair for each side for RSA asymetric key and RSA signiture and share the public keys with each other.
>
> the key genertion scenario is on the fly and we don't use any predefing key and the key for each run is different from the pre-run.
>
> the [architecture](https://viewer.diagrams.net/?tags=%7B%7D&highlight=FFFFFF&edit=_blank&layers=1&nav=1&title=drawio#Uhttps%3A%2F%2Fraw.githubusercontent.com%2Fmohammadne%2Funiversity%2Fmaster%2Fcryptography%2Fcontracts%2Farchitecture.drawio) for transefering each message has been drawn here. (click the highlighted link)

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
