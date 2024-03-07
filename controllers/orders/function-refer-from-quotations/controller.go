package controllersOrdersFunctionReferFromQuotations

import (
	apiInputReader "data-platform-request-updates-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsOrders "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/orders"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type OrdersFunctionReferFromQuotationsController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

func (controller *OrdersFunctionReferFromQuotationsController) Post() {
	inputParameter := apiInputReader.OrdersInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *OrdersFunctionReferFromQuotationsController,
) OrdersRequestFunctionReferFromQuotations(
	requestPram *apiInputReader.Request,
	input apiInputReader.OrdersSDC,
) *apiOutputFormatter.OrdersSDC {
	responseJsonData := apiOutputFormatter.OrdersSDC{}
	responseBody := apiModuleRuntimesRequestsOrders.OrdersRequestFunctionReferFromQuotations(
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
		controller.CustomLogger.Error("OrdersRequestFunctionReferFromQuotations Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersFunctionReferFromQuotationsController,
) request(
	input apiInputReader.OrdersSDC,
) {
	var response *apiOutputFormatter.OrdersSDC

	response = controller.OrdersRequestFunctionReferFromQuotations(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
