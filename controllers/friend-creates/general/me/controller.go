package controllersFriendCreatesGeneralMe

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsFriend "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/friend"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type FriendCreatesGeneralMeController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *FriendCreatesGeneralMeController) Post() {
	inputParameter := types.FriendInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *FriendCreatesGeneralMeController,
) FriendCreatesRequestGeneralMe(
	requestPram *types.Request,
	input types.FriendSDC,
) *apiOutputFormatter.FriendSDC {
	responseJsonData := apiOutputFormatter.FriendSDC{}
	responseBody := apiModuleRuntimesRequestsFriend.FriendCreatesRequestGeneralMe(
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
		controller.CustomLogger.Error("FriendCreatesRequestGeneralMe Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *FriendCreatesGeneralMeController,
) request(
	input types.FriendSDC,
) {
	var response *apiOutputFormatter.FriendSDC

	response = controller.FriendCreatesRequestGeneralMe(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
