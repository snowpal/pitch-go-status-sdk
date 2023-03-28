package tickets

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/go-status-sdk/lib"
	"github.com/snowpal/go-status-sdk/lib/helpers"
	"github.com/snowpal/go-status-sdk/lib/structs/response"
)

func GetMemberTickets(jwtToken string, memberId string) ([]response.Ticket, error) {
	var resTickets response.Tickets

	route, err := helpers.GetRoute(lib.RouteTicketsGetMemberTickets, memberId)
	if err != nil {
		fmt.Println(err)
		return resTickets.Tickets, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
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
