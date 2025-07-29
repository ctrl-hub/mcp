package server

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var (
	version = "1.0.0"
	addr    = flag.String("addr", ":8080", "HTTP server address")
	port    = flag.String("port", "", "HTTP server port (overrides addr)")
)

func NewServer(version string) error {
	flag.Parse()

	if *port != "" {
		*addr = ":" + *port
	}

	server := mcp.NewServer(&mcp.Implementation{Name: "ctrl-hub", Title: "Ctrl Hub", Version: version}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "listOrganisations", Description: "Lists the organisations that the user can access", Title: "List Organisations"}, ListOrganisations)
	server.AddPrompt(&mcp.Prompt{Name: "listOrganisations"}, PromptListOrganisations)
	server.AddResource(&mcp.Resource{
		Name:     "info",
		MIMEType: "text/plain",
		URI:      "embedded:info",
	}, handleEmbeddedResource)

	handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
		return server
	}, nil)
	log.Printf("MCP server listening at %s", *addr)
	return http.ListenAndServe(*addr, handler)
}

type Pagination struct {
	Page  int `json:"page" jsonschema:"page"`
	Limit int `json:"limit" jsonschema:"limit"`
}

type OrganisationArgs struct {
	Pagination Pagination `json:"pagination" jsonschema:"pagination"`
}

func ListOrganisations(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[OrganisationArgs]) (*mcp.CallToolResultFor[struct{}], error) {

	req, err := http.NewRequest(http.MethodGet, "https://api.ctrl-hub.dev/v3/orgs", nil)
	if err != nil {
		return nil, err
	}

	// TODO: implement auth
	// req.Header.Set("X-Session-Token", os.Getenv("SESSION_TOKEN"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}

	return &mcp.CallToolResultFor[struct{}]{
		Content: []mcp.Content{
			&mcp.TextContent{Text: buf.String()},
		},
	}, nil
}

func PromptListOrganisations(ctx context.Context, ss *mcp.ServerSession, params *mcp.GetPromptParams) (*mcp.GetPromptResult, error) {
	return &mcp.GetPromptResult{
		Description: "List organisations",
		Messages: []*mcp.PromptMessage{
			{Role: "user", Content: &mcp.TextContent{Text: "organisations: " + params.Arguments["page"]}},
		},
	}, nil
}

var embeddedResources = map[string]string{
	"info": "This is the Ctrl Hub MCP Server.",
}

func handleEmbeddedResource(_ context.Context, _ *mcp.ServerSession, params *mcp.ReadResourceParams) (*mcp.ReadResourceResult, error) {
	u, err := url.Parse(params.URI)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "embedded" {
		return nil, fmt.Errorf("incorrect scheme: %q", u.Scheme)
	}
	key := u.Opaque
	text, ok := embeddedResources[key]
	if !ok {
		return nil, fmt.Errorf("no embedded resource named %q", key)
	}
	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{
			{URI: params.URI, MIMEType: "text/plain", Text: text},
		},
	}, nil
}
