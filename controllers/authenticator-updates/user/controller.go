package controllersAuthenticatorUpdatesUser

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsAuthenticator "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/authenticator"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type AuthenticatorUpdatesUserController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *AuthenticatorUpdatesUserController) Post() {
	inputParameter := types.AuthenticatorInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *AuthenticatorUpdatesUserController,
) AuthenticatorUpdatesRequestUser(
	requestPram *types.Request,
	input types.AuthenticatorSDC,
) *apiOutputFormatter.AuthenticatorSDC {
	responseJsonData := apiOutputFormatter.AuthenticatorSDC{}
	responseBody := apiModuleRuntimesRequestsAuthenticator.AuthenticatorUpdatesRequestUser(
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
		controller.CustomLogger.Error("AuthenticatorUpdatesRequestUser Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AuthenticatorUpdatesUserController,
) request(
	input types.AuthenticatorSDC,
) {
	var response *apiOutputFormatter.AuthenticatorSDC

	response = controller.AuthenticatorUpdatesRequestUser(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
