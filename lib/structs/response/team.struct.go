package response

import "github.com/snowpal/pitch-go-status-sdk/lib/structs/common"

type Teams struct {
	Teams []Team `json:"teams"`
}

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Members struct {
	Members []common.Member `json:"members"`
}

type BlockPairedMembers struct {
	Members []common.PairedMember `json:"members"`
}
