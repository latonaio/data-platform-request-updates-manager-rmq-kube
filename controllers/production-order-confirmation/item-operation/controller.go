package controllersProductionOrderConfirmationItemOperation

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsProductionOrderConfirmation "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/production-order-confirmation"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ProductionOrderConfirmationItemOperationController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

const ()

func (controller *ProductionOrderConfirmationItemOperationController) Post() {
	inputParameter := types.ProductionOrderConfirmationInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *ProductionOrderConfirmationItemOperationController,
) ProductionOrderConfirmationRequestHeaderUpdates(
	requestPram *types.Request,
	input types.ProductionOrderConfirmationSDC,
) *apiOutputFormatter.ProductionOrderConfirmationSDC {
	responseJsonData := apiOutputFormatter.ProductionOrderConfirmationSDC{}
	responseBody := apiModuleRuntimesRequestsProductionOrderConfirmation.ProductionOrderConfirmationRequestHeaderUpdates(
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
		controller.CustomLogger.Error("ProductionOrderConfirmationRequestHeaderUpdates Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ProductionOrderConfirmationItemOperationController,
) request(
	input types.ProductionOrderConfirmationSDC,
) {
	var bRes *apiOutputFormatter.ProductionOrderConfirmationSDC

	bRes = controller.ProductionOrderConfirmationRequestHeaderUpdates(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &bRes
	controller.ServeJSON()
}
