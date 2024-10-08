package controllersQRCodeGeneratesParticipation

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsQRCode "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/qr-code"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type QRCodeGeneratesParticipationController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *QRCodeGeneratesParticipationController) Post() {
	inputParameter := types.QRCodeInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *QRCodeGeneratesParticipationController,
) QRCodeGeneratesParticipationController(
	requestPram *types.Request,
	input types.QRCodeSDC,
) *apiOutputFormatter.QRCodeSDC {
	responseJsonData := apiOutputFormatter.QRCodeSDC{}
	responseBody := apiModuleRuntimesRequestsQRCode.QRCodeGeneratesRequestParticipation(
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
		controller.CustomLogger.Error("QRCodeGeneratesParticipation Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *QRCodeGeneratesParticipationController,
) request(
	input types.QRCodeSDC,
) {
	var response *apiOutputFormatter.QRCodeSDC

	response = controller.QRCodeGeneratesParticipationController(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
