package controllersPostUpdatesHeader

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsPost "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/post"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type PostUpdatesHeaderController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *PostUpdatesHeaderController) Post() {
	inputParameter := types.PostInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *PostUpdatesHeaderController,
) PostUpdatesRequestHeader(
	requestPram *types.Request,
	input types.PostSDC,
) *apiOutputFormatter.PostSDC {
	responseJsonData := apiOutputFormatter.PostSDC{}
	responseBody := apiModuleRuntimesRequestsPost.PostUpdatesRequestHeader(
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
		controller.CustomLogger.Error("PostUpdatesRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostUpdatesHeaderController,
) request(
	input types.PostSDC,
) {
	var response *apiOutputFormatter.PostSDC

	response = controller.PostUpdatesRequestHeader(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
