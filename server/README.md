# Summary

Resources related to backend services that support the Mono iOS and Android clients can be found here.

## Structure

```
server
├── cmd
│   └── dummy
│       └── main.go
├── idl
│   └── dummy.proto
├── pkg
│   └── dummy_gen
│       ├── dummy.pb.go
│       └── dummy.pb.gw.go
├── Dockerfile
├── Dockerfile.test
├── go.mod
└── go.sum
```

* cmd

   `cmd` contains service entry points in the `main.go`. Each directory in `cmd` maps to a similarly named service.

* idl

   `idl` contains service and message definitions in `.proto` files.

* pkg

   `pkg` contains shared code and libraries. Typically, this will mean the generated gRPC service and protocol buffer messages.

## Getting started

It's assumed that the Go environment (`$GOPATH` and `$GOBIN`) have been correctly setup. In order to get started, clone this repo and install the following tools:
* [`protoc`](https://grpc.io/docs/protoc-installation/) - This is the protocol buffer compiler
* [`protoc-gen-go`](https://grpc.io/docs/languages/go/quickstart/) - This is a compiler plugin to `protoc` that enables the compiler to generate Go code
* [`grpc-gateway`](https://github.com/grpc-ecosystem/grpc-gateway/) - This compiler plugin to `protoc` generates a reverse-proxy server which translates a RESTful HTTP API into gRPC

```
$ brew install protobuf
$ GO111MODULE=on go get github.com/golang/protobuf/protoc-gen-go
$ GO111MODULE=on go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.14.7
```

* [`docker`](https://docs.docker.com/docker-for-mac/install/) - This is a containerization tool that services will run in

   Install `docker` manually by following the steps linked above

* [`googleapis`](https://github.com/googleapis/googleapis) - This repo provides some helpful tools and `.proto` files for transcoding HTTP/JSON to gRPC

```
$ go get github.com/googleapis/googleapis@v0.0.0-20200814034631-3a54e988edcb
```

After these steps, you should be able to bring up existing services using `make start` and `make stop`.

To define a new service, write a definition for that service i.e. `[NAME].proto` in `idl` and create an entry point i.e. `main.go` in a new directory
under `cmd/[NAME]`. Then, using `make proto SERVICE="[NAME]"` where `SERVICE` is the name of the service's definition and entry point i.e. `dummy`, generate
the service stubs. These can be found in `pkg/[NAME]_gen`. The final step is to implement the business logic in the entry point that was created earlier.

