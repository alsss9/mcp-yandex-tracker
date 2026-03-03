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

type UpdateCommentInput struct {
	Key       string `json:"key"        jsonschema:"Issue key, e.g. PROJ-123"`
	CommentID string `json:"comment_id" jsonschema:"Comment ID"`
	Text      string `json:"text"       jsonschema:"New comment text"`
}

type UpdateCommentOutput struct {
	Comment *tracker.Comment `json:"comment"`
}

func (t *toolServer) UpdateComment(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input UpdateCommentInput,
) (*mcp.CallToolResult, UpdateCommentOutput, error) {
	comment, err := t.Tracker.UpdateComment(ctx, input.Key, input.CommentID, input.Text)
	if err != nil {
		return nil, UpdateCommentOutput{}, err
	}
	return nil, UpdateCommentOutput{Comment: comment}, nil
}

type DeleteCommentInput struct {
	Key       string `json:"key"        jsonschema:"Issue key, e.g. PROJ-123"`
	CommentID string `json:"comment_id" jsonschema:"Comment ID"`
}

type DeleteCommentOutput struct {
	OK bool `json:"ok"`
}

func (t *toolServer) DeleteComment(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input DeleteCommentInput,
) (*mcp.CallToolResult, DeleteCommentOutput, error) {
	if err := t.Tracker.DeleteComment(ctx, input.Key, input.CommentID); err != nil {
		return nil, DeleteCommentOutput{}, err
	}
	return nil, DeleteCommentOutput{OK: true}, nil
}
