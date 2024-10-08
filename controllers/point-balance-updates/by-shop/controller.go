package controllersPointBalanceUpdatesByShop

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsPointBalance "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/point-balance"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type PointBalanceUpdatesByShopController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *PointBalanceUpdatesByShopController) Post() {
	inputParameter := types.PointBalanceInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *PointBalanceUpdatesByShopController,
) PointBalanceUpdatesRequestByShop(
	requestPram *types.Request,
	input types.PointBalanceSDC,
) *apiOutputFormatter.PointBalanceSDC {
	responseJsonData := apiOutputFormatter.PointBalanceSDC{}
	responseBody := apiModuleRuntimesRequestsPointBalance.PointBalanceUpdatesRequestByShop(
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
		controller.CustomLogger.Error("PointBalanceUpdatesRequestByShop Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointBalanceUpdatesByShopController,
) request(
	input types.PointBalanceSDC,
) {
	var response *apiOutputFormatter.PointBalanceSDC

	response = controller.PointBalanceUpdatesRequestByShop(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
