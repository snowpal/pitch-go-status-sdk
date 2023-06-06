package request

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

type TeamReqBody struct {
	Name string `json:"name"`
}

type MembersReqBody struct {
	Members []common.Member `json:"members"`
}

type BlockPairedMembersReqBody struct {
	Members []common.PairedMember `json:"members"`
}

type BlockedByMemberParam struct {
	TeamId            string
	MemberId          string
	StatusId          string
	BlockedByMemberId string
}
