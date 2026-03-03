package tracker

import (
	"context"
	"fmt"
	"net/http"
)

func (c *client) GetCurrentUser(ctx context.Context) (*User, error) {
	var result User
	if err := c.do(ctx, http.MethodGet, baseURL+"/myself", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *client) GetUser(ctx context.Context, uid string) (*User, error) {
	var result User
	if err := c.do(ctx, http.MethodGet, fmt.Sprintf("%s/users/%s", baseURL, uid), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
