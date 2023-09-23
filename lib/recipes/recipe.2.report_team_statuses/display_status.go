package teamstatuses

import (
	"github.com/snowpal/pitch-go-status-sdk/lib/endpoints/statuses/statuses.1"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func DisplayMyStatusForDay(user response.User, team response.Team, statusDate string) error {
	log.Info("Displaying status for ", user.Email, " reported on ", statusDate)
	recipes.SleepBefore()

	resStatuses, err := statuses.GetStatusByDate(user.JwtToken, request.StatusParam{
		TeamId:    team.ID,
		StartDate: statusDate,
	})
	if err != nil {
		return err
	}

	resStatus := resStatuses[0]

	log.Info("Status ID: ", resStatus.ID)

	var ticketsCount = 0
	for _, session := range resStatus.Sessions {
		ticketsCount += len(session.Tickets)
	}
	log.Info("Session Tickets Count: ", ticketsCount)

	for idx, session := range resStatus.Sessions {
		log.Info("Session ", idx, " Tickets")
		for _, ticket := range session.Tickets {
			log.Info("Ticket Url: ", ticket.Url)
		}
	}

	return nil
}
