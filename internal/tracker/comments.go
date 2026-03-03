package tracker

import (
	"context"
	"fmt"
	"net/http"
)

func (c *client) GetIssueComments(ctx context.Context, key string) ([]*Comment, error) {
	var result []*Comment
	if err := c.do(ctx, http.MethodGet, fmt.Sprintf("%s/issues/%s/comments", baseURL, key), nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) AddComment(ctx context.Context, key, text string) (*Comment, error) {
	body := map[string]string{"text": text}
	var result Comment
	if err := c.do(ctx, http.MethodPost, fmt.Sprintf("%s/issues/%s/comments", baseURL, key), body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) UpdateComment(ctx context.Context, key, commentID, text string) (*Comment, error) {
	body := map[string]string{"text": text}
	var result Comment
	url := fmt.Sprintf("%s/issues/%s/comments/%s", baseURL, key, commentID)
	if err := c.do(ctx, http.MethodPatch, url, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) DeleteComment(ctx context.Context, key, commentID string) error {
	url := fmt.Sprintf("%s/issues/%s/comments/%s", baseURL, key, commentID)
	return c.do(ctx, http.MethodDelete, url, nil, nil)
}
