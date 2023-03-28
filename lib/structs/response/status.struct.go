package response

type Status struct {
	ID            string   `json:"id"`
	Text          string   `json:"text"`
	TaggedMembers []Member `json:"taggedMembers"`
}
