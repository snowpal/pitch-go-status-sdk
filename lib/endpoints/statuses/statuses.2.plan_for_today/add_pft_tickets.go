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

func AddPftTickets(jwtToken string,
	reqBody request.AddTicketsReqBody,
	statusParam request.StatusParam,
) ([]response.Ticket, error) {
	var resTickets response.Tickets

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resTickets.Tickets, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteStatusesAddPftTickets, statusParam.TeamId, statusParam.StatusId)
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

	err = json.Unmarshal(body, &resTickets)
	if err != nil {
		fmt.Println(err)
		return resTickets.Tickets, err
	}
	return resTickets.Tickets, nil
}
