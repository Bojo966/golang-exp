# golang-todo app

## Prerequisites

### Golang

- Download and install the latest version of [Go](https://golang.org/doc/install)
- Make sure the following returns a result

```bash
echo $GOPATH
```

### Dependency management system for Go

Dep is the dependency management system used for this project.

- [Installation and Docs](https://golang.github.io/dep/)

### Protobuf

- Install [protoc](https://developers.google.com/protocol-buffers) (platform dependent) and make sure you put it in your path

For more info on gRPC with Go visit <https://grpc.io/docs/quickstart/go/>

## Installation

To get all dependencies, run

```bash
dep ensure
```

Next run the following to be sure that everything is building properly:

```bash
go build
```

## Development

### Running project locally

To run the project locally, run the following

```bash
go run server.go
```

### Middleware

<https://github.com/grpc-ecosystem/go-grpc-middleware>

### Logging

For logging the project uses [zap](https://github.com/uber-go/zap) that is widely adopted along the Golang ecosystem

### Dependencies

To add a new go dependency run the following

```bash
dep ensure -add <your-dependency>
```

### Gen files from proto files

```bash
protoc -I protoc/ protoc/todoapp.proto  --go_out=plugins=grpc:protoc/
```
