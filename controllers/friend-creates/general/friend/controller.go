package controllersFriendCreatesGeneralFriend

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsFriend "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/friend"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type FriendCreatesGeneralFriendController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *FriendCreatesGeneralFriendController) Post() {
	inputParameter := types.FriendInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *FriendCreatesGeneralFriendController,
) FriendCreatesRequestGeneralFriend(
	requestPram *types.Request,
	input types.FriendSDC,
) *apiOutputFormatter.FriendSDC {
	responseJsonData := apiOutputFormatter.FriendSDC{}
	responseBody := apiModuleRuntimesRequestsFriend.FriendCreatesRequestGeneralFriend(
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
		controller.CustomLogger.Error("FriendCreatesRequestGeneralFriend Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *FriendCreatesGeneralFriendController,
) request(
	input types.FriendSDC,
) {
	var response *apiOutputFormatter.FriendSDC

	response = controller.FriendCreatesRequestGeneralFriend(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
