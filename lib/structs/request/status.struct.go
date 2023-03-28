package request

type StatusReqBody struct {
	Text            string `json:"text"`
	TaggedMemberIds string `json:"taggedMemberIds"`
}
