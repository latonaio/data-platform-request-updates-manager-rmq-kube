package apiModuleRuntimesRequestsGoogleAccountAccessToken

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type GoogleAccountAccessTokenReq struct {
	GoogleAccountAccessToken    GoogleAccountAccessToken    `json:"GoogleAccountAccessToken"`
	APIType                     string                      `json:"api_type"`
	Accepter                    []string                    `json:"accepter"`
}

type GoogleAccountAccessToken struct {
	URL string `json:"URL"`
}

func FuncGoogleAccountAccessTokenRequests(
	requestPram *types.Request,
	input GoogleAccountAccessTokenReq,
) GoogleAccountAccessTokenReq {
	req := GoogleAccountAccessTokenReq{
		GoogleAccountAccessToken:	input.GoogleAccountAccessToken,
		APIType: "requests",
		Accepter: []string{
			"GoogleAccountAccessToken",
		},
	}
	return req
}

func GoogleAccountAccessTokenRequests(
	requestPram *types.Request,
	input types.GoogleAccountAccessTokenSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_GOOGLE_ACCOUNT_ACCESS_TOKEN_SRV"
	aPIType := "requests"

	var request GoogleAccountAccessTokenReq

	request = FuncGoogleAccountAccessTokenRequests(
		requestPram,
		GoogleAccountAccessTokenReq{
			GoogleAccountAccessToken: GoogleAccountAccessToken{
				URL:	input.GoogleAccountAccessToken.URL,
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
