package teamstatuses

import (
	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/endpoints/members"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/common"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func AddTeamMembers(user response.User, team response.Team) ([]common.Member, error) {
	var resMembers []common.Member
	var err error

	var user2 response.User
	user2, err = recipes.SignIn(lib.User2, lib.Password)
	if err != nil {
		return resMembers, err
	}

	var user3 response.User
	user3, err = recipes.SignIn(lib.User3, lib.Password)
	if err != nil {
		return resMembers, err
	}

	log.Info("Adding members to team ", team.Name, ".")
	resMembers, err = members.AddMembersToTeam(user.JwtToken, request.MembersReqBody{
		Members: []common.Member{
			{ID: user2.ID, Role: "member"},
			{ID: user3.ID, Role: "member"},
		},
	}, team.ID)
	if err != nil {
		return resMembers, err
	}

	return resMembers, err
}
