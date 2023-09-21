package request

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

type AddSessionsReqBody struct {
	Sessions []SessionReqBody `json:"sessions"`
}

type SessionReqBody struct {
	SequenceId int                       `json:"sequenceId"`
	Tickets    []SessionTicketReqBody    `json:"tickets"`
	OtherItems []SessionOtherItemReqBody `json:"otherItems"`
}

type StatusSessionOtherItemsReqBody struct {
	OtherItems []SessionOtherItemReqBody `json:"otherItems"`
}

type SessionOtherItemReqBody struct {
	Title         string                `json:"title"`
	Type          string                `json:"type"`
	TimeSpent     float32               `json:"timeSpent"`
	PairedMembers []common.PairedMember `json:"pairedWith"`
	Comment       string                `json:"comment"`
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

type TicketReqBody struct {
	TicketID    string `json:"ticketID"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Points      int    `json:"points"`
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
