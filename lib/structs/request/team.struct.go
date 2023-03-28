package request

type TeamReqBody struct {
	Name string `json:"name"`
}

type TeamMemberReqBody struct {
	Name string `json:"name"`
}

type TeamMemberParam struct {
	TeamId   string
	MemberId string
}
