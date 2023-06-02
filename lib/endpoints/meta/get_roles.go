package meta

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/go-status-sdk/lib"
	"github.com/snowpal/go-status-sdk/lib/helpers"
	"github.com/snowpal/go-status-sdk/lib/structs/response"
)

func GetRoles(jwtToken string) ([]response.Role, error) {
	var resRoles response.Roles

	route, err := helpers.GetRoute(lib.RoutesMetaGetRoles)
	if err != nil {
		fmt.Println(err)
		return resRoles.Roles, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resRoles.Roles, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resRoles.Roles, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resRoles.Roles, err
	}

	err = json.Unmarshal(body, &resRoles)
	if err != nil {
		fmt.Println(err)
		return resRoles.Roles, err
	}
	return resRoles.Roles, nil
}
