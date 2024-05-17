package apiModuleRuntimesRequestsProductionOrderConfirmation

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ProductionOrderConfirmationReq struct {
	Header   Header   `json:"ProductionOrderConfirmation"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	ProductionOrder                          int      `json:"ProductionOrder"`
	ProductionOrderItem                      int      `json:"ProductionOrderItem"`
	Operations                               int      `json:"Operations"`
	OperationsItem                           int      `json:"OperationsItem"`
	OperationID                              int      `json:"OperationID"`
	ConfirmationCountingID                   int      `json:"ConfirmationCountingID"`
	OperationPlannedQuantityInBaseUnit       float32  `json:"OperationPlannedQuantityInBaseUnit"`
	OperationPlannedQuantityInProductionUnit float32  `json:"OperationPlannedQuantityInProductionUnit"`
	OperationPlannedQuantityInOperationUnit  float32  `json:"OperationPlannedQuantityInOperationUnit"`
	ProductBaseUnit                          string   `json:"ProductBaseUnit"`
	ProductProductionUnit                    string   `json:"ProductProductionUnit"`
	ProductOperationUnit                     string   `json:"ProductOperationUnit"`
	OperationPlannedScrapInPercent           *float32 `json:"OperationPlannedScrapInPercent"`
	ConfirmationEntryDate                    *string  `json:"ConfirmationEntryDate"`
	ConfirmationEntryTime                    *string  `json:"ConfirmationEntryTime"`
	ConfirmationText                         *string  `json:"ConfirmationText"`
	IsFinalConfirmation                      *bool	  `json:"IsFinalConfirmation"`
	WorkCenter                               int      `json:"WorkCenter"`
	EmployeeIDWhoConfirmed                   int      `json:"EmployeeIDWhoConfirmed"`
	ConfirmedExecutionStartDate              *string  `json:"ConfirmedExecutionStartDate"`
	ConfirmedExecutionStartTime              *string  `json:"ConfirmedExecutionStartTime"`
	ConfirmedSetupStartDate                  *string  `json:"ConfirmedSetupStartDate"`
	ConfirmedSetupStartTime                  *string  `json:"ConfirmedSetupStartTime"`
	ConfirmedProcessingStartDate             *string  `json:"ConfirmedProcessingStartDate"`
	ConfirmedProcessingStartTime             *string  `json:"ConfirmedProcessingStartTime"`
	ConfirmedExecutionEndDate                *string  `json:"ConfirmedExecutionEndDate"`
	ConfirmedExecutionEndTime                *string  `json:"ConfirmedExecutionEndTime"`
	ConfirmedSetupEndDate                    *string  `json:"ConfirmedSetupEndDate"`
	ConfirmedSetupEndTime                    *string  `json:"ConfirmedSetupEndTime"`
	ConfirmedProcessingEndDate               *string  `json:"ConfirmedProcessingEndDate"`
	ConfirmedProcessingEndTime               *string  `json:"ConfirmedProcessingEndTime"`
	ConfirmedWaitDuration                    *float32 `json:"ConfirmedWaitDuration"`
	WaitDurationUnit                         *string  `json:"WaitDurationUnit"`
	ConfirmedQueueDuration                   *float32 `json:"ConfirmedQueueDuration"`
	QueueDurationUnit                        *string  `json:"QueueDurationUnit"`
	ConfirmedMoveDuration                    *float32 `json:"ConfirmedMoveDuration"`
	MoveDurationUnit                         *string  `json:"MoveDurationUnit"`
	ConfirmedYieldQuantity                   *float32 `json:"ConfirmedYieldQuantity"`
	ConfirmedScrapQuantity                   *float32 `json:"ConfirmedScrapQuantity"`
	OperationVarianceReason                  *string  `json:"OperationVarianceReason"`
	CreationDate                             string   `json:"CreationDate"`
	CreationTime                             string   `json:"CreationTime"`
	LastChangeDate                           string   `json:"LastChangeDate"`
	LastChangeTime                           string   `json:"LastChangeTime"`
	IsCancelled                              *bool    `json:"IsCancelled"`
}

func CreateProductionOrderConfirmationUpdatesRequestHeaderUpdates(
	requestPram *types.Request,
	input ProductionOrderConfirmationReq,
) ProductionOrderConfirmationReq {
	req := ProductionOrderConfirmationReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func ProductionOrderConfirmationRequestHeaderUpdates(
	requestPram *types.Request,
	productionOrderConfirmationHeader types.ProductionOrderConfirmationSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PRODUCTION_ORDER_CONFIRMATION_SRV"
	//aPIType := "updates"
	aPIType := "creates"

	var request ProductionOrderConfirmationReq

	request = CreateProductionOrderConfirmationUpdatesRequestHeaderUpdates(
		requestPram,
		ProductionOrderConfirmationReq{
			Header: Header{
				ProductionOrder:                          productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ProductionOrder,
				ProductionOrderItem:                      productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ProductionOrderItem,
				Operations:                               productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.Operations,
				OperationsItem:                           productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.OperationsItem,
				OperationID:                              productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.OperationID,
				ConfirmationCountingID:                   productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmationCountingID,
				OperationPlannedQuantityInBaseUnit:       productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.OperationPlannedQuantityInBaseUnit,
				OperationPlannedQuantityInProductionUnit: productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.OperationPlannedQuantityInProductionUnit,
				OperationPlannedQuantityInOperationUnit:  productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.OperationPlannedQuantityInOperationUnit,
				ProductBaseUnit:                          productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ProductBaseUnit,
				ProductProductionUnit:                    productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ProductProductionUnit,
				ProductOperationUnit:                     productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ProductOperationUnit,
				OperationPlannedScrapInPercent:           productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.OperationPlannedScrapInPercent,
				ConfirmationEntryDate:                    productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmationEntryDate,
				ConfirmationEntryTime:                    productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmationEntryTime,
				ConfirmationText:                         productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmationText,
				IsFinalConfirmation:                      productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.IsFinalConfirmation,
				WorkCenter:                               productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.WorkCenter,
				EmployeeIDWhoConfirmed:                   productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.EmployeeIDWhoConfirmed,
				ConfirmedExecutionStartDate:              productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedExecutionStartDate,
				ConfirmedExecutionStartTime:              productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedExecutionStartTime,
				ConfirmedSetupStartDate:                  productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedSetupStartDate,
				ConfirmedSetupStartTime:                  productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedSetupStartTime,
				ConfirmedProcessingStartDate:             productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedProcessingStartDate,
				ConfirmedProcessingStartTime:             productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedProcessingStartTime,
				ConfirmedExecutionEndDate:                productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedExecutionEndDate,
				ConfirmedExecutionEndTime:                productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedExecutionEndTime,
				ConfirmedSetupEndDate:                    productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedSetupEndDate,
				ConfirmedSetupEndTime:                    productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedSetupEndTime,
				ConfirmedProcessingEndDate:               productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedProcessingEndDate,
				ConfirmedProcessingEndTime:               productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedProcessingEndTime,
				ConfirmedWaitDuration:                    productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedWaitDuration,
				WaitDurationUnit:                         productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.WaitDurationUnit,
				ConfirmedQueueDuration:                   productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedQueueDuration,
				QueueDurationUnit:                        productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.QueueDurationUnit,
				ConfirmedMoveDuration:                    productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedMoveDuration,
				MoveDurationUnit:                         productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.MoveDurationUnit,
				ConfirmedYieldQuantity:                   productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedYieldQuantity,
				ConfirmedScrapQuantity:                   productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.ConfirmedScrapQuantity,
				OperationVarianceReason:                  productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.OperationVarianceReason,
				CreationDate:                             productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.CreationDate,
				CreationTime:                             productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.CreationTime,
				LastChangeDate:                           productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.LastChangeDate,
				LastChangeTime:                           productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.LastChangeTime,
				IsCancelled:                              productionOrderConfirmationHeader.ProductionOrderConfirmationHeader.IsCancelled,
			},
		},
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}
