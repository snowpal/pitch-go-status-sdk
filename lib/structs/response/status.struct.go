package response

type Statuses struct {
	Statuses []Status `json:"statuses"`
}

type Status struct {
	ID             string       `json:"ID"`
	TeamID         string       `json:"TeamID"`
	MemberId       string       `json:"MemberID"`
	StatusDate     string       `json:"StatusDate"`
	PlanForToday   PlanForToday `json:"PlanForToday"`
	Sessions       []Session    `json:"Sessions"`
	BlockedBy      BlockedBy    `json:"BlockedBy"`
	Active         bool         `json:"Active"`
	CreatedBy      string       `json:"CreatedBy"`
	UpdatedBy      string       `json:"UpdatedBy"`
	LastModifiedOn string       `json:"LastModifiedOn"`
}

type Sessions struct {
	Sessions []Session `json:"sessions"`
}

type Session struct {
	ID         string             `json:"ID"`
	SequenceId int                `json:"SequenceID"`
	Tickets    []SessionTicket    `json:"Tickets"`
	OtherItems []SessionOtherItem `json:"OtherItems"`
}

type PlanForToday struct {
	Tickets    []Ticket `json:"Tickets"`
	OtherItems []string `json:"OtherItems"`
}

type PairedMember struct {
	MemberId      string   `json:"MemberID"`
	NonMemberName string   `json:"NonMemberName"`
	Items         []string `json:"Items"`
}

type BlockedBy struct {
	Tickets    []Ticket       `json:"Tickets"`
	OtherItems []string       `json:"OtherItems"`
	Members    []PairedMember `json:"Members"`
}

type StatusOtherItems struct {
	OtherItems []string `json:"OtherItems"`
}
