package mcpserver

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sunnyyssh/mcp-yandex-tracker/internal/tracker"
)

type GetIssueTransitionsInput struct {
	Key string `json:"key" jsonschema:"Issue key, e.g. PROJ-123"`
}

type GetIssueTransitionsOutput struct {
	Transitions []*tracker.Transition `json:"transitions"`
}

func (t *toolServer) GetIssueTransitions(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input GetIssueTransitionsInput,
) (*mcp.CallToolResult, GetIssueTransitionsOutput, error) {
	transitions, err := t.Tracker.GetIssueTransitions(ctx, input.Key)
	if err != nil {
		return nil, GetIssueTransitionsOutput{}, err
	}
	return nil, GetIssueTransitionsOutput{Transitions: transitions}, nil
}

type ExecuteTransitionInput struct {
	Key          string `json:"key"                    jsonschema:"Issue key, e.g. PROJ-123"`
	TransitionID string `json:"transition_id"          jsonschema:"Transition ID from get_issue_transitions"`
	Comment      string `json:"comment,omitempty"      jsonschema:"Optional comment to add on transition"`
	Resolution   string `json:"resolution,omitempty"   jsonschema:"Resolution key when closing an issue"`
}

type ExecuteTransitionOutput struct {
	Issue *tracker.Issue `json:"issue"`
}

func (t *toolServer) ExecuteTransition(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input ExecuteTransitionInput,
) (*mcp.CallToolResult, ExecuteTransitionOutput, error) {
	issue, err := t.Tracker.ExecuteTransition(ctx, input.Key, input.TransitionID, tracker.ExecuteTransitionOptions{
		Comment:    input.Comment,
		Resolution: input.Resolution,
	})
	if err != nil {
		return nil, ExecuteTransitionOutput{}, err
	}
	return nil, ExecuteTransitionOutput{Issue: issue}, nil
}
