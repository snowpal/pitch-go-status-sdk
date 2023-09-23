package teamstatuses

import (
	"github.com/snowpal/pitch-go-status-sdk/lib/endpoints/statuses/statuses.1"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func DisplayMyStatusForDay(user response.User, team response.Team, statusDate string) error {
	log.Printf("Displaying status for %s reported on %s", user.Email, statusDate)
	recipes.SleepBefore()
	resStatuses, err := statuses.GetStatusByDate(user.JwtToken, request.StatusParam{
		TeamId:    team.ID,
		StartDate: statusDate,
	})
	if err != nil {
		return err
	}

	// TODO(1, 09/22/23): API returns a list so handle it this way for now.
	resStatus := resStatuses[0]
	log.Info("Status ID: ", resStatus.ID)

	var ticketCount = 0
	for _, session := range resStatus.Sessions {
		ticketCount += len(session.Tickets)
	}
	log.Printf("Cumulative Number of Tickets (across sessions): %d", ticketCount)

	for idx, session := range resStatus.Sessions {
		log.Printf("Session %d Tickets", idx+1)
		for _, ticket := range session.Tickets {
			log.Printf(".Ticket Url: %s", ticket.Url)
		}
	}

	return nil
}
