package tracker

// BasicRef is a lightweight reference used in nested objects (user, queue, status, etc.).
type BasicRef struct {
	Self    string `json:"self"`
	ID      string `json:"id"`
	Key     string `json:"key"`
	Display string `json:"display"`
}

// User is the full user object returned by /myself and user endpoints.
type User struct {
	Self                 string `json:"self"`
	UID                  int    `json:"uid"`
	Login                string `json:"login"`
	TrackerUID           int    `json:"trackerUid"`
	PassportUID          int    `json:"passportUid"`
	CloudUID             string `json:"cloudUid"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	Display              string `json:"display"`
	Email                string `json:"email"`
	External             bool   `json:"external"`
	HasLicense           bool   `json:"hasLicense"`
	Dismissed            bool   `json:"dismissed"`
	DisableNotifications bool   `json:"disableNotifications"`
	FirstLoginDate       string `json:"firstLoginDate"`
	LastLoginDate        string `json:"lastLoginDate"`
}

// Issue is the full issue object returned by the Tracker API.
type Issue struct {
	Self                 string     `json:"self"`
	ID                   string     `json:"id"`
	Key                  string     `json:"key"`
	Version              int        `json:"version"`
	Summary              string     `json:"summary"`
	Description          string     `json:"description"`
	CreatedAt            string     `json:"createdAt"`
	UpdatedAt            string     `json:"updatedAt"`
	LastCommentUpdatedAt string     `json:"lastCommentUpdatedAt"`
	Aliases              []string   `json:"aliases"`
	Votes                int        `json:"votes"`
	Favorite             bool       `json:"favorite"`
	Parent               *BasicRef  `json:"parent"`
	Queue                *BasicRef  `json:"queue"`
	Status               *BasicRef  `json:"status"`
	PreviousStatus       *BasicRef  `json:"previousStatus"`
	Type                 *BasicRef  `json:"type"`
	Priority             *BasicRef  `json:"priority"`
	CreatedBy            *BasicRef  `json:"createdBy"`
	UpdatedBy            *BasicRef  `json:"updatedBy"`
	Assignee             *BasicRef  `json:"assignee"`
	Followers            []BasicRef `json:"followers"`
	Sprint               []BasicRef `json:"sprint"`
}

// CreateIssueOptions is the request body for creating an issue.
type CreateIssueOptions struct {
	Summary     string `json:"summary"`
	Queue       string `json:"queue"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Priority    string `json:"priority,omitempty"`
	Assignee    string `json:"assignee,omitempty"`
	Parent      string `json:"parent,omitempty"`
}

// UpdateIssueOptions is the request body for patching an issue.
type UpdateIssueOptions struct {
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
	Priority    string `json:"priority,omitempty"`
	Assignee    string `json:"assignee,omitempty"`
}

// FindIssuesOptions is the request body for POST /issues/_search.
type FindIssuesOptions struct {
	Queue  string         `json:"queue,omitempty"`
	Keys   []string       `json:"keys,omitempty"`
	Filter map[string]any `json:"filter,omitempty"`
	Query  string         `json:"query,omitempty"`
}

// Comment is an issue comment object.
type Comment struct {
	Self      string    `json:"self"`
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	CreatedBy *BasicRef `json:"createdBy"`
	UpdatedBy *BasicRef `json:"updatedBy"`
}
