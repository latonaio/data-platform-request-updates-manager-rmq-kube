package controllersBusinessPartnerCreatesAll

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type BusinessPartnerCreatesAllController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *BusinessPartnerCreatesAllController) Post() {
	inputParameter := types.BusinessPartnerInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *BusinessPartnerCreatesAllController,
) BusinessPartnerCreatesRequestAll(
	requestPram *types.Request,
	input types.BusinessPartnerSDC,
) *apiOutputFormatter.BusinessPartnerSDC {
	responseJsonData := apiOutputFormatter.BusinessPartnerSDC{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerCreatesRequestAll(
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
		controller.CustomLogger.Error("BusinessPartnerCreatesRequestAll Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BusinessPartnerCreatesAllController,
) request(
	input types.BusinessPartnerSDC,
) {
	var response *apiOutputFormatter.BusinessPartnerSDC

	response = controller.BusinessPartnerCreatesRequestAll(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
