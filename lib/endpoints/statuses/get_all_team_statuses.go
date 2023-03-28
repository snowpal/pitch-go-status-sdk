package statuses

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/go-status-sdk/lib"
	"github.com/snowpal/go-status-sdk/lib/helpers"
	"github.com/snowpal/go-status-sdk/lib/structs/response"
)

func GetAllTeamStatuses(jwtToken string) ([]response.Status, error) {
	var resStatuses response.Statuses

	route, err := helpers.GetRoute(lib.RouteStatusesGetAllTeamStatuses)
	if err != nil {
		fmt.Println(err)
		return resStatuses.Statuses, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
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
