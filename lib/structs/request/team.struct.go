package request

type TeamReqBody struct {
	Name string `json:"name"`
}

type TeamMemberParam struct {
	TeamId   string
	MemberId string
}

type RoleReqBody struct {
	Title string `json:"title"`
}
