package controllersArticleDocCreatesHeaderDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsArticleDoc "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/article-doc"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type ArticleDocCreatesHeaderDocController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *ArticleDocCreatesHeaderDocController) Post() {
	inputParameter := types.ArticleDocInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *ArticleDocCreatesHeaderDocController,
) ArticleDocCreatesHeaderDocController(
	requestPram *types.Request,
	input types.ArticleDocSDC,
) *apiOutputFormatter.ArticleDocSDC {
	responseJsonData := apiOutputFormatter.ArticleDocSDC{}
	responseBody := apiModuleRuntimesRequestsArticleDoc.ArticleDocCreatesRequestHeaderDoc(
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
		controller.CustomLogger.Error("ArticleDocCreatesHeaderDoc Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *ArticleDocCreatesHeaderDocController,
) request(
	input types.ArticleDocSDC,
) {
	var response *apiOutputFormatter.ArticleDocSDC

	response = controller.ArticleDocCreatesHeaderDocController(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
