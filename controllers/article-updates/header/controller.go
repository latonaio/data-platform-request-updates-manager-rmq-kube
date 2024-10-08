package controllersArticleUpdatesHeader

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsArticle "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/article"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ArticleUpdatesHeaderController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *ArticleUpdatesHeaderController) Post() {
	inputParameter := types.ArticleInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *ArticleUpdatesHeaderController,
) ArticleUpdatesRequestHeader(
	requestPram *types.Request,
	input types.ArticleSDC,
) *apiOutputFormatter.ArticleSDC {
	responseJsonData := apiOutputFormatter.ArticleSDC{}
	responseBody := apiModuleRuntimesRequestsArticle.ArticleUpdatesRequestHeader(
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
		controller.CustomLogger.Error("ArticleUpdatesRequestHeader Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleUpdatesHeaderController,
) request(
	input types.ArticleSDC,
) {
	var response *apiOutputFormatter.ArticleSDC

	response = controller.ArticleUpdatesRequestHeader(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
