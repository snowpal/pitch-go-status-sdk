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

func AddSessionTicketsForMember(
	jwtToken string,
	reqBody request.AddSessionTicketsReqBody,
	sessionParam request.SessionParam,
) ([]response.SessionTicket, error) {
	var resTickets response.SessionTickets

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resTickets.Tickets, err
	}

	var route string
	route, err = helpers.GetRoute(
		lib.RouteStatusesAddSessionTicketsForMember,
		sessionParam.TeamId,
		sessionParam.MemberId,
		sessionParam.StatusId,
		sessionParam.SessionId,
	)
	if err != nil {
		fmt.Println(err)
		return resTickets.Tickets, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resTickets.Tickets, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resTickets.Tickets, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resTickets.Tickets, err
	}

	err = json.Unmarshal(body, &resTickets.Tickets)
	if err != nil {
		fmt.Println(err)
		return resTickets.Tickets, err
	}
	return resTickets.Tickets, nil
}
