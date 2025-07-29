# Ctrl Hub MCP Server

The Ctrl Hub MCP server connects to the Ctrl Hub Platform. It allows your AI agents, assistants and chatbots access to read data from Ctrl Hub.

It can be ran locally (or on your own infrastructure) using the code provided in this repository or you can use the publicly hosted version at https://api.ctrl-hub.com/mcp.

## Building from Source

The application is written in Go and can be built from source using the following command:

```
go build -o bin/ctrl-hub-mcp-server cmd/server/main.go
```
