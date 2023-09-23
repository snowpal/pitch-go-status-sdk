package teamstatuses

import (
	"github.com/snowpal/pitch-go-status-sdk/lib/endpoints/statuses/statuses.1"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/common"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func ReportMyStatusForDay(user response.User, team response.Team,
	statusDate string) (response.Status, error) {
	log.Printf("User with %s is reporting status for %s", user.Email, statusDate)

	var memberID = team.Members[1].ID // Pick a random member to pair with
	recipes.SleepBefore()
	resStatus, err := statuses.AddMyStatus(user.JwtToken, request.StatusReqBody{
		StatusDate: statusDate,
		PlanForToday: request.PlanForToday{
			Tickets:    pftTickets(),
			OtherItems: pftOtherItems(),
		},
		Sessions: []request.SessionReqBody{
			{
				SequenceId: 1,
				Tickets:    sessionOneTickets(memberID),
				OtherItems: sessionOneOtherItems(),
			},
			{
				SequenceId: 2,
				Tickets:    sessionTwoTickets(memberID),
				OtherItems: sessionTwoOtherItems(memberID),
			},
		},
		BlockedBy: request.BlockedBy{
			Tickets:    blockedByTickets(),
			Members:    blockedByMembers(memberID),
			OtherItems: blockedByOtherItems(),
		},
	}, request.TeamParam{
		TeamId: team.ID,
	})
	if err != nil {
		return resStatus, err
	}
	recipes.SleepAfter()
	log.Printf(".User with %s has reported status for %s", user.Email, statusDate)

	return resStatus, nil
}

// NOTE: We use hardcoded values simply for illustration purposes. In the real world, you would be
// constructing these via params, of course.
func pftTickets() []request.TicketReqBody {
	return []request.TicketReqBody{
		{
			TicketID:    "GH124",
			Type:        "github",
			Url:         "https://github.com/repos/abc/issues/2",
			Description: "sample desc",
			Status:      "Open",
			Points:      3,
		},
	}
}

func sessionOneTickets(memberID string) []request.SessionTicketReqBody {
	return []request.SessionTicketReqBody{
		{
			TicketID:    "GH123",
			Type:        "azure",
			Url:         "https://github.com/repos/abc/issues/1",
			Description: "sample desc",
			TimeSpent:   3.5,
			Status:      "In Progress",
			PairedMembers: []common.PairedMember{
				{MemberId: &memberID},
				{NonMemberName: "another Team Member, Larry"},
			},
			PullRequest: common.PullRequest{
				Url:         "https://github.com/repos/abc/pr/1",
				Description: "PR is ready for review",
			},
		},
	}
}

func sessionTwoTickets(memberID string) []request.SessionTicketReqBody {
	return []request.SessionTicketReqBody{
		{
			TicketID:    "GH124",
			Type:        "github",
			Url:         "https://github.com/repos/abc/issues/2",
			Description: "",
			TimeSpent:   1.5,
			Status:      "Closed",
			PairedMembers: []common.PairedMember{
				{MemberId: &memberID},
				{NonMemberName: "another Team Member, Larry"},
			},
			PullRequest: common.PullRequest{
				Url:         "https://github.com/repos/abc/pr/2",
				Description: "Draft PR is not ready for review",
			},
		},
	}
}

func blockedByTickets() []request.TicketReqBody {
	return []request.TicketReqBody{
		{
			TicketID:    "GH124",
			Type:        "jira",
			Url:         "https://github.com/repos/abc/issues/2",
			Status:      "Open",
			Description: "Not so complex an issue",
		},
	}
}

func blockedByMembers(memberID string) []common.PairedMember {
	return []common.PairedMember{
		{
			MemberId: &memberID,
			Items: []string{
				"Need UI component",
			},
		},
	}
}

func pftOtherItems() []string {
	return []string{
		"Meet with Product Managers",
		"Talk to QA",
	}
}

func sessionOneOtherItems() []request.SessionOtherItemReqBody {
	return []request.SessionOtherItemReqBody{
		{
			Title:     "Feature discussion with Raj",
			Type:      "Meeting",
			TimeSpent: 1.5,
			PairedMembers: []common.PairedMember{
				{NonMemberName: "Joe"},
			},
			Comment: "It was a fruitful meeting!",
		},
	}
}

func sessionTwoOtherItems(memberID string) []request.SessionOtherItemReqBody {
	return []request.SessionOtherItemReqBody{
		{
			Title:     "Bug discussion with Kumar",
			Type:      "Meeting",
			TimeSpent: 2.0,
			PairedMembers: []common.PairedMember{
				{MemberId: &memberID},
				{NonMemberName: "Jorge"},
			},
			Comment: "It was a fruitful meeting!",
		},
	}
}

func blockedByOtherItems() []string {
	return []string{
		"poker game!",
		"workout sessions",
	}
}
