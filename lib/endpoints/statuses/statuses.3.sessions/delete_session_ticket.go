package statuses

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
)

func DeleteSessionTicket(jwtToken string, ticketParam request.TicketParam) error {
	route, err := helpers.GetRoute(
		lib.RouteStatusesDeleteSessionTicket,
		ticketParam.TeamId,
		ticketParam.StatusId,
		ticketParam.SessionId,
		ticketParam.TicketId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
