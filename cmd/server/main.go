package main

import "github.com/ctrl-hub/mcp/internal/server"

var (
	version = "1.0.0"
)

func main() {

	server.NewServer(version)

}
