package controllersBusinessPartnerDocCreatesGeneralDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsBusinessPartnerDoc "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/business-partner-doc"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type BusinessPartnerDocCreatesGeneralDocController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *BusinessPartnerDocCreatesGeneralDocController) Post() {
	inputParameter := types.BusinessPartnerDocInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *BusinessPartnerDocCreatesGeneralDocController,
) BusinessPartnerDocCreatesGeneralDocController(
	requestPram *types.Request,
	input types.BusinessPartnerDocSDC,
) *apiOutputFormatter.BusinessPartnerDocSDC {
	responseJsonData := apiOutputFormatter.BusinessPartnerDocSDC{}
	responseBody := apiModuleRuntimesRequestsBusinessPartnerDoc.BusinessPartnerDocCreatesRequestGeneralDoc(
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
		controller.CustomLogger.Error("BusinessPartnerDocCreatesGeneralDoc Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BusinessPartnerDocCreatesGeneralDocController,
) request(
	input types.BusinessPartnerDocSDC,
) {
	var response *apiOutputFormatter.BusinessPartnerDocSDC

	response = controller.BusinessPartnerDocCreatesGeneralDocController(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
