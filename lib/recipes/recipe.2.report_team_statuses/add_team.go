package teamstatuses

import (
	"github.com/snowpal/pitch-go-status-sdk/lib/endpoints/teams"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func AddNewTeam(user response.User, teamName string) (response.Team, error) {
	var resTeam response.Team

	log.Info("Adding a new team, ", teamName, ".")
	resTeam, err := teams.AddTeam(
		user.JwtToken,
		request.TeamReqBody{Name: teamName})
	if err != nil {
		return resTeam, err
	}
	log.Info(".Team successfully added and the team id is, ", resTeam.ID)

	return resTeam, err
}
