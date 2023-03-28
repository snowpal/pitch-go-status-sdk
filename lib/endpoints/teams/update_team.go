package teams

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

func UpdateTeam(jwtToken string, reqBody request.TeamReqBody, teamId string) (response.Team, error) {
	var resTeam response.Team

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resTeam, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteTeamsUpdateTeam, teamId)
	if err != nil {
		fmt.Println(err)
		return resTeam, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resTeam, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resTeam, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resTeam, err
	}

	err = json.Unmarshal(body, &resTeam)
	if err != nil {
		fmt.Println(err)
		return resTeam, err
	}
	return resTeam, nil
}
