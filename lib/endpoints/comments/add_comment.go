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

func AddComment(jwtToken string, reqBody request.CommentReqBody, statusId string) (response.Comment, error) {
	var resComment response.Comment

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resComment, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteCommentsAddComment, statusId)
	if err != nil {
		fmt.Println(err)
		return resComment, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
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