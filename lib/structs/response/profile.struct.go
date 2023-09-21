package response

type Profile struct {
	UserID    string `json:"userID"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
