package tracker

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

const baseURL = "https://api.tracker.yandex.net/v3"

// Config holds authentication and organization settings for the Tracker API.
type Config struct {
	// Token is the OAuth or IAM token.
	// OAuth: "OAuth <token>", IAM: "Bearer <token>"
	Token string

	// OrgID is the Yandex 360 organization ID (X-Org-ID header).
	OrgID string

	// CloudOrgID is the Yandex Cloud organization ID (X-Cloud-Org-ID header).
	// Takes priority over OrgID when set.
	CloudOrgID string
}

// Client is the interface for interacting with the Yandex Tracker API.
type Client interface {
	// Issues
	GetIssue(ctx context.Context, key string) (*Issue, error)
	CreateIssue(ctx context.Context, opts CreateIssueOptions) (*Issue, error)
	UpdateIssue(ctx context.Context, key string, opts UpdateIssueOptions) (*Issue, error)
	FindIssues(ctx context.Context, opts FindIssuesOptions) ([]*Issue, error)

	// Transitions
	GetIssueTransitions(ctx context.Context, key string) ([]*Transition, error)
	ExecuteTransition(ctx context.Context, key, transitionID string, opts ExecuteTransitionOptions) (*Issue, error)

	// Comments
	GetIssueComments(ctx context.Context, key string) ([]*Comment, error)
	AddComment(ctx context.Context, key, text string) (*Comment, error)
	UpdateComment(ctx context.Context, key, commentID, text string) (*Comment, error)
	DeleteComment(ctx context.Context, key, commentID string) error

	// Links
	GetIssueLinks(ctx context.Context, key string) ([]*Link, error)
	CreateIssueLink(ctx context.Context, key string, opts CreateLinkOptions) (*Link, error)
	DeleteIssueLink(ctx context.Context, key, linkID string) error

	// Queues
	GetQueues(ctx context.Context) ([]*Queue, error)
	GetQueue(ctx context.Context, queueID string) (*Queue, error)

	// Users
	GetCurrentUser(ctx context.Context) (*User, error)
	GetUser(ctx context.Context, uid string) (*User, error)
}

type client struct {
	http    *resty.Client
	headers map[string]string
}

func authHeader(token string) string {
	if strings.HasPrefix(token, "OAuth ") || strings.HasPrefix(token, "Bearer ") {
		return token
	}
	return "OAuth " + token
}

func New(cfg Config) Client {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": authHeader(cfg.Token),
	}
	if cfg.CloudOrgID != "" {
		headers["X-Cloud-Org-ID"] = cfg.CloudOrgID
	} else {
		headers["X-Org-ID"] = cfg.OrgID
	}
	return &client{
		http:    resty.New(),
		headers: headers,
	}
}

func (c *client) do(ctx context.Context, method, url string, body, result any) error {
	req := c.http.R().SetContext(ctx).SetHeaders(c.headers)
	if body != nil {
		req = req.SetBody(body)
	}

	var (
		resp *resty.Response
		err  error
	)
	resp, err = req.Execute(method, url)
	if err != nil {
		return fmt.Errorf("request %s %s: %w", method, url, err)
	}

	code := resp.StatusCode()
	if code < 200 || code >= 300 {
		return fmt.Errorf("status %d from %s %s: %s", code, method, url, resp.Body())
	}

	if result != nil {
		if err := json.Unmarshal(resp.Body(), result); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}
	return nil
}
