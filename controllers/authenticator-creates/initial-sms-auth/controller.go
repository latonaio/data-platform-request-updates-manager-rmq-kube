package controllersAuthenticatorCreatesInitialSMSAuth

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsAuthenticator "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/authenticator"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type AuthenticatorCreatesInitialSMSAuthController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *AuthenticatorCreatesInitialSMSAuthController) Post() {
	inputParameter := types.AuthenticatorInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *AuthenticatorCreatesInitialSMSAuthController,
) AuthenticatorCreatesRequestInitialSMSAuth(
	requestPram *types.Request,
	input types.AuthenticatorSDC,
) *apiOutputFormatter.AuthenticatorSDC {
	responseJsonData := apiOutputFormatter.AuthenticatorSDC{}
	responseBody := apiModuleRuntimesRequestsAuthenticator.AuthenticatorCreatesRequestInitialSMSAuth(
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
		controller.CustomLogger.Error("AuthenticatorCreatesRequestInitialSMSAuth Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AuthenticatorCreatesInitialSMSAuthController,
) request(
	input types.AuthenticatorSDC,
) {
	var response *apiOutputFormatter.AuthenticatorSDC

	response = controller.AuthenticatorCreatesRequestInitialSMSAuth(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
