// Generates Bitcoin Core's RPC documentation.
//
// What is necessary to run this:
//   (1) install golang
//   (2) install bitcoin core, set it up to use regtest
//   (3) run bitcoind
//   (4) from contrib/doc-gen, with bitcoin-cli in PATH, run `go run .`
//   (5) add the generated files to git
package main

import (
	"io"
	"log"
	"os"
)

type CommandData struct {
	Version     string
	Name        string
	Description string
	Group       string
	Permalink   string
}

func main() {
	generateRPC()
}

func open(path string) io.Writer {
	f, err := os.Create(path)
	// not closing, program will close sooner
	if err != nil {
		log.Fatalf("Cannot open file %s: %s", path, err.Error())
	}
	return f
}
