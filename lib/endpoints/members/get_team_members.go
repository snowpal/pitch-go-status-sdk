package members

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/response"
)

func GetTeamMembers(jwtToken string, teamId string) ([]response.Member, error) {
	var resMembers response.Members

	route, err := helpers.GetRoute(lib.RouteMembersGetTeamMembers, teamId)
	if err != nil {
		fmt.Println(err)
		return resMembers.Members, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resMembers.Members, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resMembers.Members, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resMembers.Members, err
	}

	err = json.Unmarshal(body, &resMembers)
	if err != nil {
		fmt.Println(err)
		return resMembers.Members, err
	}
	return resMembers.Members, nil
}
