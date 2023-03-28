package response

type Statuses struct {
	Statuses []Status `json:"statuses"`
}

type Status struct {
	ID            string   `json:"id"`
	Text          string   `json:"text"`
	TaggedMembers []Member `json:"taggedMembers"`
}
