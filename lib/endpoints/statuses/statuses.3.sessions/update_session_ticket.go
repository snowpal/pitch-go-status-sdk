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

func UpdateSessionTicket(jwtToken string, reqBody request.TicketReqBody, statusParam request.StatusParam) (response.Status, error) {
	var resStatus response.Status

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resStatus, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteStatusesUpdateSessionTicket, statusParam.TeamId, statusParam.StatusId, statusParam.TicketId)
	if err != nil {
		fmt.Println(err)
		return resStatus, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
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
