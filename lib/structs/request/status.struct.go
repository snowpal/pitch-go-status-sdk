package request

type StatusReqBody struct {
	Text            string `json:"text"`
	TaggedMemberIds string `json:"taggedMemberIds"`
}

type StatusOtherItemsReqBody struct {
	OtherStatusItems []string `json:"otherItems"`
}

type StatusSessionOtherItemsReqBody struct {
	OtherStatusSessionItems []OtherStatusSessionItem `json:"otherItems"`
}

type OtherStatusSessionItem struct {
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	TimeSpent     float32        `json:"timeSpent"`
	PairedMembers []PairedMember `json:"pairedWith"`
	Comment       string         `json:"comment"`
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
