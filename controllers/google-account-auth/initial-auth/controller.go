package controllersGoogleAccountAuthInitialAuth

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsGoogleAccountInitialAuth "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/google-account-intial-auth"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type GoogleAccountAuthInitialAuthController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *GoogleAccountAuthInitialAuthController) Post() {
	inputParameter := types.GoogleAccountInitialAuthInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *GoogleAccountAuthInitialAuthController,
) GoogleAccountAuthRequest(
	requestPram *types.Request,
	input types.GoogleAccountInitialAuthSDC,
) *apiOutputFormatter.GoogleAccountInitialAuthSDC {
	responseJsonData := apiOutputFormatter.GoogleAccountInitialAuthSDC{}
	responseBody := apiModuleRuntimesRequestsGoogleAccountInitialAuth.GoogleAccountInitialAuthRequests(
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
		controller.CustomLogger.Error("GoogleAccountAuthRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *GoogleAccountAuthInitialAuthController,
) request(
	input types.GoogleAccountInitialAuthSDC,
) {
	var response *apiOutputFormatter.GoogleAccountInitialAuthSDC

	response = controller.GoogleAccountAuthRequest(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
