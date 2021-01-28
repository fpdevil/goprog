// Package github provides a Go API for the GitHub issue tracker.
// The main metadata configuration containing all the GitHub url's
// are available here: https://api.github.com
// The format for issue tracker url is:
// https://api.github.com/search/issues?q={query}{&page,per_page,sort,order}
//
package github

import "time"

// IssuesURL is the main url for issue tracker API
const IssuesURL = "https://api.github.com/search/issues"

// SearchResults is the main struct type containing top level fields
type SearchResults struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue struct contains the child elements of Items issues
type Issue struct {
	HTMLURL           string `json:"html_url"`
	Number            int
	Title             string
	State             string
	CreatedAt         time.Time `json:"created_at"`
	AuthorAssociation string
	Body              string
	Draft             bool
	User              *User
	Score             float64
}

// User struct contains user level details
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
	ID      int
}
