package controllersSiteUpdatesCounter

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsSite "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/site"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type SiteUpdatesCounterController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *SiteUpdatesCounterController) Post() {
	inputParameter := types.SiteInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *SiteUpdatesCounterController,
) SiteUpdatesRequestCounter(
	requestPram *types.Request,
	input types.SiteSDC,
) *apiOutputFormatter.SiteSDC {
	responseJsonData := apiOutputFormatter.SiteSDC{}
	responseBody := apiModuleRuntimesRequestsSite.SiteUpdatesRequestCounter(
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
		controller.CustomLogger.Error("SiteUpdatesRequestCounter Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteUpdatesCounterController,
) request(
	input types.SiteSDC,
) {
	var response *apiOutputFormatter.SiteSDC

	response = controller.SiteUpdatesRequestCounter(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
