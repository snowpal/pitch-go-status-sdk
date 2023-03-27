package comments

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

func GetComment(jwtToken string, param request.CommentParam) (response.Comment, error) {
	resComment := response.Comment{}
	route, err := helpers.GetRoute(lib.RouteCommentsGetComment, param.StatusId, param.CommentId)
	if err != nil {
		fmt.Println(err)
		return resComment, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resComment, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resComment, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resComment, err
	}

	err = json.Unmarshal(body, &resComment)
	if err != nil {
		fmt.Println(err)
		return resComment, err
	}
	return resComment, nil
}
