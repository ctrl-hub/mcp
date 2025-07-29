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

## Running with Docker

We publish public docker images to the [releases page](https://github.com/ctrl-hub/mcp/releases) for the latest versions of the container.

The images are hosted on Github, and can be pulled using the following command:

```bash
docker pull ghcr.io/ctrl-hub/mcp:<version>
```

For example:

```bash
docker run ghcr.io/ctrl-hub/mcp:latest
```

## Using an Inspector

If you're just getting started and want to check that the server is operational, you can use the [MCP Inspector](https://github.com/modelcontextprotocol/inspector) to test a running MCP server.

This has the benefit of also generating the servers configuration for you to add to your own MCP client (Claude, ChatGPT etc). A list of compatible clients can be found [here](https://modelcontextprotocol.io/clients).

You can launch it with the following command:

```bash
npx @modelcontextprotocol/inspector
```

The terminal output will contain a link with the auth token prefilled, which will look similar to `http://localhost:6274/?MCP_PROXY_AUTH_TOKEN=71143ab7b5f55e3d618a26b99cd4470d6bbb8d36e140a444beb44cd1f9798d8c`.

Clicking on this will launch the web UI.


### Connecting to the official remote server (recommended)

In the sidebar, change the transport type to `Streamable HTTP` and the URL to `https://api.ctrl-hub.com/mcp`.

You can now press the `Connect` button to establish a connection with the server.


### Connecting to a locally running instance

In the sidebar, change the transport type to `Streamable HTTP` and the URL to `http://localhost:8080` (the default MCP server address).

You can now press the `Connect` button to establish a connection with the server.


## Releasing

[Releases](https://github.com/ctrl-hub/mcp/releases) are managed using GitHub Actions. To create a release, you can tag the repository - this will build binaries from the Go source for multiple platforms and docker images.
