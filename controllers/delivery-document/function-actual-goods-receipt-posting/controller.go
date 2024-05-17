package controllersDeliveryDocumentFunctionActualGoodsReceiptPosting

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsDeliveryDocuement "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/delivery-document"
	apiModuleRuntimesResponsesDeliveryDocument "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-responses/delivery-document"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"time"
)

type DeliveryDocumentFunctionActualGoodsReceiptPostingController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *DeliveryDocumentFunctionActualGoodsReceiptPostingController) Post() {
	inputParameter := types.DeliveryDocumentInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *DeliveryDocumentFunctionActualGoodsReceiptPostingController,
) createDeliveryDocumentRequestItem(
	requestPram *types.Request,
	input types.DeliveryDocumentSDC,
) *apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes {
	responseJsonData := apiModuleRuntimesResponsesDeliveryDocument.DeliveryDocumentRes{}
	responseBody := apiModuleRuntimesRequestsDeliveryDocuement.DeliveryDocumentReads(
		requestPram,
		input,
		&controller.Controller,
		"Item",
	)

	err := json.Unmarshal(responseBody, &responseJsonData)
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("createDeliveryDocumentRequestItem Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *DeliveryDocumentFunctionActualGoodsReceiptPostingController,
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
	controller *DeliveryDocumentFunctionActualGoodsReceiptPostingController,
) request(
	input types.DeliveryDocumentSDC,
) {
	var response *apiOutputFormatter.DeliveryDocumentSDC

	deliveryDocumentItemRes := controller.createDeliveryDocumentRequestItem(
		controller.UserInfo,
		input,
	)

	currentDateTime := time.Now()
	currentDate := currentDateTime.Format("2006-01-02")
	currentTime := currentDateTime.Format("15:04:05")

	for _, item := range *deliveryDocumentItemRes.Message.Item {
		input.Header.Item[0].ActualGoodsReceiptDate = &currentDate
		input.Header.Item[0].ActualGoodsReceiptTime = &currentTime
		input.Header.Item[0].ActualGoodsReceiptQuantity = &item.PlannedGoodsReceiptQuantity
		input.Header.Item[0].ActualGoodsReceiptQtyInBaseUnit = &item.PlannedGoodsReceiptQtyInBaseUnit
	}

	response = controller.DeliveryDocumentRequestItemUpdates(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
