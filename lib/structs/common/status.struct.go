package common

type PlanForToday struct {
	Tickets    []TicketReqBody `json:"tickets"`
	OtherItems []string        `json:"otherItems"`
}

type BlockedBy struct {
	Tickets    []TicketReqBody `json:"tickets"`
	OtherItems []string        `json:"otherItems"`
	Members    []PairedMember  `json:"members"`
}

type PairedMember struct {
	MemberId      string   `json:"memberID"`
	NonMemberName string   `json:"nonMemberName"`
	Items         []string `json:"items"`
}

type Session struct {
	SequenceId int                       `json:"sequenceId"`
	Tickets    []TicketReqBody           `json:"tickets"`
	OtherItems []SessionOtherItemReqBody `json:"otherItems"`
}

type TicketReqBody struct {
	ID          string `json:"ticketID"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Points      int    `json:"points"`
}

type SessionOtherItemReqBody struct {
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	TimeSpent     float32        `json:"timeSpent"`
	PairedMembers []PairedMember `json:"pairedWith"`
	Comment       string         `json:"comment"`
}
