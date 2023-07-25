package apiModuleRuntimesRequestsBillOfMaterial

import (
	apiInputReader "data-platform-request-updates-manager-rmq-kube/api-input-reader"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego"
)

type BillOfMaterialReq struct {
	Header   Header   `json:"BillOfMaterial"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	BillOfMaterial                          int     `json:"BillOfMaterial"`
	ProductStandardQuantityInBaseUnit       float32 `json:"ProductStandardQuantityInBaseUnit"`
	ProductStandardQuantityInDeliveryUnit   float32 `json:"ProductStandardQuantityInDeliveryUnit"`
	ProductStandardQuantityInProductionUnit float32 `json:"ProductStandardQuantityInProductionUnit"`
	BillOfMaterialHeaderText                *string `json:"BillOfMaterialHeaderText"`
	ValidityStartDate                       *string `json:"ValidityStartDate"`
	ValidityEndDate                         *string `json:"ValidityEndDate"`
	Item                                    []Item  `json:"Item"`
}

type Item struct {
	BillOfMaterial                                 int      `json:"BillOfMaterial"`
	BillOfMaterialItem                             int      `json:"BillOfMaterialItem"`
	ComponentProductStandardQuantityInBaseUnit     float32  `json:"ComponentProductStandardQuantityInBaseUnit"`
	ComponentProductStandardQuantityInDeliveryUnit float32  `json:"ComponentProductStandardQuantityInDeliveryUnit"`
	ComponentProductStandardScrapInPercent         *float32 `json:"ComponentProductStandardScrapInPercent"`
	IsMarkedForBackflush                           *bool    `json:"IsMarkedForBackflush"`
	BillOfMaterialItemText                         *string  `json:"BillOfMaterialItemText"`
	ValidityStartDate                              *string  `json:"ValidityStartDate"`
	ValidityEndDate                                *string  `json:"ValidityEndDate"`
}

func CreateBillOfMaterialUpdatesRequestHeaderUpdates(
	requestPram *apiInputReader.Request,
	input BillOfMaterialReq,
) BillOfMaterialReq {
	req := BillOfMaterialReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func CreateBillOfMaterialUpdatesRequestItemUpdates(
	requestPram *apiInputReader.Request,
	input BillOfMaterialReq,
) BillOfMaterialReq {
	req := BillOfMaterialReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Item",
		},
	}
	return req
}

func BillOfMaterialRequestHeaderUpdates(
	requestPram *apiInputReader.Request,
	billOfMaterialHeader apiInputReader.BillOfMaterialSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BILL_OF_MATERIAL_SRV"
	aPIType := "updates"

	var request BillOfMaterialReq

	request = CreateBillOfMaterialUpdatesRequestHeaderUpdates(
		requestPram,
		BillOfMaterialReq{
			Header: Header{
				BillOfMaterial:                          billOfMaterialHeader.BillOfMaterialHeader.BillOfMaterial,
				ProductStandardQuantityInBaseUnit:       *billOfMaterialHeader.BillOfMaterialHeader.ProductStandardQuantityInBaseUnit,
				ProductStandardQuantityInDeliveryUnit:   *billOfMaterialHeader.BillOfMaterialHeader.ProductStandardQuantityInDeliveryUnit,
				ProductStandardQuantityInProductionUnit: *billOfMaterialHeader.BillOfMaterialHeader.ProductStandardQuantityInProductionUnit,
				BillOfMaterialHeaderText:                billOfMaterialHeader.BillOfMaterialHeader.BillOfMaterialHeaderText,
				ValidityStartDate:                       billOfMaterialHeader.BillOfMaterialHeader.ValidityStartDate,
				ValidityEndDate:                         billOfMaterialHeader.BillOfMaterialHeader.ValidityEndDate,
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

func BillOfMaterialRequestItemUpdates(
	requestPram *apiInputReader.Request,
	billOfMaterialHeader apiInputReader.BillOfMaterialSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BILL_OF_MATERIAL_SRV"
	aPIType := "updates"

	var request BillOfMaterialReq

	var items []Item

	items = append(items, Item{
		BillOfMaterial:     *billOfMaterialHeader.BillOfMaterialHeader.BillOfMaterialItem[0].BillOfMaterial,
		BillOfMaterialItem: *billOfMaterialHeader.BillOfMaterialHeader.BillOfMaterialItem[0].BillOfMaterialItem,
		ComponentProductStandardQuantityInBaseUnit: *billOfMaterialHeader.BillOfMaterialHeader.BillOfMaterialItem[0].ComponentProductStandardQuantityInBaseUnit,
		//ComponentProductStandardQuantityInDeliveryUnit: *billOfMaterialHeader.BillOfMaterialHeader.BillOfMaterialItem[0].ComponentProductStandardQuantityInDeliveryUnit,
		ComponentProductStandardScrapInPercent: billOfMaterialHeader.BillOfMaterialHeader.BillOfMaterialItem[0].ComponentProductStandardScrapInPercent,
		IsMarkedForBackflush:                   billOfMaterialHeader.BillOfMaterialHeader.BillOfMaterialItem[0].IsMarkedForBackflush,
		BillOfMaterialItemText:                 billOfMaterialHeader.BillOfMaterialHeader.BillOfMaterialItem[0].BillOfMaterialItemText,
		ValidityStartDate:                      billOfMaterialHeader.BillOfMaterialHeader.BillOfMaterialItem[0].ValidityStartDate,
		ValidityEndDate:                        billOfMaterialHeader.BillOfMaterialHeader.BillOfMaterialItem[0].ValidityEndDate,
	})

	request = CreateBillOfMaterialUpdatesRequestItemUpdates(
		requestPram,
		BillOfMaterialReq{
			Header: Header{
				BillOfMaterial: billOfMaterialHeader.BillOfMaterialHeader.BillOfMaterial,
				Item:           items,
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
