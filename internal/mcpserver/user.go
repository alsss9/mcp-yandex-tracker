package mcpserver

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sunnyyssh/mcp-yandex-tracker/internal/tracker"
)

type GetCurrentUserInput struct{}

type GetCurrentUserOutput struct {
	User *tracker.User `json:"user"`
}

func (t *toolServer) GetCurrentUser(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	_ GetCurrentUserInput,
) (*mcp.CallToolResult, GetCurrentUserOutput, error) {
	user, err := t.Tracker.GetCurrentUser(ctx)
	if err != nil {
		return nil, GetCurrentUserOutput{}, err
	}
	return nil, GetCurrentUserOutput{User: user}, nil
}

type GetUserInput struct {
	UID string `json:"uid" jsonschema:"User UID or login"`
}

type GetUserOutput struct {
	User *tracker.User `json:"user"`
}

func (t *toolServer) GetUser(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input GetUserInput,
) (*mcp.CallToolResult, GetUserOutput, error) {
	user, err := t.Tracker.GetUser(ctx, input.UID)
	if err != nil {
		return nil, GetUserOutput{}, err
	}
	return nil, GetUserOutput{User: user}, nil
}
