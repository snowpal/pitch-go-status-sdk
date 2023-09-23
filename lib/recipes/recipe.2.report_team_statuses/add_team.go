package teamstatuses

import (
	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/endpoints/members"
	"github.com/snowpal/pitch-go-status-sdk/lib/endpoints/profiles"
	"github.com/snowpal/pitch-go-status-sdk/lib/endpoints/teams"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/common"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func AddNewTeam(user response.User, teamName string) (response.Team, error) {
	var resTeam response.Team

	log.Info("Adding a new team, ", teamName, ".")
	recipes.SleepBefore()
	resTeam, err := teams.AddTeam(
		user.JwtToken,
		request.TeamReqBody{Name: teamName})
	if err != nil {
		return resTeam, err
	}
	recipes.SleepAfter()
	log.Info(".Team successfully added and the team id is, ", resTeam.ID)

	return resTeam, err
}

func AddTeamMembers(user response.User, team response.Team) (response.Team, error) {
	var resTeam response.Team
	var err error

	log.Info("Adding members to team, ", team.Name, ".")
	recipes.SleepBefore()
	var secondProfile response.Profile
	secondProfile, err = getUserProfile(lib.SecondUser)
	if err != nil {
		return resTeam, err
	}

	var thirdProfile response.Profile
	thirdProfile, err = getUserProfile(lib.ThirdUser)
	if err != nil {
		return resTeam, err
	}

	resTeam, err = members.AddMembersToTeam(user.JwtToken, request.MembersReqBody{
		Members: []common.MemberReqBody{
			{ID: secondProfile.ID, Role: "member"},
			{ID: thirdProfile.ID, Role: "admin"},
		},
	}, team.ID)
	if err != nil {
		return resTeam, err
	}
	recipes.SleepBefore()
	log.Info(".Members were added to team, ", team.Name, ".")

	return resTeam, err
}

// private functions

func getUserProfile(userEmail string) (response.Profile, error) {
	var resProfile response.Profile

	user, err := recipes.SignIn(userEmail, lib.Password)
	if err != nil {
		return resProfile, err
	}

	resProfile, err = profiles.GetMyProfile(user.JwtToken)
	if err != nil {
		return resProfile, err
	}

	return resProfile, nil
}
