package request

import "github.com/snowpal/go-status-sdk/lib/structs/common"

type SessionReqBody struct {
	SequenceId int                       `json:"sequenceId"`
	Tickets    []TicketReqBody           `json:"tickets"`
	OtherItems []SessionOtherItemReqBody `json:"otherItems"`
}

type TicketReqBody struct {
	TicketID    string `json:"ticketID"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Points      int    `json:"points"`
}

type SessionOtherItemReqBody struct {
	Title         string                `json:"title"`
	Type          string                `json:"type"`
	TimeSpent     float32               `json:"timeSpent"`
	PairedMembers []common.PairedMember `json:"pairedWith"`
	Comment       string                `json:"comment"`
}
