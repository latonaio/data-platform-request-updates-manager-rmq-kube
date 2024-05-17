package controllersDeliveryDocumentFunctionReferFromOrders

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsDeliveryDocument "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/delivery-document"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type DeliveryDocumentFunctionReferFromOrdersController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *DeliveryDocumentFunctionReferFromOrdersController) Post() {
	inputParameter := types.DeliveryDocumentInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *DeliveryDocumentFunctionReferFromOrdersController,
) DeliveryDocumentFunctionReferFromOrders(
	requestPram *types.Request,
	input types.DeliveryDocumentSDC,
) *apiOutputFormatter.DeliveryDocumentSDC {
	responseJsonData := apiOutputFormatter.DeliveryDocumentSDC{}
	responseBody := apiModuleRuntimesRequestsDeliveryDocument.DeliveryDocumentRequestFunctionReferFromOrders(
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
		controller.CustomLogger.Error("DeliveryDocumentFunctionReferFromOrders Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentFunctionReferFromOrdersController,
) request(
	input types.DeliveryDocumentSDC,
) {
	var response *apiOutputFormatter.DeliveryDocumentSDC

	response = controller.DeliveryDocumentFunctionReferFromOrders(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
