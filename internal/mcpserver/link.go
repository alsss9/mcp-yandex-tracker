package mcpserver

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sunnyyssh/mcp-yandex-tracker/internal/tracker"
)

type GetIssueLinksInput struct {
	Key string `json:"key" jsonschema:"Issue key, e.g. PROJ-123"`
}

type GetIssueLinksOutput struct {
	Links []*tracker.Link `json:"links"`
}

func (t *toolServer) GetIssueLinks(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input GetIssueLinksInput,
) (*mcp.CallToolResult, GetIssueLinksOutput, error) {
	links, err := t.Tracker.GetIssueLinks(ctx, input.Key)
	if err != nil {
		return nil, GetIssueLinksOutput{}, err
	}
	return nil, GetIssueLinksOutput{Links: links}, nil
}

type CreateIssueLinkInput struct {
	Key          string `json:"key"          jsonschema:"Issue key, e.g. PROJ-123"`
	Relationship string `json:"relationship" jsonschema:"Relationship type: relates, depends, blocks, duplicates, epic, subtask"`
	Issue        string `json:"issue"        jsonschema:"Key of the issue to link to"`
}

type CreateIssueLinkOutput struct {
	Link *tracker.Link `json:"link"`
}

func (t *toolServer) CreateIssueLink(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input CreateIssueLinkInput,
) (*mcp.CallToolResult, CreateIssueLinkOutput, error) {
	link, err := t.Tracker.CreateIssueLink(ctx, input.Key, tracker.CreateLinkOptions{
		Relationship: input.Relationship,
		Issue:        input.Issue,
	})
	if err != nil {
		return nil, CreateIssueLinkOutput{}, err
	}
	return nil, CreateIssueLinkOutput{Link: link}, nil
}

type DeleteIssueLinkInput struct {
	Key    string `json:"key"     jsonschema:"Issue key, e.g. PROJ-123"`
	LinkID string `json:"link_id" jsonschema:"Link ID"`
}

type DeleteIssueLinkOutput struct {
	OK bool `json:"ok"`
}

func (t *toolServer) DeleteIssueLink(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input DeleteIssueLinkInput,
) (*mcp.CallToolResult, DeleteIssueLinkOutput, error) {
	if err := t.Tracker.DeleteIssueLink(ctx, input.Key, input.LinkID); err != nil {
		return nil, DeleteIssueLinkOutput{}, err
	}
	return nil, DeleteIssueLinkOutput{OK: true}, nil
}
