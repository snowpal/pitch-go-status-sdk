package statuses

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
)

func DeleteSessionOtherItem(jwtToken string, otherItemParam request.OtherItemParam) error {
	route, err := helpers.GetRoute(
		lib.RouteStatusesDeleteSessionOtherItem,
		otherItemParam.TeamId,
		otherItemParam.StatusId,
		otherItemParam.SessionId,
		otherItemParam.OtherItemId,
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
