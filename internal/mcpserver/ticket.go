package mcpserver

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sunnyyssh/mcp-yandex-tracker/internal/tracker"
)

type FindIssuesInput struct {
	Queue  string            `json:"queue,omitempty"  jsonschema:"Queue key to search in"`
	Query  string            `json:"query,omitempty"  jsonschema:"Tracker query language expression"`
	Filter map[string]string `json:"filter,omitempty" jsonschema:"Key-value field filters"`
}

type FindIssuesOutput struct {
	Issues []*tracker.Issue `json:"issues"`
}

func (t *toolServer) FindIssues(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input FindIssuesInput,
) (*mcp.CallToolResult, FindIssuesOutput, error) {
	filter := make(map[string]interface{}, len(input.Filter))
	for k, v := range input.Filter {
		filter[k] = v
	}
	issues, err := t.Tracker.FindIssues(ctx, tracker.FindIssuesOptions{
		Queue:  input.Queue,
		Query:  input.Query,
		Filter: filter,
	})
	if err != nil {
		return nil, FindIssuesOutput{}, err
	}
	return nil, FindIssuesOutput{Issues: issues}, nil
}

type GetIssueInput struct {
	Key string `json:"key" jsonschema:"Issue key, e.g. PROJ-123"`
}

type GetIssueOutput struct {
	Issue *tracker.Issue `json:"issue"`
}

func (t *toolServer) GetIssue(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input GetIssueInput,
) (*mcp.CallToolResult, GetIssueOutput, error) {
	issue, err := t.Tracker.GetIssue(ctx, input.Key)
	if err != nil {
		return nil, GetIssueOutput{}, err
	}
	return nil, GetIssueOutput{Issue: issue}, nil
}

type CreateIssueInput struct {
	Summary     string `json:"summary"               jsonschema:"Issue title"`
	Queue       string `json:"queue"                 jsonschema:"Queue key, e.g. PROJ"`
	Description string `json:"description,omitempty" jsonschema:"Issue description"`
	Type        string `json:"type,omitempty"        jsonschema:"Issue type key"`
	Priority    string `json:"priority,omitempty"    jsonschema:"Priority key"`
	Assignee    string `json:"assignee,omitempty"    jsonschema:"Assignee login"`
	Parent      string `json:"parent,omitempty"      jsonschema:"Parent issue key"`
}

type CreateIssueOutput struct {
	Issue *tracker.Issue `json:"issue"`
}

func (t *toolServer) CreateIssue(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input CreateIssueInput,
) (*mcp.CallToolResult, CreateIssueOutput, error) {
	issue, err := t.Tracker.CreateIssue(ctx, tracker.CreateIssueOptions{
		Summary:     input.Summary,
		Queue:       input.Queue,
		Description: input.Description,
		Type:        input.Type,
		Priority:    input.Priority,
		Assignee:    input.Assignee,
		Parent:      input.Parent,
	})
	if err != nil {
		return nil, CreateIssueOutput{}, err
	}
	return nil, CreateIssueOutput{Issue: issue}, nil
}

type UpdateIssueInput struct {
	Key         string `json:"key"                   jsonschema:"Issue key, e.g. PROJ-123"`
	Summary     string `json:"summary,omitempty"     jsonschema:"New issue title"`
	Description string `json:"description,omitempty" jsonschema:"New issue description"`
	Priority    string `json:"priority,omitempty"    jsonschema:"New priority key"`
	Assignee    string `json:"assignee,omitempty"    jsonschema:"New assignee login"`
}

type UpdateIssueOutput struct {
	Issue *tracker.Issue `json:"issue"`
}

func (t *toolServer) UpdateIssue(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input UpdateIssueInput,
) (*mcp.CallToolResult, UpdateIssueOutput, error) {
	issue, err := t.Tracker.UpdateIssue(ctx, input.Key, tracker.UpdateIssueOptions{
		Summary:     input.Summary,
		Description: input.Description,
		Priority:    input.Priority,
		Assignee:    input.Assignee,
	})
	if err != nil {
		return nil, UpdateIssueOutput{}, err
	}
	return nil, UpdateIssueOutput{Issue: issue}, nil
}
