package response

type Roles struct {
	Roles []Role `json:"roles"`
}

type Role struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
