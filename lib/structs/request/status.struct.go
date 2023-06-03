package request

import "github.com/snowpal/go-status-sdk/lib/structs/common"

type StatusReqBody struct {
	MemberId     string              `json:"memberID"`
	StatusDate   string              `json:"statusDate"`
	PlanForToday common.PlanForToday `json:"planForToday"`
	Sessions     []common.Session    `json:"sessions"`
	BlockedBy    common.BlockedBy    `json:"blockedBy"`
}

type StatusOtherItemsReqBody struct {
	StatusOtherItems []string `json:"otherItems"`
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

type TicketParam struct {
	TeamId    string
	MemberId  string
	StatusId  string
	SessionId string
	TicketId  string
}

type OtherItemParam struct {
	TeamId      string
	MemberId    string
	StatusId    string
	SessionId   string
	OtherItemId string
}
