package domain

type PullRequest struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	User    User   `json:"user"`
	Head    Branch `json:"head"`
	Base    Branch `json:"base"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}