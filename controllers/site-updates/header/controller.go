package controllersSiteUpdatesHeader

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsSite "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/site"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type SiteUpdatesHeaderController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *SiteUpdatesHeaderController) Post() {
	inputParameter := types.SiteInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *SiteUpdatesHeaderController,
) SiteUpdatesRequestHeader(
	requestPram *types.Request,
	input types.SiteSDC,
) *apiOutputFormatter.SiteSDC {
	responseJsonData := apiOutputFormatter.SiteSDC{}
	responseBody := apiModuleRuntimesRequestsSite.SiteUpdatesRequestHeader(
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
		controller.CustomLogger.Error("SiteUpdatesRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteUpdatesHeaderController,
) request(
	input types.SiteSDC,
) {
	var response *apiOutputFormatter.SiteSDC

	response = controller.SiteUpdatesRequestHeader(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
