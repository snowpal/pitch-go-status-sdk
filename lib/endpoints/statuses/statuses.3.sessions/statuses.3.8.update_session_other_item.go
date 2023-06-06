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

func UpdateSessionOtherItem(
	jwtToken string,
	reqBody request.SessionOtherItemReqBody,
	otherItemParam request.OtherItemParam,
) (response.SessionOtherItem, error) {
	var resOtherItem response.SessionOtherItem

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resOtherItem, err
	}

	var route string
	route, err = helpers.GetRoute(
		lib.RouteStatusesUpdateSessionOtherItem,
		otherItemParam.TeamId,
		otherItemParam.StatusId,
		otherItemParam.SessionId,
		otherItemParam.OtherItemId,
	)
	if err != nil {
		fmt.Println(err)
		return resOtherItem, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resOtherItem, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resOtherItem, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resOtherItem, err
	}

	err = json.Unmarshal(body, &resOtherItem)
	if err != nil {
		fmt.Println(err)
		return resOtherItem, err
	}
	return resOtherItem, nil
}
