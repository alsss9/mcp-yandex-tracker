package tracker

import (
	"context"
	"net/http"
)

func (c *client) GetCurrentUser(ctx context.Context) (*User, error) {
	var result User
	if err := c.do(ctx, http.MethodGet, baseURL+"/myself", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
