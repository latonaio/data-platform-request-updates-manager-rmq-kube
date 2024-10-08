package controllersPostCreatesHeader

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsPost "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/post"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type PostCreatesHeaderController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *PostCreatesHeaderController) Post() {
	inputParameter := types.PostInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *PostCreatesHeaderController,
) PostCreatesRequestHeader(
	requestPram *types.Request,
	input types.PostSDC,
) *apiOutputFormatter.PostSDC {
	responseJsonData := apiOutputFormatter.PostSDC{}
	responseBody := apiModuleRuntimesRequestsPost.PostCreatesRequestHeader(
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
		controller.CustomLogger.Error("PostCreatesRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostCreatesHeaderController,
) request(
	input types.PostSDC,
) {
	var response *apiOutputFormatter.PostSDC

	response = controller.PostCreatesRequestHeader(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
