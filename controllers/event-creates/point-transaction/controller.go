package controllersEventCreatesPointTransaction

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsEvent "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/event"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type EventCreatesPointTransactionController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *EventCreatesPointTransactionController) Post() {
	inputParameter := types.EventInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *EventCreatesPointTransactionController,
) EventCreatesRequestPointTransaction(
	requestPram *types.Request,
	input types.EventSDC,
) *apiOutputFormatter.EventSDC {
	responseJsonData := apiOutputFormatter.EventSDC{}
	responseBody := apiModuleRuntimesRequestsEvent.EventCreatesRequestPointTransaction(
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
		controller.CustomLogger.Error("EventCreatesRequestPointTransaction Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *EventCreatesPointTransactionController,
) request(
	input types.EventSDC,
) {
	var response *apiOutputFormatter.EventSDC

	response = controller.EventCreatesRequestPointTransaction(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
