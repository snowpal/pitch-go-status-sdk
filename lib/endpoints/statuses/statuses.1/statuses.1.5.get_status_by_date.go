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

func GetStatusByDate(jwtToken string, statusParam request.StatusParam) ([]response.Status, error) {
	var resStatuses []response.Status

	route, err := helpers.GetRoute(lib.RouteStatusesGetStatusByDate, statusParam.TeamId, statusParam.StartDate)
	if err != nil {
		fmt.Println(err)
		return resStatuses, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resStatuses, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStatuses, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resStatuses, err
	}

	err = json.Unmarshal(body, &resStatuses)
	if err != nil {
		fmt.Println(err)
		return resStatuses, err
	}
	return resStatuses, nil
}
