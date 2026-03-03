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
