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

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_issue_comments",
		Description: "Get all comments for a Yandex Tracker issue",
	}, tools.GetIssueComments)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "add_comment",
		Description: "Add a comment to a Yandex Tracker issue",
	}, tools.AddComment)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_current_user",
		Description: "Get information about the authenticated Yandex Tracker user",
	}, tools.GetCurrentUser)

	return server
}

type toolServer struct {
	Tracker tracker.Client
}
