package apiModuleRuntimesRequestsInspectionLot

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type InspectionLotReq struct {
	Header   Header   `json:"InspectionLot"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	InspectionLot                  int     `json:"InspectionLot"`
	InspectionPlan                 int     `json:"InspectionPlan"`
	InspectionPlantBusinessPartner int     `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                string  `json:"InspectionPlant"`
	Product                        string  `json:"Product"`
	ProductSpecification           *string `json:"ProductSpecification"`
	InspectionSpecification        *string `json:"InspectionSpecification"`
	ProductionOrder                *int    `json:"ProductionOrder"`
	ProductionOrderItem            *int    `json:"ProductionOrderItem"`
	InspectionLotHeaderText        *string `json:"InspectionLotHeaderText"`
	Inspection                     []Inspection  `json:"Inspection"`
}

type Inspection struct {
	InspectionLot	                            int	        `json:"InspectionLot"`
    Inspection	                                int	        `json:"Inspection"`
    InspectionType                            	string	    `json:"InspectionType"`
    InspectionTypeValueUnit	                    *string	    `json:"InspectionTypeValueUnit"`
    InspectionTypePlannedValue	                *float32	`json:"InspectionTypePlannedValue"`
    InspectionTypeCertificateType	            *string	    `json:"InspectionTypeCertificateType"`
    InspectionTypeCertificateValueInText	    *string	    `json:"InspectionTypeCertificateValueInText"`
    InspectionTypeCertificateValueInQuantity	*float32	`json:"InspectionTypeCertificateValueInQuantity"`
    InspectionLotInspectionText	                *string	    `json:"InspectionLotInspectionText"`
}

func CreateInspectionLotUpdatesRequestHeaderUpdates(
	requestPram *types.Request,
	input InspectionLotReq,
) InspectionLotReq {
	req := InspectionLotReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreateInspectionLotUpdatesRequestInspectionUpdates(
	requestPram *types.Request,
	input InspectionLotReq,
) InspectionLotReq {
	req := InspectionLotReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Inspection",
		},
	}
	return req
}

func InspectionLotRequestHeaderUpdates(
	requestPram *types.Request,
	inspectionLotHeader types.InspectionLotSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_INSPECTION_LOT_SRV"
	aPIType := "updates"

	var request InspectionLotReq

	request = CreateInspectionLotUpdatesRequestHeaderUpdates(
		requestPram,
		InspectionLotReq{
			Header: Header{
				InspectionLot:                           inspectionLotHeader.InspectionLotHeader.InspectionLot,
				InspectionPlan:                          inspectionLotHeader.InspectionLotHeader.InspectionPlan,
				InspectionPlantBusinessPartner:          inspectionLotHeader.InspectionLotHeader.InspectionPlantBusinessPartner,
				InspectionPlant:                         inspectionLotHeader.InspectionLotHeader.InspectionPlant,
				Product:                                 inspectionLotHeader.InspectionLotHeader.Product,
				ProductSpecification:                    inspectionLotHeader.InspectionLotHeader.ProductSpecification,
				InspectionSpecification:                 inspectionLotHeader.InspectionLotHeader.InspectionSpecification,
				ProductionOrder:                         inspectionLotHeader.InspectionLotHeader.ProductionOrder,
				ProductionOrderItem:                     inspectionLotHeader.InspectionLotHeader.ProductionOrderItem,
				InspectionLotHeaderText:                 inspectionLotHeader.InspectionLotHeader.InspectionLotHeaderText,
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

func InspectionLotRequestInspectionUpdates(
	requestPram *types.Request,
	inspectionLotHeader types.InspectionLotSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_INSPECTION_LOT_SRV"
	aPIType := "updates"

	var request InspectionLotReq

	var inspections []Inspection

	inspections = append(inspections, Inspection{
		InspectionLot:     							inspectionLotHeader.InspectionLotHeader.InspectionLotInspection[0].InspectionLot,
		Inspection:		 							inspectionLotHeader.InspectionLotHeader.InspectionLotInspection[0].Inspection,
		InspectionType: 							inspectionLotHeader.InspectionLotHeader.InspectionLotInspection[0].InspectionType,
		InspectionTypeValueUnit: 					inspectionLotHeader.InspectionLotHeader.InspectionLotInspection[0].InspectionTypeValueUnit,
		InspectionTypePlannedValue:                 inspectionLotHeader.InspectionLotHeader.InspectionLotInspection[0].InspectionTypePlannedValue,
		InspectionTypeCertificateType:              inspectionLotHeader.InspectionLotHeader.InspectionLotInspection[0].InspectionTypeCertificateType,
		InspectionTypeCertificateValueInText:       inspectionLotHeader.InspectionLotHeader.InspectionLotInspection[0].InspectionTypeCertificateValueInText,
		InspectionTypeCertificateValueInQuantity:   inspectionLotHeader.InspectionLotHeader.InspectionLotInspection[0].InspectionTypeCertificateValueInQuantity,
		InspectionLotInspectionText:   				inspectionLotHeader.InspectionLotHeader.InspectionLotInspection[0].InspectionLotInspectionText,
	})

	request = CreateInspectionLotUpdatesRequestInspectionUpdates(
		requestPram,
		InspectionLotReq{
			Header: Header{
				InspectionLot: inspectionLotHeader.InspectionLotHeader.InspectionLot,
				Inspection:    inspections,
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
