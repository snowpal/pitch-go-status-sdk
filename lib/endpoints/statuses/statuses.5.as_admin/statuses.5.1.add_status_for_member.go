package statuses

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

func AddStatusForMember(
	jwtToken string,
	reqBody request.StatusReqBody,
	teamParam request.TeamParam,
) (response.Status, error) {
	var resStatus response.Status

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resStatus, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteStatusesAddStatusForMember, teamParam.TeamId, teamParam.MemberId)
	if err != nil {
		fmt.Println(err)
		return resStatus, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resStatus, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStatus, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resStatus, err
	}

	err = json.Unmarshal(body, &resStatus)
	if err != nil {
		fmt.Println(err)
		return resStatus, err
	}
	return resStatus, nil
}
