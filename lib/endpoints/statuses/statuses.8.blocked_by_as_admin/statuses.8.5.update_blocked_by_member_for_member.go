package statuses

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-go-status-sdk/lib"
	"github.com/snowpal/pitch-go-status-sdk/lib/helpers"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/common"
	"github.com/snowpal/pitch-go-status-sdk/lib/structs/request"
)

func UpdateBlockedByMemberForMember(
	jwtToken string,
	reqBody common.PairedMember,
	blockedByMemberParam request.BlockedByMemberParam,
) (common.PairedMember, error) {
	var resPairedMember common.PairedMember

	payload, err := helpers.GetRequestPayload(reqBody)
	if err != nil {
		fmt.Println(err)
		return resPairedMember, err
	}

	var route string
	route, err = helpers.GetRoute(
		lib.RouteStatusesUpdateBlockedByMemberForMember,
		blockedByMemberParam.TeamId,
		blockedByMemberParam.MemberId,
		blockedByMemberParam.StatusId,
		blockedByMemberParam.BlockedByMemberId,
	)
	if err != nil {
		fmt.Println(err)
		return resPairedMember, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resPairedMember, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPairedMember, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPairedMember, err
	}

	err = json.Unmarshal(body, &resPairedMember)
	if err != nil {
		fmt.Println(err)
		return resPairedMember, err
	}
	return resPairedMember, nil
}
