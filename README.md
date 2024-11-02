# Go-protoc

## Pure Go version of `protoc`

This repo contains a pure-Go re-implementation of `protoc`.
This new version of `protoc`, named `goprotoc`, will use `go run` to execute the `protoc-gen-go` and `protoc-gen-go-grpc` plugins, expecting to be in a Go module.
This is useful for Go projects that want to avoid a dependency.

The implementaiton delegates to a `protoc` executable on the path, driving it as if it were a plugin, for generating
C++, C#, Objective-C, Java, JavaScript, Python, PHP, and Ruby code (since they are implemented in `protoc` itself).
But it provides descriptors to `protoc`, parsed by `goprotoc`, instead of having `protoc` re-parse all of the source
code.
