package controllersOrdersItemSingleUnit

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsOrders "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/orders"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type OrdersItemSingleUnitController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *OrdersItemSingleUnitController) Post() {
	inputParameter := types.OrdersInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *OrdersItemSingleUnitController,
) OrdersRequestItemUpdates(
	requestPram *types.Request,
	input types.OrdersSDC,
) *apiOutputFormatter.OrdersSDC {
	responseJsonData := apiOutputFormatter.OrdersSDC{}
	responseBody := apiModuleRuntimesRequestsOrders.OrdersRequestItemUpdates(
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
		controller.CustomLogger.Error("OrdersRequestItemUpdates Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *OrdersItemSingleUnitController,
) request(
	input types.OrdersSDC,
) {
	var response *apiOutputFormatter.OrdersSDC

	response = controller.OrdersRequestItemUpdates(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
