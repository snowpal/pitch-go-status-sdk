package teamstatuses

import (
	"github.com/snowpal/pitch-go-status-sdk/lib/endpoints/statuses/statuses.1"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func DisplayMyStatusForDay(user response.User, team response.Team, statusDate string) error {
	resStatus, err := statuses.GetStatusByDate(user.JwtToken, request.StatusParam{
		TeamId:    team.ID,
		MemberId:  user.ID,
		StartDate: statusDate,
	})
	if err != nil {
		return err
	}

	log.Info("Status ID: ", resStatus.ID)

	return nil
}
