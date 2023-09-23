package response

type Tickets struct {
	Tickets []Ticket `json:"tickets"`
}

type Ticket struct {
	ID          string `json:"ID"`
	TicketID    string `json:"TicketID"`
	Type        string `json:"Type"`
	Url         string `json:"Url"`
	Description string `json:"Description"`
	Status      string `json:"Status"`
	Points      int    `json:"Points"`
}

type SessionTickets struct {
	Tickets []SessionTicket `json:"Tickets"`
}

type PullRequest struct {
	Url         string `json:"Url"`
	Description string `json:"Description"`
}

type SessionTicket struct {
	ID            string         `json:"ID"`
	TicketID      string         `json:"TicketID"`
	Type          string         `json:"Type"`
	Url           string         `json:"Url"`
	Description   string         `json:"Description"`
	Status        string         `json:"Status"`
	TimeSpent     float32        `json:"TimeSpent"`
	PullRequest   PullRequest    `json:"PullRequest"`
	PairedMembers []PairedMember `json:"PairedWith"`
}

type SessionOtherItem struct {
	ID            string         `json:"ID"`
	Title         string         `json:"Title"`
	Type          string         `json:"Type"`
	TimeSpent     float32        `json:"TimeSpent"`
	PairedMembers []PairedMember `json:"PairedWith"`
	Comment       string         `json:"Comment"`
}

type SessionOtherItems struct {
	OtherItems []SessionOtherItem `json:"OtherItems"`
}
