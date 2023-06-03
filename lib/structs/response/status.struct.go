package response

import "github.com/snowpal/go-status-sdk/lib/structs/common"

type Statuses struct {
	Statuses []Status `json:"statuses"`
}

type Status struct {
	ID           string              `json:"id"`
	MemberId     string              `json:"memberID"`
	StatusDate   string              `json:"statusDate"`
	PlanForToday common.PlanForToday `json:"planForToday"`
	Sessions     []common.Session    `json:"sessions"`
	BlockedBy    common.BlockedBy    `json:"blockedBy"`
}
