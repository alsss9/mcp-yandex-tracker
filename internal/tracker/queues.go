package tracker

import (
	"context"
	"fmt"
	"net/http"
)

func (c *client) GetQueues(ctx context.Context) ([]*Queue, error) {
	var result []*Queue
	if err := c.do(ctx, http.MethodGet, baseURL+"/queues", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetQueue(ctx context.Context, queueID string) (*Queue, error) {
	var result Queue
	if err := c.do(ctx, http.MethodGet, fmt.Sprintf("%s/queues/%s", baseURL, queueID), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
