package controllersSiteCreatesAll

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsSite "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/site"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type SiteCreatesAllController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *SiteCreatesAllController) Post() {
	inputParameter := types.SiteInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *SiteCreatesAllController,
) SiteCreatesRequestAll(
	requestPram *types.Request,
	input types.SiteSDC,
) *apiOutputFormatter.SiteSDC {
	responseJsonData := apiOutputFormatter.SiteSDC{}
	responseBody := apiModuleRuntimesRequestsSite.SiteCreatesRequestAll(
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
		controller.CustomLogger.Error("SiteCreatesRequestAll Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteCreatesAllController,
) request(
	input types.SiteSDC,
) {
	var response *apiOutputFormatter.SiteSDC

	response = controller.SiteCreatesRequestAll(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
