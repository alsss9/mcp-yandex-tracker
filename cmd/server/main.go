package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/auth"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sunnyyssh/mcp-yandex-tracker/internal/mcpserver"
	"github.com/sunnyyssh/mcp-yandex-tracker/internal/tracker"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	secret := os.Getenv("MCP_SECRET")
	if secret == "" {
		log.Fatal("MCP_SECRET env var is required")
	}
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		log.Fatal("DOMAIN env var is required")
	}

	trackerClient := tracker.New(tracker.Config{
		Token:      os.Getenv("TRACKER_TOKEN"),
		OrgID:      os.Getenv("TRACKER_ORG_ID"),
		CloudOrgID: os.Getenv("TRACKER_CLOUD_ORG_ID"),
	})

	server := mcpserver.NewServer(trackerClient)

	verifier := func(_ context.Context, token string, _ *http.Request) (*auth.TokenInfo, error) {
		if token != secret {
			return nil, auth.ErrInvalidToken
		}
		return &auth.TokenInfo{}, nil
	}

	mcpHandler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
		return server
	}, nil)

	http.Handle("/mcp", auth.RequireBearerToken(verifier, nil)(mcpHandler))

	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain),
		Cache:      autocert.DirCache("/var/cache/autocert"),
	}

	// HTTP server redirects to HTTPS and handles ACME HTTP-01 challenges.
	go func() {
		log.Fatal(http.ListenAndServe(":80", m.HTTPHandler(nil)))
	}()

	tlsServer := &http.Server{
		Addr:      ":443",
		TLSConfig: m.TLSConfig(),
	}
	log.Printf("listening on https://%s", domain)
	log.Fatal(tlsServer.ListenAndServeTLS("", ""))
}
