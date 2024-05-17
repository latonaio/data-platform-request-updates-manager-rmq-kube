package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type ProductionOrderConfirmationSDC struct {
	ConnectionKey                     string                            `json:"connection_key"`
	Result                            bool                              `json:"result"`
	RedisKey                          string                            `json:"redis_key"`
	Filepath                          string                            `json:"filepath"`
	APIStatusCode                     int                               `json:"api_status_code"`
	RuntimeSessionID                  string                            `json:"runtime_session_id"`
	BusinessPartner                   *int                              `json:"business_partner"`
	ServiceLabel                      string                            `json:"service_label"`
	APIType                           string                            `json:"api_type"`
	ProductionOrderConfirmationHeader ProductionOrderConfirmationHeader `json:"ProductionOrderConfirmation"`
	APISchema                         string                            `json:"api_schema"`
	Accepter                          []string                          `json:"accepter"`
	Deleted                           bool                              `json:"deleted"`
}

type ProductionOrderConfirmationHeader struct {
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

func ProductionOrderConfirmationInputRead(
	controller *beego.Controller,
) ProductionOrderConfirmationSDC {
	var requestBody ProductionOrderConfirmationSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
