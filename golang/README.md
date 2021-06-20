# Summary

Resources related to backend services that support the Mono iOS and Android clients can be found here.

## Structure

```
server
├── cmd
│   └── dummy
│       ├── BUILD.bazel
│       └── main.go
├── pkg
│   └── dummy
│       └── v1
│           ├── BUILD.bazel
│           └── dummy.proto
├── Dockerfile
├── Dockerfile.test
├── BUILD
├── WORKSPACE
├── tools.go
├── go.mod
└── go.sum
```

* cmd

   `cmd` contains service entry points in the `main.go`. Each directory in `cmd` maps to a similarly named service.

* pkg

   `pkg` contains shared code, libraries, and service definitions. In keeping with the [`buf`](https://buf.build/docs/style-guide/#files-and-packages) styleguide,
   protobuf messages and services are in versioned packages.

## Getting started

With the introduction of the [Bazel](https://bazel.build/) system, getting up and running is fairly simple. The first of two requirements is that `bazel` be installed like so:

```
$ brew install bazel
```
Next, install [`docker`](https://docs.docker.com/docker-for-mac/install/) manually by following the linked steps.

### Existing services

An existing service can be brought up individually using the
```
$ bazel run //cmd/[SERVICE]:[SERVICE]
```
command. To bring up all existing services in containers, just run
```
$ make coldstart
```

### New services

Start by creating a versioned service definition in a package named after the service (i.e. `[SERVICE]/v1/[SERVICE].proto`). Be sure to follow
the [`buf`](https://buf.build/docs/style-guide/#services) styleguide when declaring RPCs and messages. This versioned package should be under `pkg`.

Then, create a `BUILD.bazel` that has instructions for how to build the service. The `BUILD.bazel` for the `dummer` service is a good example of a service with both external and local
dependencies. The `dummy` service contains an example for a service with no local dependencies.

Finally, create a new service entrypoint and `BUILD.bazel` file in a package named after the service under `cmd` (i.e. `cmd/[SERVICE]/main.go`). Once service endpoints
are implemented, use
```
$ bazel run //cmd/[SERVICE]:SERVICE
```
to bring up the service locally. Also, be sure to add the service configuration to the `docker-compose.yml` file so that the new service can be brought up in a container.

## [Deprecated] Getting started

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

## Helpful notes

[Bazel protobuf docs](https://github.com/bazelbuild/rules_proto)

[Bazel gazelle docs](https://github.com/bazelbuild/bazel-gazelle/blob/master/repository.rst)

[Bazel go docs](https://github.com/bazelbuild/rules_go)

[Bazel go + protobuf docs](https://github.com/bazelbuild/rules_go/blob/master/proto/core.rst)

[gRPC gateway docs](https://github.com/grpc-ecosystem/grpc-gateway)

[Buf docs](https://buf.build/docs/introduction)

