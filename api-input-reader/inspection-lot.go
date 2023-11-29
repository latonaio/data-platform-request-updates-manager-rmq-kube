package apiInputReader

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type InspectionLotSDC struct {
	ConnectionKey        string               `json:"connection_key"`
	Result               bool                 `json:"result"`
	RedisKey             string               `json:"redis_key"`
	Filepath             string               `json:"filepath"`
	APIStatusCode        int                  `json:"api_status_code"`
	RuntimeSessionID     string               `json:"runtime_session_id"`
	BusinessPartner      *int                 `json:"business_partner"`
	ServiceLabel         string               `json:"service_label"`
	APIType              string               `json:"api_type"`
	InspectionLotHeader  InspectionLotHeader  `json:"InspectionLot"`
	APISchema            string               `json:"api_schema"`
	Accepter             []string             `json:"accepter"`
	Deleted              bool                 `json:"deleted"`
}

type InspectionLotHeader struct {
	InspectionLot                  int     `json:"InspectionLot"`
	InspectionPlan                 *int    `json:"InspectionPlan"`
	InspectionPlantBusinessPartner *int    `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                *string `json:"InspectionPlant"`
	Product                        *string `json:"Product"`
	ProductSpecification           *string `json:"ProductSpecification"`
	InspectionSpecification        *string `json:"InspectionSpecification"`
	ProductionOrder                *int    `json:"ProductionOrder"`
	ProductionOrderItem            *int    `json:"ProductionOrderItem"`
	InspectionLotHeaderText        *string `json:"InspectionLotHeaderText"`
	InspectionLotInspection        []InspectionLotInspection	`json:"Inspection"`
}

type InspectionLotInspection struct {
	InspectionLot	                            int	        `json:"InspectionLot"`
    Inspection	                                int	        `json:"Inspection"`
    InspectionType                            	*string	    `json:"InspectionType"`
    InspectionTypeValueUnit	                    *string	    `json:"InspectionTypeValueUnit"`
    InspectionTypePlannedValue	                *float32	`json:"InspectionTypePlannedValue"`
    InspectionTypeCertificateType	            *string	    `json:"InspectionTypeCertificateType"`
    InspectionTypeCertificateValueInText	    *string	    `json:"InspectionTypeCertificateValueInText"`
    InspectionTypeCertificateValueInQuantity	*float32	`json:"InspectionTypeCertificateValueInQuantity"`
    InspectionLotInspectionText	                *string	    `json:"InspectionLotInspectionText"`
}

func InspectionLotInputRead(
	controller *beego.Controller,
) InspectionLotSDC {
	var requestBody InspectionLotSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
