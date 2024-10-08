package controllersMessageCreatesAll

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsMessage "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/message"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type MessageCreatesAllController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *MessageCreatesAllController) Post() {
	inputParameter := types.MessageInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *MessageCreatesAllController,
) MessageCreatesRequestAll(
	requestPram *types.Request,
	input types.MessageSDC,
) *apiOutputFormatter.MessageSDC {
	responseJsonData := apiOutputFormatter.MessageSDC{}
	responseBody := apiModuleRuntimesRequestsMessage.MessageCreatesRequestAll(
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
		controller.CustomLogger.Error("MessageCreatesRequestAll Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *MessageCreatesAllController,
) request(
	input types.MessageSDC,
) {
	var response *apiOutputFormatter.MessageSDC

	response = controller.MessageCreatesRequestAll(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
