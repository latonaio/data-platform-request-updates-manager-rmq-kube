package apiModuleRuntimesRequestsSMSAuthToken

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/formatter"
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type SMSAuthTokenReq struct {
	SMSAuthToken SMSAuthToken `json:"SMSAuthToken"`
	APIType      string       `json:"api_type"`
	Accepter     []string     `json:"accepter"`
}

type SMSAuthToken struct {
	UserID            string `json:"UserID"`
	MobilePhoneNumber string `json:"MobilePhoneNumber"`
}

func FuncSMSAuthTokenRequests(
	requestPram *types.Request,
	input SMSAuthTokenReq,
) SMSAuthTokenReq {
	req := SMSAuthTokenReq{
		SMSAuthToken: input.SMSAuthToken,
		APIType:      "generates",
		Accepter: []string{
			"SMSAuthToken",
		},
	}
	return req
}

func SMSAuthTokenRequests(
	requestPram *types.Request,
	inputSDC types.SMSAuthTokenSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SMS_AUTH_TOKEN_SRV"
	aPIType := "generates"

	var request SMSAuthTokenReq

	input, err := formatter.ConvertToSMSAuthToken(inputSDC)

	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	request = FuncSMSAuthTokenRequests(
		requestPram,
		SMSAuthTokenReq{
			SMSAuthToken: SMSAuthToken{
				UserID:            input.SMSAuthToken.UserID,
				MobilePhoneNumber: input.SMSAuthToken.MobilePhoneNumber,
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
