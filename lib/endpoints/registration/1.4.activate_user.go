package registration

import (
	"fmt"
	"net/http"

	"github.com/snowpal/go-status-sdk/lib"
	"github.com/snowpal/go-status-sdk/lib/helpers"
)

func ActivateUser(userId string) error {
	route, err := helpers.GetRoute(lib.RouteRegistrationActivateUser, userId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddAppHeaders(req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
