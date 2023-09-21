package teamstatuses

import (
	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func ReportTeamStatuses() {
	log.Info("Objective: Add a new team, add couple of members to that and report status for each 1 of them")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	var user response.User
	user, err = recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}
	log.Info("user ", user.Email, " has been logged in")

	var team response.Team
	team, err = AddNewTeam(user, lib.TeamName)
	if err != nil {
		return
	}

	_, err = AddTeamMembers(user, team)
	if err != nil {
		return
	}

	_, err = ReportMyStatusForDay(user, team, "2023-04-17T00:00:00Z")
	if err != nil {
		return
	}

	DisplayMyStatusForDay(user, team, "2023-04-17T00:00:00Z")
}
