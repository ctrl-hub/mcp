# Ctrl Hub MCP Server

The Ctrl Hub MCP server connects to the Ctrl Hub Platform. It allows your AI agents, assistants and chatbots access to read data from Ctrl Hub.

It can be ran locally (or on your own infrastructure) using the code provided in this repository or you can use the publicly hosted version at https://api.ctrl-hub.com/mcp.

Transport to the server is via the streamable HTTP protocol. SSE is not supported since the protocol only maintains it for [backward compatibility](https://modelcontextprotocol.io/specification/2025-06-18/basic/transports#backwards-compatibility). stdio will likely be supported in the future, but since the most typical use case is using this server as a remote server, it is not currently supported.

## Building from Source

The application is written in Go and can be built from source using the following command:

```bash
# build the application
go build -o bin/ctrl-hub-mcp-server cmd/server/main.go

# [optionally] run it
./bin/ctrl-hub-mcp-server
```

## Inspector

You can use the [MCP Inspector](https://github.com/modelcontextprotocol/inspector) to test a running MCP server. You can launch it with the following command:

```bash
npx @modelcontextprotocol/inspector
```

The terminal output will contain a link with the auth token prefilled, which will look similar to `http://localhost:6274/?MCP_PROXY_AUTH_TOKEN=71143ab7b5f55e3d618a26b99cd4470d6bbb8d36e140a444beb44cd1f9798d8c`.

Clicking on this will launch the web UI. In the sidebar, change the transport type to `Streamable HTTP` and the URL to `http://localhost:8080` (the default MCP server address).

You can now press the `Connect` button to establish a connection with the server.
