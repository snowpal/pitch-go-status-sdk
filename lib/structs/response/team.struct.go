package response

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

type Teams struct {
	Teams []Team `json:"teams"`
}

type Team struct {
	ID                  string          `json:"ID"`
	Name                string          `json:"Name"`
	Members             []common.Member `json:"Members"`
	CreatedBy           string          `json:"CreatedBy"`
	UpdatedBy           string          `json:"UpdatedBy"`
	LastModifiedOn      string          `json:"LastModifiedOn"`
	DeletedMembersCount int             `json:"DeletedMembersCount"`
}

type Members struct {
	Members []common.Member `json:"members"`
}

type BlockPairedMembers struct {
	Members []common.PairedMember `json:"members"`
}
