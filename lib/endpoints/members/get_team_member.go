package members

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/go-status-sdk/lib"
	"github.com/snowpal/go-status-sdk/lib/helpers"
	"github.com/snowpal/go-status-sdk/lib/structs/request"
	"github.com/snowpal/go-status-sdk/lib/structs/response"
)

func GetTeamMember(jwtToken string, param request.TeamMemberParam) (response.Member, error) {
	var resMember response.Member

	route, err := helpers.GetRoute(lib.RouteMembersGetTeamMember, param.TeamId, param.MemberId)
	if err != nil {
		fmt.Println(err)
		return resMember, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resMember, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resMember, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resMember, err
	}

	err = json.Unmarshal(body, &resMember)
	if err != nil {
		fmt.Println(err)
		return resMember, err
	}
	return resMember, nil
}
