package teamstatuses

import (
	"errors"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func AddTeamAndMembersAsAdmin() (response.Team, error) {
	var team response.Team

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return team, err
	}

	team, err = AddNewTeam(user, lib.TeamName)
	if err != nil {
		return team, err
	}

	team, err = AddTeamMembers(user, team)
	if err != nil {
		return team, err
	}
	return team, nil
}

func ReportStatusAsTeamMember(team response.Team) error {
	user, err := recipes.SignIn(lib.SecondUser, lib.Password)
	if err != nil {
		return err
	}

	if len(team.Members) == 0 {
		return errors.New(".no members in the team")
	}

	var statusDate = "2023-05-23T00:00:00Z"
	_, err = ReportMyStatusForDay(user, team.ID, team.Members[1].ID, statusDate)
	if err != nil {
		return err
	}

	DisplayMyStatusForDay(user, team, statusDate)

	return nil
}

func ReportTeamStatuses() {
	log.Info("Objective: Add a new team, add couple of members to that and report status for each 1 of them")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	var team response.Team
	team, err = AddTeamAndMembersAsAdmin()
	if err != nil {
		return
	}

	ReportStatusAsTeamMember(team)
}
