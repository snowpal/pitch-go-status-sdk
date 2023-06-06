package common

type Member struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

type PairedMember struct {
	MemberId      string   `json:"memberID"`
	NonMemberName string   `json:"nonMemberName"`
	Items         []string `json:"items"`
}

type PullRequest struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}
