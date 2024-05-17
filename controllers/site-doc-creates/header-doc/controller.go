package controllersSiteDocCreatesHeaderDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsSiteDoc "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/site-doc"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type SiteDocCreatesHeaderDocController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *SiteDocCreatesHeaderDocController) Post() {
	inputParameter := types.SiteDocInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *SiteDocCreatesHeaderDocController,
) SiteDocCreatesHeaderDocController(
	requestPram *types.Request,
	input types.SiteDocSDC,
) *apiOutputFormatter.SiteDocSDC {
	responseJsonData := apiOutputFormatter.SiteDocSDC{}
	responseBody := apiModuleRuntimesRequestsSiteDoc.SiteDocCreatesRequestHeaderDoc(
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
		controller.CustomLogger.Error("SiteDocCreatesHeaderDoc Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *SiteDocCreatesHeaderDocController,
) request(
	input types.SiteDocSDC,
) {
	var response *apiOutputFormatter.SiteDocSDC

	response = controller.SiteDocCreatesHeaderDocController(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
