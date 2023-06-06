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

func AddSessionOtherItems(
	jwtToken string,
	reqBody request.StatusSessionOtherItemsReqBody,
	sessionParam request.SessionParam,
) ([]response.SessionOtherItem, error) {
	var resOtherItems response.SessionOtherItems

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resOtherItems.OtherItems, err
	}

	var route string
	route, err = helpers.GetRoute(
		lib.RouteStatusesAddSessionOtherItems,
		sessionParam.TeamId,
		sessionParam.StatusId,
		sessionParam.SessionId,
	)
	if err != nil {
		fmt.Println(err)
		return resOtherItems.OtherItems, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resOtherItems.OtherItems, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resOtherItems.OtherItems, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resOtherItems.OtherItems, err
	}

	err = json.Unmarshal(body, &resOtherItems)
	if err != nil {
		fmt.Println(err)
		return resOtherItems.OtherItems, err
	}
	return resOtherItems.OtherItems, nil
}
