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

func AddSessionsForMember(
	jwtToken string,
	reqBody request.AddSessionsReqBody,
	statusParam request.StatusParam,
) ([]response.Session, error) {
	var resSessions response.Sessions

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resSessions.Sessions, err
	}

	var route string
	route, err = helpers.GetRoute(
		lib.RouteStatusesAddSessionsForMember,
		statusParam.TeamId,
		statusParam.MemberId,
		statusParam.StatusId,
	)
	if err != nil {
		fmt.Println(err)
		return resSessions.Sessions, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resSessions.Sessions, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resSessions.Sessions, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resSessions.Sessions, err
	}

	err = json.Unmarshal(body, &resSessions.Sessions)
	if err != nil {
		fmt.Println(err)
		return resSessions.Sessions, err
	}
	return resSessions.Sessions, nil
}
