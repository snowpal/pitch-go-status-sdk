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

func UpdateTeamMember(
	jwtToken string,
	reqBody request.TeamMemberReqBody,
	param request.TeamMemberParam,
) (response.Member, error) {
	var resMember response.Member

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resMember, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteMembersUpdateTeamMember, param.TeamId, param.MemberId)
	if err != nil {
		fmt.Println(err)
		return resMember, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
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
