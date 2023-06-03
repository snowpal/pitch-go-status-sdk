package response

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

type Tickets struct {
	Tickets []Ticket `json:"tickets"`
}

type Ticket struct {
	ID          string `json:"id"`
	TicketID    string `json:"ticketID"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Points      int    `json:"points"`
}

type SessionTickets struct {
	Tickets []SessionTicket `json:"tickets"`
}

type SessionTicket struct {
	ID            string                `json:"id"`
	TicketID      string                `json:"ticketID"`
	Type          string                `json:"type"`
	Url           string                `json:"url"`
	Description   string                `json:"description"`
	Status        string                `json:"status"`
	TimeSpent     float32               `json:"timeSpent"`
	PullRequest   common.PullRequest    `json:"pullRequest"`
	PairedMembers []common.PairedMember `json:"pairedWith"`
}

type SessionOtherItem struct {
	ID            string                `json:"id"`
	Title         string                `json:"title"`
	Type          string                `json:"type"`
	TimeSpent     float32               `json:"timeSpent"`
	PairedMembers []common.PairedMember `json:"pairedWith"`
	Comment       string                `json:"comment"`
}

type SessionOtherItems struct {
	OtherItems []SessionOtherItem `json:"otherItems"`
}
