package response

type Teams struct {
	Teams []Team `json:"teams"`
}

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Roles struct {
	Roles []Role `json:"roles"`
}

type Role struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
