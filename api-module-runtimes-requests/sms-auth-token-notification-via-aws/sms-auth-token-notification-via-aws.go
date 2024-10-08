package apiModuleRuntimesRequestsSMSAuthTokenNotificationViaAWS

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type SMSAuthTokenNotificationViaAWSReq struct {
	SMSAuthTokenNotificationViaAWS SMSAuthTokenNotificationViaAWS `json:"SMSAuthTokenNotificationViaAWS"`
	APIType                        string                         `json:"api_type"`
	Accepter                       []string                       `json:"accepter"`
}

type SMSAuthTokenNotificationViaAWS struct {
	MobilePhoneNumber  string `json:"MobilePhoneNumber"`
	AuthenticationCode int    `json:"AuthenticationCode"`
}

func FuncSMSAuthTokenNotificationViaAWSRequests(
	requestPram *types.Request,
	input SMSAuthTokenNotificationViaAWSReq,
) SMSAuthTokenNotificationViaAWSReq {
	req := SMSAuthTokenNotificationViaAWSReq{
		SMSAuthTokenNotificationViaAWS: input.SMSAuthTokenNotificationViaAWS,
		APIType:                        "generates",
		Accepter: []string{
			"SMSAuthTokenNotificationViaAWS",
		},
	}
	return req
}

func SMSAuthTokenNotificationViaAWSRequests(
	requestPram *types.Request,
	input types.SMSAuthTokenNotificationViaAWSSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SMS_AUTH_TOKEN_NOTIFICATION_VIA_AWS_SRV"
	aPIType := "generates"

	var request SMSAuthTokenNotificationViaAWSReq

	mobilePhoneNumber, err := services.ConvertToInternationalPhoneNumber(input.SMSAuthTokenNotificationViaAWS.MobilePhoneNumber)

	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	request = FuncSMSAuthTokenNotificationViaAWSRequests(
		requestPram,
		SMSAuthTokenNotificationViaAWSReq{
			SMSAuthTokenNotificationViaAWS: SMSAuthTokenNotificationViaAWS{
				MobilePhoneNumber:  mobilePhoneNumber,
				AuthenticationCode: input.SMSAuthTokenNotificationViaAWS.AuthenticationCode,
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
