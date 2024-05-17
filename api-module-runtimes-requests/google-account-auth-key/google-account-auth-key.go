package apiModuleRuntimesRequestsGoogleAccountAuthKey

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type GoogleAccountAuthKeyReq struct {
	GoogleAccountAuthKey    GoogleAccountAuthKey    `json:"GoogleAccountAuthKey"`
	APIType                 string                  `json:"api_type"`
	Accepter                []string                `json:"accepter"`
}

type GoogleAccountAuthKey struct {
	URL string `json:"URL"`
}

func FuncGoogleAccountAuthKeyRequests(
	requestPram *types.Request,
	input GoogleAccountAuthKeyReq,
) GoogleAccountAuthKeyReq {
	req := GoogleAccountAuthKeyReq{
		GoogleAccountAuthKey:	input.GoogleAccountAuthKey,
		APIType: "requests",
		Accepter: []string{
			"GoogleAccountAuthKey",
		},
	}
	return req
}

func GoogleAccountAuthKeyRequests(
	requestPram *types.Request,
	input types.GoogleAccountAuthKeySDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_GOOGLE_ACCOUNT_AUTH_KEY_SRV"
	aPIType := "requests"

	var request GoogleAccountAuthKeyReq

	request = FuncGoogleAccountAuthKeyRequests(
		requestPram,
		GoogleAccountAuthKeyReq{
			GoogleAccountAuthKey: GoogleAccountAuthKey{
				URL:	input.GoogleAccountAuthKey.URL,
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
