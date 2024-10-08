package controllersPointTransactionCancelsHeader

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsPointTransaction "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/point-transaction"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type PointTransactionCancelsHeaderController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *PointTransactionCancelsHeaderController) Post() {
	inputParameter := types.PointTransactionInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *PointTransactionCancelsHeaderController,
) PointTransactionCancelsRequestHeader(
	requestPram *types.Request,
	input types.PointTransactionSDC,
) *apiOutputFormatter.PointTransactionSDC {
	responseJsonData := apiOutputFormatter.PointTransactionSDC{}
	responseBody := apiModuleRuntimesRequestsPointTransaction.PointTransactionCancelsRequestHeader(
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
		controller.CustomLogger.Error("PointTransactionCancelsRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PointTransactionCancelsHeaderController,
) request(
	input types.PointTransactionSDC,
) {
	var response *apiOutputFormatter.PointTransactionSDC

	response = controller.PointTransactionCancelsRequestHeader(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
