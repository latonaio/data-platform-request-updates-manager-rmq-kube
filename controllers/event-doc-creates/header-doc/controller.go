package controllersEventDocCreatesHeaderDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsEventDoc "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/event-doc"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type EventDocCreatesHeaderDocController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *EventDocCreatesHeaderDocController) Post() {
	inputParameter := types.EventDocInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *EventDocCreatesHeaderDocController,
) EventDocCreatesHeaderDocController(
	requestPram *types.Request,
	input types.EventDocSDC,
) *apiOutputFormatter.EventDocSDC {
	responseJsonData := apiOutputFormatter.EventDocSDC{}
	responseBody := apiModuleRuntimesRequestsEventDoc.EventDocCreatesRequestHeaderDoc(
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
		controller.CustomLogger.Error("EventDocCreatesHeaderDoc Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *EventDocCreatesHeaderDocController,
) request(
	input types.EventDocSDC,
) {
	var response *apiOutputFormatter.EventDocSDC

	response = controller.EventDocCreatesHeaderDocController(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
