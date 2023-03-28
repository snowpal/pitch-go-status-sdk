package storyPoints

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/go-status-sdk/lib"
	"github.com/snowpal/go-status-sdk/lib/helpers"
	"github.com/snowpal/go-status-sdk/lib/structs/response"
)

func GetTeamStoryPoints(jwtToken string, teamId string) ([]response.StoryPoint, error) {
	var resStoryPoints response.StoryPoints

	route, err := helpers.GetRoute(lib.RouteStoryPointsGetTeamStoryPoints, teamId)
	if err != nil {
		fmt.Println(err)
		return resStoryPoints.StoryPoints, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resStoryPoints.StoryPoints, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStoryPoints.StoryPoints, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resStoryPoints.StoryPoints, err
	}

	err = json.Unmarshal(body, &resStoryPoints)
	if err != nil {
		fmt.Println(err)
		return resStoryPoints.StoryPoints, err
	}
	return resStoryPoints.StoryPoints, nil
}
