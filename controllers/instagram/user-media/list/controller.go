package controllersInstagramUserMediaList

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsInstagramUserMedia "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/instagram-user-media"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type InstagramUserMediaListController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *InstagramUserMediaListController) Post() {
	inputParameter := types.InstagramUserMediaInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *InstagramUserMediaListController,
) InstagramUserMediaRequestMedias(
	requestPram *types.Request,
	input types.InstagramUserMediaSDC,
) *apiOutputFormatter.InstagramUserMediaSDC {
	responseJsonData := apiOutputFormatter.InstagramUserMediaSDC{}
	responseBody := apiModuleRuntimesRequestsInstagramUserMedia.InstagramUserMediaRequestsMedias(
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
		controller.CustomLogger.Error("InstagramUserMediaRequestMedias Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *InstagramUserMediaListController,
) request(
	input types.InstagramUserMediaSDC,
) {
	var response *apiOutputFormatter.InstagramUserMediaSDC

	response = controller.InstagramUserMediaRequestMedias(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
