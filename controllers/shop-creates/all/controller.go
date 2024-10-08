package controllersShopCreatesAll

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsShop "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/shop"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ShopCreatesAllController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *ShopCreatesAllController) Post() {
	inputParameter := types.ShopInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *ShopCreatesAllController,
) ShopCreatesRequestAll(
	requestPram *types.Request,
	input types.ShopSDC,
) *apiOutputFormatter.ShopSDC {
	responseJsonData := apiOutputFormatter.ShopSDC{}
	responseBody := apiModuleRuntimesRequestsShop.ShopCreatesRequestAll(
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
		controller.CustomLogger.Error("ShopCreatesRequestAll Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ShopCreatesAllController,
) request(
	input types.ShopSDC,
) {
	var response *apiOutputFormatter.ShopSDC

	response = controller.ShopCreatesRequestAll(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
