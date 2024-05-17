package controllersEventCreatesAll

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsEvent "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/event"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type EventCreatesAllController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *EventCreatesAllController) Post() {
	inputParameter := types.EventInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *EventCreatesAllController,
) EventCreatesRequestAll(
	requestPram *types.Request,
	input types.EventSDC,
) *apiOutputFormatter.EventSDC {
	responseJsonData := apiOutputFormatter.EventSDC{}
	responseBody := apiModuleRuntimesRequestsEvent.EventCreatesRequestAll(
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
		controller.CustomLogger.Error("EventCreatesRequestAll Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *EventCreatesAllController,
) request(
	input types.EventSDC,
) {
	var response *apiOutputFormatter.EventSDC

	response = controller.EventCreatesRequestAll(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
