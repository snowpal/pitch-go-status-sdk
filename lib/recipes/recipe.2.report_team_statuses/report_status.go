package teamstatuses

import (
	"github.com/snowpal/pitch-go-status-sdk/lib/endpoints/statuses/statuses.1"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/common"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"
)

func ReportMyStatusForDay(user response.User, team response.Team, statusDate string) (response.Status, error) {
	var resStatus response.Status
	var err error

	resStatus, err = statuses.AddMyStatus(user.JwtToken, request.StatusReqBody{
		MemberId:   user.ID,
		StatusDate: statusDate,
		PlanForToday: request.PlanForToday{
			Tickets: []request.TicketReqBody{
				{
					TicketID:    "GH124",
					Type:        "github",
					Url:         "https://github.com/repos/abc/issues/2",
					Description: "sample desc",
					Status:      "Open",
					Points:      3,
				},
			},
			OtherItems: []string{
				"Meet with Product Managers",
				"Talk to QA",
			},
		},
		Sessions: []request.SessionReqBody{
			{
				SequenceId: 1,
				Tickets: []request.SessionTicketReqBody{
					{
						TicketID:    "GH123",
						Type:        "azure",
						Url:         "https://github.com/repos/abc/issues/1",
						Description: "sample desc",
						TimeSpent:   3.5,
						Status:      "In Progress",
						PairedMembers: []common.PairedMember{
							{
								MemberId: "",
							},
							{
								NonMemberName: "another Team Member, Larry",
							},
						},
						PullRequest: common.PullRequest{
							Url:         "https://github.com/repos/abc/pr/1",
							Description: "PR is ready for review",
						},
					},
				},
				OtherItems: []request.SessionOtherItemReqBody{
					{
						Title:     "Feature discussion with Raj",
						Type:      "Meeting",
						TimeSpent: 1.5,
						PairedMembers: []common.PairedMember{
							{
								NonMemberName: "",
							},
						},
						Comment: "It was a fruitful meeting!",
					},
				},
			},
			{
				SequenceId: 2,
				Tickets: []request.SessionTicketReqBody{
					{
						TicketID:    "GH124",
						Type:        "github",
						Url:         "https://github.com/repos/abc/issues/2",
						Description: "",
						TimeSpent:   1.5,
						Status:      "Closed",
						PairedMembers: []common.PairedMember{
							{
								MemberId: "",
							},
							{
								NonMemberName: "another Team Member, David",
							},
						},
						PullRequest: common.PullRequest{
							Url:         "https://github.com/repos/abc/pr/2",
							Description: "Draft PR is not ready for review",
						},
					},
				},
				OtherItems: []request.SessionOtherItemReqBody{
					{
						Title:     "Bug discussion with Kumar",
						Type:      "Meeting",
						TimeSpent: 2.0,
						PairedMembers: []common.PairedMember{
							{
								MemberId: "",
							},
							{
								NonMemberName: "Jorge",
							},
						},
						Comment: "It was a fruitful meeting!",
					},
				},
			},
		},
		BlockedBy: request.BlockedBy{
			Tickets: []request.TicketReqBody{
				{
					TicketID:    "GH124",
					Type:        "jira",
					Url:         "https://github.com/repos/abc/issues/2",
					Status:      "Open",
					Description: "Not so complex an issue",
				},
			},
			Members: []common.PairedMember{
				{
					MemberId: "",
					Items: []string{
						"Need UI component",
					},
				},
			},
			OtherItems: []string{
				"poker game!",
				"workout sessions",
			},
		},
	}, request.TeamParam{
		TeamId: team.ID,
	})

	if err != nil {
		return resStatus, err
	}

	return resStatus, nil
}
