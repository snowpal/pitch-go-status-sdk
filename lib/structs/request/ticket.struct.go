package request

type AddSessionsReqBody struct {
	Sessions []Session `json:"sessions"`
}

type Session struct {
	Tickets    []TicketReqBody           `json:"tickets"`
	OtherItems []SessionOtherItemReqBody `json:"otherItems"`
}

type StatusSessionOtherItemsReqBody struct {
	SessionOtherItems []SessionOtherItemReqBody `json:"otherItems"`
}

type SessionOtherItemReqBody struct {
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	TimeSpent     float32        `json:"timeSpent"`
	PairedMembers []PairedMember `json:"pairedWith"`
	Comment       string         `json:"comment"`
}

type AddTicketsReqBody struct {
	Tickets []TicketReqBody `json:"tickets"`
}

type TicketReqBody struct {
	ID            string         `json:"ticketID"`
	Type          string         `json:"type"`
	Url           string         `json:"url"`
	Description   string         `json:"description"`
	Status        string         `json:"status"`
	Pionts        int            `json:"points"`
	TimeSpent     float32        `json:"timeSpent"`
	PullRequest   PullRequest    `json:"pullRequest"`
	PairedMembers []PairedMember `json:"pairedWith"`
}

type PullRequest struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}

type PairedMember struct {
	MemberId      string `json:"memberID"`
	NonMemberName string `json:"nonMemberName"`
}
