# Go-protoc

## Pure Go version of `protoc`

This repo contains a pure-Go re-implementation of `protoc`.
This project is forked from [jhump/goprotoc](https://github.com/jhump/goprotoc) and modified to hanlde opts flags and to override the `go` and `go_grpc` plugins with built-ins that use `go run`.
This is useful for Go projects that want to avoid a dependency.

The implementaiton delegates to a `protoc` executable on the path, driving it as if it were a plugin, for generating
C++, C#, Objective-C, Java, JavaScript, Python, PHP, and Ruby code (since they are implemented in `protoc` itself).
But it provides descriptors to `protoc`, parsed by `goprotoc`, instead of having `protoc` re-parse all of the source
code.

### Installation

```shell
go get github.com/duckbrain/goprotoc@latest
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Then in your module, you can replace `protoc` with `go run github.com/duckbrain/goprotoc`.
So long as the invocation uses only the `go` and `go_grpc` plugins, you won't need to install any local generation dependencies besides the go toolchain.

```go
//go:generate go run github.com/duckbrain/goprotoc --go_out=./output --go_opt=paths=source_relative --go_grpc_out=./output --go_grpc_opt=paths=source_relative schema.proto
```

### Goals

This project is intended as a replacement for `protoc` in projects that want to avoid requireing that as a dev dependency. It's only intended to work within a Go module where the plugin dependencies are already included.

Future scope could include a way to specify additional plugins that are written in pure Go. With the Go 1.24 [propoal to manage tool dependencies](https://go.googlesource.com/proposal/+/54d6775ff71ccbc00c276db2a4e4841d67011cf4/design/48429-go-tool-modules.md), this will likely be that all plugins will be invoked with `go tool <plugin_name>` instead of the current `go run <import_path>`.
