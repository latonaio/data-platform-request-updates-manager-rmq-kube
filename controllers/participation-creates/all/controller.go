package controllersParticipationCreatesAll

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsParticipation "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/participation"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ParticipationCreatesAllController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *ParticipationCreatesAllController) Post() {
	inputParameter := types.ParticipationInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *ParticipationCreatesAllController,
) ParticipationCreatesRequestAll(
	requestPram *types.Request,
	input types.ParticipationSDC,
) *apiOutputFormatter.ParticipationSDC {
	responseJsonData := apiOutputFormatter.ParticipationSDC{}
	responseBody := apiModuleRuntimesRequestsParticipation.ParticipationCreatesRequestAll(
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
		controller.CustomLogger.Error("ParticipationCreatesRequestAll Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationCreatesAllController,
) request(
	input types.ParticipationSDC,
) {
	var response *apiOutputFormatter.ParticipationSDC

	response = controller.ParticipationCreatesRequestAll(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
