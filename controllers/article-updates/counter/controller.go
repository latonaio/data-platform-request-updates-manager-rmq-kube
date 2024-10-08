package controllersArticleUpdatesCounter

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsArticle "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/article"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ArticleUpdatesCounterController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *ArticleUpdatesCounterController) Post() {
	inputParameter := types.ArticleInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *ArticleUpdatesCounterController,
) ArticleUpdatesRequestCounter(
	requestPram *types.Request,
	input types.ArticleSDC,
) *apiOutputFormatter.ArticleSDC {
	responseJsonData := apiOutputFormatter.ArticleSDC{}
	responseBody := apiModuleRuntimesRequestsArticle.ArticleUpdatesRequestCounter(
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
		controller.CustomLogger.Error("ArticleUpdatesRequestCounter Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleUpdatesCounterController,
) request(
	input types.ArticleSDC,
) {
	var response *apiOutputFormatter.ArticleSDC

	response = controller.ArticleUpdatesRequestCounter(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
