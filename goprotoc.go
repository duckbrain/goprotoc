// Command protoc is 100% Go implementation of protoc. It can generate
// code by invoking other plugins, shelling out to external programs in
// the same way that the standard protoc does. It can also link in Go
// plugins that register protoc plugins via plugins.RegisterPlugin during
// their initialization. It aims to provide much of the same functionality
// as protoc, including the ability to read and write descriptors and to
// encode and decode files that contain text- or binary-encoded protocol
// buffer messages.
//
// Unlike the standard protoc, it does not provide any builtin code
// generation logic: it can only execute plugins to generate code. In order
// to generate code that is built into the standard protoc (such as Python,
// C++, Java, etc), this program can shell out to the standard protoc,
// driving it as if it were a plugin. In this mode, it provides to protoc
// the file descriptors it has already parsed, instead of asking protoc to
// re-parse all of the source code.
package main

import (
	"os/exec"

	"github.com/duckbrain/goprotoc/internal/goprotoc"
	"github.com/duckbrain/goprotoc/internal/plugins"
)

func main() {
	goprotoc.RegisterPlugin("go", func(req *plugins.CodeGenRequest, res *plugins.CodeGenResponse) error {
		return plugins.Exec(exec.Command("go", "run", "google.golang.org/protobuf/cmd/protoc-gen-go"), req, res)
	})
	goprotoc.RegisterPlugin("go_grpc", func(req *plugins.CodeGenRequest, res *plugins.CodeGenResponse) error {
		return plugins.Exec(exec.Command("go", "run", "google.golang.org/grpc/cmd/protoc-gen-go-grpc"), req, res)
	})
	goprotoc.Main()
}
