package controllersBusinessPartnerUpdatesPersonGoogleAccountAuth

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsBusinessPartner "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/business-partner"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type BusinessPartnerUpdatesPersonGoogleAccountAuthController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *BusinessPartnerUpdatesPersonGoogleAccountAuthController) Post() {
	inputParameter := types.BusinessPartnerInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *BusinessPartnerUpdatesPersonGoogleAccountAuthController,
) BusinessPartnerUpdatesRequestPersonGoogleAccountAuth(
	requestPram *types.Request,
	input types.BusinessPartnerSDC,
) *apiOutputFormatter.BusinessPartnerSDC {
	responseJsonData := apiOutputFormatter.BusinessPartnerSDC{}
	responseBody := apiModuleRuntimesRequestsBusinessPartner.BusinessPartnerUpdatesRequestPersonGoogleAccountAuth(
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
		controller.CustomLogger.Error("BusinessPartnerUpdatesRequestPersonGoogleAccountAuth Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *BusinessPartnerUpdatesPersonGoogleAccountAuthController,
) request(
	input types.BusinessPartnerSDC,
) {
	var response *apiOutputFormatter.BusinessPartnerSDC

	response = controller.BusinessPartnerUpdatesRequestPersonGoogleAccountAuth(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
