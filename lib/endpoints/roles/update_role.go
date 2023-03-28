package roles

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

func UpdateRole(jwtToken string, reqBody request.RoleReqBody, roleId string) (response.Role, error) {
	var resRole response.Role

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resRole, err
	}

	var route string
	route, err = helpers.GetRoute(lib.RouteRolesUpdateRole, roleId)
	if err != nil {
		fmt.Println(err)
		return resRole, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resRole, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resRole, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resRole, err
	}

	err = json.Unmarshal(body, &resRole)
	if err != nil {
		fmt.Println(err)
		return resRole, err
	}
	return resRole, nil
}
