package controllersSMSAuthToken

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsSMSAuthToken "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/sms-auth-token"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type SMSAuthTokenController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *SMSAuthTokenController) Post() {
	inputParameter := types.SMSAuthTokenInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *SMSAuthTokenController,
) SMSAuthTokenRequest(
	requestPram *types.Request,
	input types.SMSAuthTokenSDC,
) *apiOutputFormatter.SMSAuthTokenSDC {
	responseJsonData := apiOutputFormatter.SMSAuthTokenSDC{}
	responseBody := apiModuleRuntimesRequestsSMSAuthToken.SMSAuthTokenRequests(
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
		controller.CustomLogger.Error("SMSAuthTokenRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SMSAuthTokenController,
) request(
	input types.SMSAuthTokenSDC,
) {
	var response *apiOutputFormatter.SMSAuthTokenSDC

	response = controller.SMSAuthTokenRequest(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
