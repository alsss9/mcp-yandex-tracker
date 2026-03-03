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

// CreateIssueOptions is the request body for POST /issues/.
type CreateIssueOptions struct {
	Summary     string   `json:"summary"`
	Queue       string   `json:"queue"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"type,omitempty"`
	Priority    string   `json:"priority,omitempty"`
	Assignee    string   `json:"assignee,omitempty"`
	Parent      string   `json:"parent,omitempty"`
	Sprint      []string `json:"sprint,omitempty"`
}

// UpdateIssueOptions is the request body for PATCH /issues/{key}.
type UpdateIssueOptions struct {
	Summary     string    `json:"summary,omitempty"`
	Description string    `json:"description,omitempty"`
	Priority    string    `json:"priority,omitempty"`
	Assignee    string    `json:"assignee,omitempty"`
	Sprint      *[]string `json:"sprint,omitempty"`
}

// FindIssuesOptions is the request body for POST /issues/_search.
type FindIssuesOptions struct {
	Queue  string         `json:"queue,omitempty"`
	Keys   []string       `json:"keys,omitempty"`
	Filter map[string]any `json:"filter,omitempty"`
	Query  string         `json:"query,omitempty"`
}

// Transition is an available workflow transition for an issue.
type Transition struct {
	Self    string    `json:"self"`
	ID      string    `json:"id"`
	Display string    `json:"display"`
	To      *BasicRef `json:"to"`
}

// ExecuteTransitionOptions is the optional body for POST /issues/{key}/transitions/{id}/_execute.
type ExecuteTransitionOptions struct {
	// Comment to add when executing the transition.
	Comment string `json:"comment,omitempty"`
	// Resolution to set (key or id), used when closing an issue.
	Resolution string `json:"resolution,omitempty"`
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

// Link represents a relation between two issues.
type Link struct {
	Self          string    `json:"self"`
	ID            string    `json:"id"`
	Type          *LinkType `json:"type"`
	Direction     string    `json:"direction"`
	Object        *BasicRef `json:"object"`
	CreatedBy     *BasicRef `json:"createdBy"`
	UpdatedBy     *BasicRef `json:"updatedBy"`
	CreatedAt     string    `json:"createdAt"`
	UpdatedAt     string    `json:"updatedAt"`
}

// LinkType describes the relationship type between issues.
type LinkType struct {
	Self      string `json:"self"`
	ID        string `json:"id"`
	Inward    string `json:"inward"`
	Outward   string `json:"outward"`
}

// CreateLinkOptions is the request body for POST /issues/{key}/links.
type CreateLinkOptions struct {
	// Relationship type key, e.g. "relates", "depends", "blocks", "duplicates", "epic", "subtask".
	Relationship string `json:"relationship"`
	// Key of the issue to link to.
	Issue string `json:"issue"`
}

// Queue is a Yandex Tracker queue.
type Queue struct {
	Self        string    `json:"self"`
	ID          int       `json:"id"`
	Key         string    `json:"key"`
	Version     int       `json:"version"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Lead        *BasicRef `json:"lead"`
	Assignee    *BasicRef `json:"assignee"`
}
