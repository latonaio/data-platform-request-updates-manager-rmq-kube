package apiModuleRuntimesRequestsGoogleAccountInitialAuth

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type GoogleAccountInitialAuthReq struct {
	GoogleAccountInitialAuth GoogleAccountInitialAuth `json:"GoogleAccountInitialAuth"`
	APIType                  string                   `json:"api_type"`
	Accepter                 []string                 `json:"accepter"`
}

type GoogleAccountInitialAuth struct {
}

func FuncGoogleAccountInitialAuthRequests(
	requestPram *types.Request,
	input GoogleAccountInitialAuthReq,
) GoogleAccountInitialAuthReq {
	req := GoogleAccountInitialAuthReq{
		APIType: "requests",
		Accepter: []string{
			"GoogleAccountInitialAuth",
		},
	}
	return req
}

func GoogleAccountInitialAuthRequests(
	requestPram *types.Request,
	input types.GoogleAccountInitialAuthSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_GOOGLE_ACCOUNT_INITIAL_AUTH_SRV"
	aPIType := "requests"

	var request GoogleAccountInitialAuthReq

	request = FuncGoogleAccountInitialAuthRequests(
		requestPram,
		GoogleAccountInitialAuthReq{},
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
