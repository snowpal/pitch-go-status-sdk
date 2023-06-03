package response

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

type Statuses struct {
	Statuses []Status `json:"statuses"`
}

type Status struct {
	ID           string       `json:"id"`
	MemberId     string       `json:"memberID"`
	StatusDate   string       `json:"statusDate"`
	PlanForToday PlanForToday `json:"planForToday"`
	Sessions     []Session    `json:"sessions"`
	BlockedBy    BlockedBy    `json:"blockedBy"`
}

type Sessions struct {
	Sessions []Session `json:"sessions"`
}

type Session struct {
	ID         string             `json:"id"`
	SequenceId int                `json:"sequenceId"`
	Tickets    []Ticket           `json:"tickets"`
	OtherItems []SessionOtherItem `json:"otherItems"`
}

type PlanForToday struct {
	Tickets    []Ticket `json:"tickets"`
	OtherItems []string `json:"otherItems"`
}

type BlockedBy struct {
	Tickets    []Ticket              `json:"tickets"`
	OtherItems []string              `json:"otherItems"`
	Members    []common.PairedMember `json:"members"`
}

type StatusOtherItems struct {
	OtherItems []string `json:"otherItems"`
}
