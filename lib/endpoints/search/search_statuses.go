package search

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

func SrchStatuses(jwtToken string, reqBody request.SearchReqBody) ([]response.Status, error) {
	var resStatuses response.Statuses

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resStatuses.Statuses, err
	}

	route, err := helpers.GetRoute(lib.RouteSearchSrchStatuses)
	if err != nil {
		fmt.Println(err)
		return resStatuses.Statuses, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, payload)
	if err != nil {
		fmt.Println(err)
		return resStatuses.Statuses, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStatuses.Statuses, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resStatuses.Statuses, err
	}

	err = json.Unmarshal(body, &resStatuses)
	if err != nil {
		fmt.Println(err)
		return resStatuses.Statuses, err
	}
	return resStatuses.Statuses, nil
}
