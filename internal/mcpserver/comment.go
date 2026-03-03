package mcpserver

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sunnyyssh/mcp-yandex-tracker/internal/tracker"
)

type GetIssueCommentsInput struct {
	Key string `json:"key" jsonschema:"Issue key, e.g. PROJ-123"`
}

type GetIssueCommentsOutput struct {
	Comments []*tracker.Comment `json:"comments"`
}

func (t *toolServer) GetIssueComments(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input GetIssueCommentsInput,
) (*mcp.CallToolResult, GetIssueCommentsOutput, error) {
	comments, err := t.Tracker.GetIssueComments(ctx, input.Key)
	if err != nil {
		return nil, GetIssueCommentsOutput{}, err
	}
	return nil, GetIssueCommentsOutput{Comments: comments}, nil
}

type AddCommentInput struct {
	Key  string `json:"key"  jsonschema:"Issue key, e.g. PROJ-123"`
	Text string `json:"text" jsonschema:"Comment text"`
}

type AddCommentOutput struct {
	Comment *tracker.Comment `json:"comment"`
}

func (t *toolServer) AddComment(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input AddCommentInput,
) (*mcp.CallToolResult, AddCommentOutput, error) {
	comment, err := t.Tracker.AddComment(ctx, input.Key, input.Text)
	if err != nil {
		return nil, AddCommentOutput{}, err
	}
	return nil, AddCommentOutput{Comment: comment}, nil
}
