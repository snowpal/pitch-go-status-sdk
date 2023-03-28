package response

type Teams struct {
	Teams []Team `json:"teams"`
}

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Members struct {
	Members []Member `json:"members"`
}

type Member struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Team Team   `json:"team"`
}
