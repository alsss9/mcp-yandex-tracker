package tracker

import (
	"context"
	"fmt"
	"net/http"
)

func (c *client) GetIssue(ctx context.Context, key string) (*Issue, error) {
	var result Issue
	if err := c.do(ctx, http.MethodGet, fmt.Sprintf("%s/issues/%s", baseURL, key), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) CreateIssue(ctx context.Context, opts CreateIssueOptions) (*Issue, error) {
	var result Issue
	if err := c.do(ctx, http.MethodPost, baseURL+"/issues/", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) UpdateIssue(ctx context.Context, key string, opts UpdateIssueOptions) (*Issue, error) {
	var result Issue
	if err := c.do(ctx, http.MethodPatch, fmt.Sprintf("%s/issues/%s", baseURL, key), opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) FindIssues(ctx context.Context, opts FindIssuesOptions) ([]*Issue, error) {
	var result []*Issue
	if err := c.do(ctx, http.MethodPost, baseURL+"/issues/_search", opts, &result); err != nil {
		return nil, err
	}
	return result, nil
}
