package pullRequests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/go-status-sdk/lib"
	"github.com/snowpal/go-status-sdk/lib/helpers"
	"github.com/snowpal/go-status-sdk/lib/structs/response"
)

func GetMemberPullRequests(jwtToken string, memberId string) ([]response.PullRequest, error) {
	var resPullRequests response.PullRequests

	route, err := helpers.GetRoute(lib.RoutePullRequestsGetMemberPullRequests, memberId)
	if err != nil {
		fmt.Println(err)
		return resPullRequests.PullRequests, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resPullRequests.PullRequests, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPullRequests.PullRequests, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPullRequests.PullRequests, err
	}

	err = json.Unmarshal(body, &resPullRequests)
	if err != nil {
		fmt.Println(err)
		return resPullRequests.PullRequests, err
	}
	return resPullRequests.PullRequests, nil
}
