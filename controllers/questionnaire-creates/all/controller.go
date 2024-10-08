package controllersQuestionnaireCreatesAll

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsQuestionnaire "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/questionnaire"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type QuestionnaireCreatesAllController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *QuestionnaireCreatesAllController) Post() {
	input := types.QuestionnaireInputRead(&controller.Controller)
	controller.request(input)
}

func (
	controller *QuestionnaireCreatesAllController,
) QuestionnaireCreatesRequestAll(
	requestPram *types.Request,
	input types.QuestionnaireSDC,
) *apiOutputFormatter.QuestionnaireSDC {
	responseJsonData := apiOutputFormatter.QuestionnaireSDC{}
	responseBody := apiModuleRuntimesRequestsQuestionnaire.QuestionnaireCreatesRequestAll(
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
		controller.CustomLogger.Error("QuestionnaireCreatesRequestAll Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *QuestionnaireCreatesAllController,
) request(
	input types.QuestionnaireSDC,
) {
	var response *apiOutputFormatter.QuestionnaireSDC

	response = controller.QuestionnaireCreatesRequestAll(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
