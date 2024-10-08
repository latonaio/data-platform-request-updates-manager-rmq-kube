package controllersShopUpdatesHeader

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsShop "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/shop"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ShopUpdatesHeaderController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *ShopUpdatesHeaderController) Post() {
	inputParameter := types.ShopInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *ShopUpdatesHeaderController,
) ShopUpdatesRequestHeader(
	requestPram *types.Request,
	input types.ShopSDC,
) *apiOutputFormatter.ShopSDC {
	responseJsonData := apiOutputFormatter.ShopSDC{}
	responseBody := apiModuleRuntimesRequestsShop.ShopUpdatesRequestHeader(
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
		controller.CustomLogger.Error("ShopUpdatesRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ShopUpdatesHeaderController,
) request(
	input types.ShopSDC,
) {
	var response *apiOutputFormatter.ShopSDC

	response = controller.ShopUpdatesRequestHeader(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
