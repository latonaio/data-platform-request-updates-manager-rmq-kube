package apiModuleRuntimesRequestsGoogleAccountUserInfo

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type GoogleAccountUserInfoReq struct {
	GoogleAccountUserInfo GoogleAccountUserInfo `json:"GoogleAccountUserInfo"`
	APIType               string                `json:"api_type"`
	Accepter              []string              `json:"accepter"`
}

type GoogleAccountUserInfo struct {
	AccessToken string `json:"AccessToken"`
}

func FuncGoogleAccountUserInfoRequests(
	requestPram *types.Request,
	input GoogleAccountUserInfoReq,
) GoogleAccountUserInfoReq {
	req := GoogleAccountUserInfoReq{
		GoogleAccountUserInfo: input.GoogleAccountUserInfo,
		APIType:               "requests",
		Accepter: []string{
			"GoogleAccountUserInfo",
		},
	}
	return req
}

func GoogleAccountUserInfoRequests(
	requestPram *types.Request,
	input types.GoogleAccountUserInfoSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_GOOGLE_ACCOUNT_USER_INFO_SRV"
	aPIType := "requests"

	var request GoogleAccountUserInfoReq

	request = FuncGoogleAccountUserInfoRequests(
		requestPram,
		GoogleAccountUserInfoReq{
			GoogleAccountUserInfo: GoogleAccountUserInfo{
				AccessToken: input.GoogleAccountUserInfo.AccessToken,
			},
		},
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}
