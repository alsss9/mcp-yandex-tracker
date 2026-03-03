package tracker

import (
	"context"
	"fmt"
	"net/http"
)

func (c *client) GetIssueLinks(ctx context.Context, key string) ([]*Link, error) {
	var result []*Link
	if err := c.do(ctx, http.MethodGet, fmt.Sprintf("%s/issues/%s/links", baseURL, key), nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) CreateIssueLink(ctx context.Context, key string, opts CreateLinkOptions) (*Link, error) {
	var result Link
	if err := c.do(ctx, http.MethodPost, fmt.Sprintf("%s/issues/%s/links", baseURL, key), opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) DeleteIssueLink(ctx context.Context, key, linkID string) error {
	return c.do(ctx, http.MethodDelete, fmt.Sprintf("%s/issues/%s/links/%s", baseURL, key, linkID), nil, nil)
}
