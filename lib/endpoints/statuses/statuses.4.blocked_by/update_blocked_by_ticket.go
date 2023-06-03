package statuses

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

func UpdateBlockedByTicket(
	jwtToken string,
	reqBody request.TicketReqBody,
	ticketParam request.TicketParam,
) (response.Ticket, error) {
	var resTicket response.Ticket

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resTicket, err
	}

	var route string
	route, err = helpers.GetRoute(
		lib.RouteStatusesUpdateBlockedByTicket,
		ticketParam.TeamId,
		ticketParam.StatusId,
		ticketParam.TicketId,
	)
	if err != nil {
		fmt.Println(err)
		return resTicket, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resTicket, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resTicket, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resTicket, err
	}

	err = json.Unmarshal(body, &resTicket)
	if err != nil {
		fmt.Println(err)
		return resTicket, err
	}
	return resTicket, nil
}
