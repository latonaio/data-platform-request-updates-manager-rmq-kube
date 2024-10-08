package controllersPostCreatesHeadersWithInstagramMedia

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/formatter"
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsInstagramUserMedia "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/instagram-user-media"
	apiModuleRuntimesRequestsPost "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/post"
	apiModuleRuntimesRequestsPostDoc "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/post-doc"
	apiModuleRuntimesResponsesInstagramUserMedia "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-responses/instagram-user-media"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type PostCreatesHeadersWithInstagramMediaController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *PostCreatesHeadersWithInstagramMediaController) Post() {
	//postInputParameter := types.PostInputRead(&controller.Controller)
	controller.UserInfo = services.UserRequestParams(
		services.RequestWrapperController{
			Controller:   &controller.Controller,
			CustomLogger: controller.CustomLogger,
		},
	)
	instagramUserMediaInputParameter := types.InstagramUserMediaInputRead(&controller.Controller)
	controller.request(instagramUserMediaInputParameter)
}

func (
	controller *PostCreatesHeadersWithInstagramMediaController,
) PostCreatesRequestHeadersWithInstagramMedia(
	requestPram *types.Request,
	input types.PostSDC,
) *apiOutputFormatter.PostSDC {
	responseJsonData := apiOutputFormatter.PostSDC{}
	responseBody := apiModuleRuntimesRequestsPost.PostCreatesRequestHeadersWithInstagramMedia(
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
		controller.CustomLogger.Error("PostCreatesRequestHeadersWithInstagramMedia Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostCreatesHeadersWithInstagramMediaController,
) PostDocCreatesRequestHeaderDoc(
	requestPram *types.Request,
	input types.PostDocSDC,
) *apiOutputFormatter.PostDocSDC {
	responseJsonData := apiOutputFormatter.PostDocSDC{}
	responseBody := apiModuleRuntimesRequestsPostDoc.PostDocCreatesRequestHeaderDocs(
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
		controller.CustomLogger.Error("InstagramUserMediaRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostCreatesHeadersWithInstagramMediaController,
) InstagramUserMediaRequest(
	requestPram *types.Request,
	input types.InstagramUserMediaSDC,
) *apiModuleRuntimesResponsesInstagramUserMedia.InstagramUserMediaRes {
	responseJsonData := apiModuleRuntimesResponsesInstagramUserMedia.InstagramUserMediaRes{}
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
		controller.CustomLogger.Error("InstagramUserMediaRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *PostCreatesHeadersWithInstagramMediaController,
) request(
	input types.InstagramUserMediaSDC,
) {
	var instagramUserMediaResponse *apiModuleRuntimesResponsesInstagramUserMedia.InstagramUserMediaRes

	instagramUserMediaResponse = controller.InstagramUserMediaRequest(
		controller.UserInfo,
		input,
	)

	convertedToPostCreatesRequestHeadersWithInstagramMedia, error := formatter.ConvertToPostCreatesHeadersWithInstagramMedia(
		controller.UserInfo,
		instagramUserMediaResponse,
	)
	if error != nil {
		services.HandleError(
			&controller.Controller,
			error,
			nil,
		)
		controller.CustomLogger.Error("ConvertToPostCreatesHeadersWithInstagramMedia error")
	}

	postResponse := controller.PostCreatesRequestHeadersWithInstagramMedia(
		controller.UserInfo,
		*convertedToPostCreatesRequestHeadersWithInstagramMedia,
	)

	convertedToPostDocCreatesHeaderDoc, error := formatter.ConvertToPostDocCreatesHeaderDoc(
		controller.UserInfo,
		postResponse,
	)
	if error != nil {
		services.HandleError(
			&controller.Controller,
			error,
			nil,
		)
		controller.CustomLogger.Error("ConvertToPostDocCreatesHeaderDoc error")
	}

	postDocCreatesResponse := controller.PostDocCreatesRequestHeaderDoc(
		controller.UserInfo,
		*convertedToPostDocCreatesHeaderDoc,
	)

	controller.Data["json"] = map[string]interface{}{
		"postResponse":           postResponse,
		"postDocCreatesResponse": postDocCreatesResponse,
	}
	controller.ServeJSON()
}
