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

func UpdateSessionTicketForMember(
	jwtToken string,
	reqBody request.SessionTicketReqBody,
	ticketParam request.TicketParam,
) (response.SessionTicket, error) {
	var resTicket response.SessionTicket

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resTicket, err
	}

	var route string
	route, err = helpers.GetRoute(
		lib.RouteStatusesUpdateSessionTicket,
		ticketParam.TeamId,
		ticketParam.StatusId,
		ticketParam.SessionId,
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
