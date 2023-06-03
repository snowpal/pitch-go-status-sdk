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

func UpdatePftOtherItems(
	jwtToken string,
	reqBody request.StatusOtherItemsReqBody,
	statusParam request.StatusParam,
) (response.StatusOtherItems, error) {
	var resOtherItems response.StatusOtherItems

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resOtherItems, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteStatusesUpdatePftOtherItems, statusParam.TeamId, statusParam.StatusId)
	if err != nil {
		fmt.Println(err)
		return resOtherItems, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resOtherItems, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resOtherItems, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resOtherItems, err
	}

	err = json.Unmarshal(body, &resOtherItems)
	if err != nil {
		fmt.Println(err)
		return resOtherItems, err
	}
	return resOtherItems, nil
}
