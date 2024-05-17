package controllersGoogleAccountAuthUserInfo

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsGoogleAccountUserInfo "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/google-account-user-info"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type GoogleAccountAuthUserInfoController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *GoogleAccountAuthUserInfoController) Post() {
	inputParameter := types.GoogleAccountUserInfoInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *GoogleAccountAuthUserInfoController,
) GoogleAccountAuthRequest(
	requestPram *types.Request,
	input types.GoogleAccountUserInfoSDC,
) *apiOutputFormatter.GoogleAccountUserInfoSDC {
	responseJsonData := apiOutputFormatter.GoogleAccountUserInfoSDC{}
	responseBody := apiModuleRuntimesRequestsGoogleAccountUserInfo.GoogleAccountUserInfoRequests(
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
	controller *GoogleAccountAuthUserInfoController,
) request(
	input types.GoogleAccountUserInfoSDC,
) {
	var response *apiOutputFormatter.GoogleAccountUserInfoSDC

	response = controller.GoogleAccountAuthRequest(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
