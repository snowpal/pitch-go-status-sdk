package response

type Tickets struct {
	Tickets []Ticket `json:"tickets"`
}

type Ticket struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Team   *Team   `json:"team"`
	Member *Member `json:"member"`
}
