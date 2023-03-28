package response

type PullRequests struct {
	PullRequests []PullRequest `json:"pullRequests"`
}

type PullRequest struct {
	ID     string  `json:"id"`
	Url    string  `json:"url"`
	Team   *Team   `json:"team"`
	Member *Member `json:"member"`
}
