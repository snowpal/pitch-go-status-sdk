package members

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"
)

func AddMemberToTeam(jwtToken string, reqBody request.AddMemberReqBody, teamId string) (response.Member, error) {
	var resMember response.Member

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resMember, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteMembersAddMemberToTeam, teamId)
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
