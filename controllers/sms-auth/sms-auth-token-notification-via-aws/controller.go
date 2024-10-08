package controllersSMSAuthTokenNotificationViaAWS

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsSMSAuthTokenNotificationViaAWS "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/sms-auth-token-notification-via-aws"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type SMSAuthTokenNotificationViaAWSController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *SMSAuthTokenNotificationViaAWSController) Post() {
	inputParameter := types.SMSAuthTokenNotificationViaAWSInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *SMSAuthTokenNotificationViaAWSController,
) SMSAuthTokenNotificationViaAWSRequest(
	requestPram *types.Request,
	input types.SMSAuthTokenNotificationViaAWSSDC,
) *apiOutputFormatter.SMSAuthTokenNotificationViaAWSSDC {
	responseJsonData := apiOutputFormatter.SMSAuthTokenNotificationViaAWSSDC{}
	responseBody := apiModuleRuntimesRequestsSMSAuthTokenNotificationViaAWS.SMSAuthTokenNotificationViaAWSRequests(
		requestPram,
		input,
		&controller.Controller,
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("SMSAuthTokenNotificationViaAWSRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SMSAuthTokenNotificationViaAWSController,
) request(
	input types.SMSAuthTokenNotificationViaAWSSDC,
) {
	var response *apiOutputFormatter.SMSAuthTokenNotificationViaAWSSDC

	response = controller.SMSAuthTokenNotificationViaAWSRequest(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
