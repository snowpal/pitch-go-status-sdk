package request

type StatusReqBody struct {
	Text            string `json:"text"`
	TaggedMemberIds string `json:"taggedMemberIds"`
}

type StatusOtherItemsReqBody struct {
	OtherStatusItems []string `json:"otherItems"`
}

type TeamParam struct {
	TeamId string
}

type StatusParam struct {
	TeamId      string
	StatusId    string
	StartDate   string
	EndDate     string
	SessionId   string
	MemberId    string
	TicketId    string
	OtherItemId string
}
