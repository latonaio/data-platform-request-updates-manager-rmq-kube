package controllersDeliveryDocumentItemSingleUnit

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsDeliveryDocuement "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/delivery-document"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type DeliveryDocumentItemSingleUnitController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *DeliveryDocumentItemSingleUnitController) Post() {
	inputParameter := types.DeliveryDocumentInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *DeliveryDocumentItemSingleUnitController,
) DeliveryDocumentRequestItemUpdates(
	requestPram *types.Request,
	input types.DeliveryDocumentSDC,
) *apiOutputFormatter.DeliveryDocumentSDC {
	responseJsonData := apiOutputFormatter.DeliveryDocumentSDC{}
	responseBody := apiModuleRuntimesRequestsDeliveryDocuement.DeliveryDocumentRequestItemUpdates(
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
		controller.CustomLogger.Error("DeliveryDocumentRequestItemUpdates Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentItemSingleUnitController,
) request(
	input types.DeliveryDocumentSDC,
) {
	var response *apiOutputFormatter.DeliveryDocumentSDC

	response = controller.DeliveryDocumentRequestItemUpdates(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
