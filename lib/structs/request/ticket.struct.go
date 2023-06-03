package request

import "github.com/snowpal/go-status-sdk/lib/structs/common"

type AddSessionsReqBody struct {
	Sessions []common.Session `json:"sessions"`
}

type StatusSessionOtherItemsReqBody struct {
	SessionOtherItems []common.SessionOtherItemReqBody `json:"otherItems"`
}

type AddSessionTicketsReqBody struct {
	Tickets []SessionTicketReqBody `json:"tickets"`
}

type SessionTicketReqBody struct {
	ID            string                `json:"ticketID"`
	Type          string                `json:"type"`
	Url           string                `json:"url"`
	Description   string                `json:"description"`
	Status        string                `json:"status"`
	TimeSpent     float32               `json:"timeSpent"`
	PullRequest   PullRequest           `json:"pullRequest"`
	PairedMembers []common.PairedMember `json:"pairedWith"`
}

type AddTicketsReqBody struct {
	Tickets []common.TicketReqBody `json:"tickets"`
}

type PullRequest struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}
