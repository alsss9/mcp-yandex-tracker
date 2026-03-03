package mcpserver

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/sunnyyssh/mcp-yandex-tracker/internal/tracker"
)

func NewServer(trackerClient tracker.Client) *mcp.Server {
	tools := &toolServer{Tracker: trackerClient}

	server := mcp.NewServer(&mcp.Implementation{
		Name:       "yandex-tracker",
		Title:      "Yandex Tracker",
		Version:    "v1.0.0",
		WebsiteURL: "https://github.com/sunnyyssh/mcp-yandex-tracker",
	}, nil)

	// Issues
	mcp.AddTool(server, &mcp.Tool{
		Name:        "find_issues",
		Description: "Search for issues in Yandex Tracker by queue, query, or field filters",
	}, tools.FindIssues)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_issue",
		Description: "Get a single Yandex Tracker issue by its key",
	}, tools.GetIssue)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_issue",
		Description: "Create a new issue in Yandex Tracker",
	}, tools.CreateIssue)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_issue",
		Description: "Update fields of an existing Yandex Tracker issue",
	}, tools.UpdateIssue)

	// Transitions
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_issue_transitions",
		Description: "Get available workflow transitions for a Yandex Tracker issue",
	}, tools.GetIssueTransitions)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "execute_transition",
		Description: "Execute a workflow transition on a Yandex Tracker issue (e.g. move to In Progress, close)",
	}, tools.ExecuteTransition)

	// Comments
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_issue_comments",
		Description: "Get all comments for a Yandex Tracker issue",
	}, tools.GetIssueComments)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "add_comment",
		Description: "Add a comment to a Yandex Tracker issue",
	}, tools.AddComment)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_comment",
		Description: "Update an existing comment on a Yandex Tracker issue",
	}, tools.UpdateComment)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_comment",
		Description: "Delete a comment from a Yandex Tracker issue",
	}, tools.DeleteComment)

	// Links
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_issue_links",
		Description: "Get all links (relations) for a Yandex Tracker issue",
	}, tools.GetIssueLinks)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_issue_link",
		Description: "Create a link between two Yandex Tracker issues",
	}, tools.CreateIssueLink)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_issue_link",
		Description: "Delete a link between Yandex Tracker issues",
	}, tools.DeleteIssueLink)

	// Queues
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_queues",
		Description: "List all accessible Yandex Tracker queues",
	}, tools.GetQueues)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_queue",
		Description: "Get a single Yandex Tracker queue by its key or ID",
	}, tools.GetQueue)

	// Users
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_current_user",
		Description: "Get information about the authenticated Yandex Tracker user",
	}, tools.GetCurrentUser)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_user",
		Description: "Get information about a Yandex Tracker user by UID or login",
	}, tools.GetUser)

	return server
}

type toolServer struct {
	Tracker tracker.Client
}
