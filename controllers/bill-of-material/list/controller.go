package controllersBillOfMaterialList

import (
	apiInputReader "data-platform-request-updates-manager-rmq-kube/api-input-reader"
	apiModuleRuntimesRequestsBillOfMaterial "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/bill-of-material"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type BillOfMaterialListController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *apiInputReader.Request
	CustomLogger *logger.Logger
}

const (
	OwnerProductionPlantBusinessPartner = "ownerProductionPlantBusinessPartner"
)

func (controller *BillOfMaterialListController) Post() {
	billOfMaterial := apiInputReader.BillOfMaterialInputRead(&controller.Controller)

	billOfMaterialHeader := billOfMaterial

	controller.request(billOfMaterialHeader)
}

func (
	controller *BillOfMaterialListController,
) BillOfMaterialRequestUpdates(
	requestPram *apiInputReader.Request,
	input apiInputReader.BillOfMaterialSDC,
) *apiOutputFormatter.BillOfMaterialSDC {
	responseJsonData := apiOutputFormatter.BillOfMaterialSDC{}
	responseBody := apiModuleRuntimesRequestsBillOfMaterial.BillOfMaterialRequestHeaderUpdates(
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
		controller.CustomLogger.Error("BillOfMaterialRequestUpdates Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BillOfMaterialListController,
) request(
	input apiInputReader.BillOfMaterialSDC,
) {
	var bRes *apiOutputFormatter.BillOfMaterialSDC

	bRes = controller.BillOfMaterialRequestUpdates(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &bRes
	controller.ServeJSON()
}
