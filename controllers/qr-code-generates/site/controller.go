package controllersQRCodeGeneratesSite

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsQRCode "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/qr-code"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type QRCodeGeneratesSiteController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *QRCodeGeneratesSiteController) Post() {
	inputParameter := types.QRCodeInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *QRCodeGeneratesSiteController,
) QRCodeGeneratesSiteController(
	requestPram *types.Request,
	input types.QRCodeSDC,
) *apiOutputFormatter.QRCodeSDC {
	responseJsonData := apiOutputFormatter.QRCodeSDC{}
	responseBody := apiModuleRuntimesRequestsQRCode.QRCodeGeneratesRequestSite(
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
		controller.CustomLogger.Error("QRCodeGeneratesSite Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *QRCodeGeneratesSiteController,
) request(
	input types.QRCodeSDC,
) {
	var response *apiOutputFormatter.QRCodeSDC

	response = controller.QRCodeGeneratesSiteController(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
