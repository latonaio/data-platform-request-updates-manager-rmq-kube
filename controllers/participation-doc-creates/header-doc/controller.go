package controllersParticipationDocCreatesHeaderDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsParticipationDoc "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/participation-doc"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ParticipationDocCreatesHeaderDocController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *ParticipationDocCreatesHeaderDocController) Post() {
	inputParameter := types.ParticipationDocInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *ParticipationDocCreatesHeaderDocController,
) ParticipationDocCreatesHeaderDocController(
	requestPram *types.Request,
	input types.ParticipationDocSDC,
) *apiOutputFormatter.ParticipationDocSDC {
	responseJsonData := apiOutputFormatter.ParticipationDocSDC{}
	responseBody := apiModuleRuntimesRequestsParticipationDoc.ParticipationDocCreatesRequestHeaderDoc(
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
		controller.CustomLogger.Error("ParticipationDocCreatesHeaderDoc Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ParticipationDocCreatesHeaderDocController,
) request(
	input types.ParticipationDocSDC,
) {
	var response *apiOutputFormatter.ParticipationDocSDC

	response = controller.ParticipationDocCreatesHeaderDocController(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
