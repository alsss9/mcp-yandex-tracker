package mcpserver

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sunnyyssh/mcp-yandex-tracker/internal/tracker"
)

type GetQueuesInput struct{}

type GetQueuesOutput struct {
	Queues []*tracker.Queue `json:"queues"`
}

func (t *toolServer) GetQueues(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	_ GetQueuesInput,
) (*mcp.CallToolResult, GetQueuesOutput, error) {
	queues, err := t.Tracker.GetQueues(ctx)
	if err != nil {
		return nil, GetQueuesOutput{}, err
	}
	return nil, GetQueuesOutput{Queues: queues}, nil
}

type GetQueueInput struct {
	QueueID string `json:"queue_id" jsonschema:"Queue key or ID, e.g. PROJ"`
}

type GetQueueOutput struct {
	Queue *tracker.Queue `json:"queue"`
}

func (t *toolServer) GetQueue(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input GetQueueInput,
) (*mcp.CallToolResult, GetQueueOutput, error) {
	queue, err := t.Tracker.GetQueue(ctx, input.QueueID)
	if err != nil {
		return nil, GetQueueOutput{}, err
	}
	return nil, GetQueueOutput{Queue: queue}, nil
}
