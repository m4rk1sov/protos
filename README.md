# PROTOS
contracts for microservices
## GRPC and Protobuffer package dependencies
### Protobuf installation
Linux, using apt or apt-get, for example:
```
apt install -y protobuf-compiler
protoc --version  # Ensure compiler version is 3+
```
MacOS, using Homebrew:
``` 
brew install protobuf
protoc --version  # Ensure compiler version is 3+
```
Windows, using Winget
```
winget install protobuf
protoc --version # Ensure compiler version is 3+
```
### Go plugins for the protocol compiler:
```
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
NOTE: You should add the `protoc-gen-go-grpc` to your PATH
```
PATH="${PATH}:${HOME}/go/bin"
```
### Running the service (generating the contracts)
Requires the Taskfile, instructions can be found here: https://taskfile.dev/installation/
```
task generate
```