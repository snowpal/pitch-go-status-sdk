package request

type TeamReqBody struct {
	Name string `json:"name"`
}

type AddMemberReqBody struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

type UpdateMemberReqBody struct {
	Name string `json:"name"`
}

type TeamMemberParam struct {
	TeamId   string
	MemberId string
}
