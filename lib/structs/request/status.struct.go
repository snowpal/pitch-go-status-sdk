package request

import "github.com/snowpal/go-status-sdk/lib/structs/common"

type StatusReqBody struct {
	MemberId     string           `json:"memberID"`
	StatusDate   string           `json:"statusDate"`
	PlanForToday PlanForToday     `json:"planForToday"`
	Sessions     []SessionReqBody `json:"sessions"`
	BlockedBy    BlockedBy        `json:"blockedBy"`
}

type PlanForToday struct {
	Tickets    []TicketReqBody `json:"tickets"`
	OtherItems []string        `json:"otherItems"`
}

type BlockedBy struct {
	Tickets    []TicketReqBody       `json:"tickets"`
	OtherItems []string              `json:"otherItems"`
	Members    []common.PairedMember `json:"members"`
}

type StatusOtherItemsReqBody struct {
	OtherItems []string `json:"otherItems"`
}

type TeamParam struct {
	TeamId   string
	MemberId string
}

type StatusParam struct {
	TeamId    string
	MemberId  string
	StatusId  string
	StartDate string
	EndDate   string
}

type SessionParam struct {
	TeamId    string
	MemberId  string
	StatusId  string
	SessionId string
}
