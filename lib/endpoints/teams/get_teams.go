package teams

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/go-status-sdk/lib"
	"github.com/snowpal/go-status-sdk/lib/helpers"
	"github.com/snowpal/go-status-sdk/lib/structs/response"
)

func GetTeams(jwtToken string) ([]response.Team, error) {
	var resTeams response.Teams

	route, err := helpers.GetRoute(lib.RouteTeamsGetTeams)
	if err != nil {
		fmt.Println(err)
		return resTeams.Teams, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resTeams.Teams, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resTeams.Teams, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resTeams.Teams, err
	}

	err = json.Unmarshal(body, &resTeams)
	if err != nil {
		fmt.Println(err)
		return resTeams.Teams, err
	}
	return resTeams.Teams, nil
}
