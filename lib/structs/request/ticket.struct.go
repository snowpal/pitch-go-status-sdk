package request

import "github.com/snowpal/go-status-sdk/lib/structs/common"

type AddSessionsReqBody struct {
	Sessions []SessionReqBody `json:"sessions"`
}

type StatusSessionOtherItemsReqBody struct {
	OtherItems []SessionOtherItemReqBody `json:"otherItems"`
}

type AddSessionTicketsReqBody struct {
	Tickets []SessionTicketReqBody `json:"tickets"`
}

type SessionTicketReqBody struct {
	TicketID      string                `json:"ticketID"`
	Type          string                `json:"type"`
	Url           string                `json:"url"`
	Description   string                `json:"description"`
	Status        string                `json:"status"`
	TimeSpent     float32               `json:"timeSpent"`
	PullRequest   common.PullRequest    `json:"pullRequest"`
	PairedMembers []common.PairedMember `json:"pairedWith"`
}

type AddTicketsReqBody struct {
	Tickets []TicketReqBody `json:"tickets"`
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
