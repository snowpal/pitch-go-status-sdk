package statuses

import (
	"fmt"
	"net/http"

	"github.com/snowpal/go-status-sdk/lib"
	"github.com/snowpal/go-status-sdk/lib/helpers"
	"github.com/snowpal/go-status-sdk/lib/structs/request"
)

func DeleteSessionOtherItemForMember(jwtToken string, statusParam request.StatusParam) error {
	route, err := helpers.GetRoute(lib.RouteStatusesDeleteSessionOtherItemForMember, statusParam.TeamId, statusParam.MemberId, statusParam.StatusId, statusParam.SessionId, statusParam.OtherItemId)
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