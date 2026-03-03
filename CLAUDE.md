# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Build
go build -o bin/command ./cmd/command
go build -o bin/server ./cmd/server

# Run (stdio/CLI mode)
TRACKER_TOKEN=<token> TRACKER_ORG_ID=<org_id> ./bin/command

# Run (HTTP server mode)
TRACKER_TOKEN=<token> TRACKER_ORG_ID=<org_id> MCP_SECRET=<secret> DOMAIN=<domain> ./bin/server

# Test
go test ./...

# Lint
go vet ./...
```

## Architecture

Three-layer architecture:

**Transport Layer** (`cmd/`)
- `cmd/command/` — stdio transport for direct CLI/MCP client integration
- `cmd/server/` — HTTP server with auto-TLS via Let's Encrypt ACME; requires `MCP_SECRET` bearer auth

**MCP Protocol Layer** (`internal/mcpserver/`)
- `server.go` — creates the MCP server and registers all 7 tools
- Tool handlers delegate directly to the tracker `Client` interface; each tool has a typed `Input`/`Output` struct with `jsonschema` tags

**API Client Layer** (`internal/tracker/`)
- `client.go` — resty-based HTTP client; `authHeader()` auto-detects OAuth vs IAM Bearer token format
- `models.go` — shared data structures
- `issues.go`, `comments.go`, `users.go` — one file per API resource

## Key Conventions

**Environment variables:**
| Var | Purpose |
|-----|---------|
| `TRACKER_TOKEN` | OAuth or IAM token (auto-detected) |
| `TRACKER_ORG_ID` | Yandex 360 org ID (`X-Org-ID` header) |
| `TRACKER_CLOUD_ORG_ID` | Cloud org ID (`X-Cloud-Org-ID`, takes priority over OrgID) |
| `MCP_SECRET` | Bearer token for HTTP server auth |
| `DOMAIN` | Domain for ACME TLS certificate |

**MCP tool pattern:** Every tool follows `func (t *toolServer) ToolName(ctx, *mcp.CallToolRequest, InputStruct) (*mcp.CallToolResult, OutputStruct, error)`.

**Adding a new API resource:** Create a file in `internal/tracker/` implementing the operations, add methods to the `Client` interface in `client.go`, then create a corresponding handler file in `internal/mcpserver/` and register the tool in `server.go`.

**Base URL:** `https://api.tracker.yandex.net/v3`
