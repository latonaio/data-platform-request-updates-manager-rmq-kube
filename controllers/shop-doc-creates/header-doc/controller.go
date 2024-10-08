package controllersShopDocCreatesHeaderDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsShopDoc "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/shop-doc"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ShopDocCreatesHeaderDocController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *ShopDocCreatesHeaderDocController) Post() {
	inputParameter := types.ShopDocInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *ShopDocCreatesHeaderDocController,
) ShopDocCreatesHeaderDocController(
	requestPram *types.Request,
	input types.ShopDocSDC,
) *apiOutputFormatter.ShopDocSDC {
	responseJsonData := apiOutputFormatter.ShopDocSDC{}
	responseBody := apiModuleRuntimesRequestsShopDoc.ShopDocCreatesRequestHeaderDoc(
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
		controller.CustomLogger.Error("ShopDocCreatesHeaderDoc Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ShopDocCreatesHeaderDocController,
) request(
	input types.ShopDocSDC,
) {
	var response *apiOutputFormatter.ShopDocSDC

	response = controller.ShopDocCreatesHeaderDocController(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
