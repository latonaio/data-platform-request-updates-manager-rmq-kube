package types

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type BillOfMaterialSDC struct {
	ConnectionKey        string               `json:"connection_key"`
	Result               bool                 `json:"result"`
	RedisKey             string               `json:"redis_key"`
	Filepath             string               `json:"filepath"`
	APIStatusCode        int                  `json:"api_status_code"`
	RuntimeSessionID     string               `json:"runtime_session_id"`
	BusinessPartner      *int                 `json:"business_partner"`
	ServiceLabel         string               `json:"service_label"`
	APIType              string               `json:"api_type"`
	BillOfMaterialHeader BillOfMaterialHeader `json:"BillOfMaterial"`
	APISchema            string               `json:"api_schema"`
	Accepter             []string             `json:"accepter"`
	Deleted              bool                 `json:"deleted"`
}

type BillOfMaterialHeader struct {
	BillOfMaterial                          int                  `json:"BillOfMaterial"`
	ProductStandardQuantityInBaseUnit       *float32             `json:"ProductStandardQuantityInBaseUnit"`
	ProductStandardQuantityInDeliveryUnit   *float32             `json:"ProductStandardQuantityInDeliveryUnit"`
	ProductStandardQuantityInProductionUnit *float32             `json:"ProductStandardQuantityInProductionUnit"`
	BillOfMaterialHeaderText                *string              `json:"BillOfMaterialHeaderText"`
	ValidityStartDate                       *string              `json:"ValidityStartDate"`
	ValidityEndDate                         *string              `json:"ValidityEndDate"`
	BillOfMaterialItem                      []BillOfMaterialItem `json:"Item"`
}

type BillOfMaterialItem struct {
	BillOfMaterial                                 *int     `json:"BillOfMaterial"`
	BillOfMaterialItem                             *int     `json:"BillOfMaterialItem"`
	ComponentProductStandardQuantityInBaseUnit     *float32 `json:"ComponentProductStandardQuantityInBaseUnit"`
	ComponentProductStandardQuantityInDeliveryUnit *float32 `json:"ComponentProductStandardQuantityInDeliveryUnit"`
	ComponentProductStandardScrapInPercent         *float32 `json:"ComponentProductStandardScrapInPercent"`
	IsMarkedForBackflush                           *bool    `json:"IsMarkedForBackflush"`
	BillOfMaterialItemText                         *string  `json:"BillOfMaterialItemText"`
	ValidityStartDate                              *string  `json:"ValidityStartDate"`
	ValidityEndDate                                *string  `json:"ValidityEndDate"`
}

func BillOfMaterialInputRead(
	controller *beego.Controller,
) BillOfMaterialSDC {
	var requestBody BillOfMaterialSDC
	err := json.Unmarshal(controller.Ctx.Input.RequestBody, &requestBody)

	if err != nil {
		controller.Data["json"] = map[string]interface{}{
			"error": fmt.Sprintf("json decode error: %v", err),
		}
		controller.ServeJSON()
	}

	return requestBody
}
