package controllersGoogleAccountAuthKey

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsGoogleAccountAccessToken "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/google-account-access-token"
	apiModuleRuntimesRequestsGoogleAccountAuthKey "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/google-account-auth-key"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/config"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"strconv"
)

type GoogleAccountAuthKeyController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *GoogleAccountAuthKeyController) Get() {
	host := controller.Ctx.Input.Host()
	port := strconv.Itoa(controller.Ctx.Input.Port())
	var protocol string

	if port == "443" {
		protocol = "https://"
	} else {
		protocol = "http://"
	}

	uri := controller.Ctx.Input.URI()

	absoluteUrl := protocol + host + ":" + port + uri

	controller.request(absoluteUrl)
}

func (
	controller *GoogleAccountAuthKeyController,
) GoogleAccountAuthRequest(
	requestPram *types.Request,
	absoluteUrl string,
) *apiOutputFormatter.GoogleAccountAuthKeySDC {
	var input types.GoogleAccountAuthKeySDC

	input = types.GoogleAccountAuthKeySDC{
		GoogleAccountAuthKey: types.GoogleAccountAuthKey{
			URL: absoluteUrl,
		},
	}

	responseJsonData := apiOutputFormatter.GoogleAccountAuthKeySDC{}
	responseBody := apiModuleRuntimesRequestsGoogleAccountAuthKey.GoogleAccountAuthKeyRequests(
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
		controller.CustomLogger.Error("GoogleAccountAuthRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *GoogleAccountAuthKeyController,
) GoogleAccountAccessTokenRequest(
	requestPram *types.Request,
	authKeyRequestResponse *apiOutputFormatter.GoogleAccountAuthKeySDC,
) *apiOutputFormatter.GoogleAccountAccessTokenSDC {
	var input types.GoogleAccountAccessTokenSDC

	for _, v := range *authKeyRequestResponse.Message.GoogleAccountAuthKey {
		input = types.GoogleAccountAccessTokenSDC{
			GoogleAccountAccessToken: types.GoogleAccountAccessToken{
				URL: v.URL,
			},
		}
	}

	responseJsonData := apiOutputFormatter.GoogleAccountAccessTokenSDC{}
	responseBody := apiModuleRuntimesRequestsGoogleAccountAccessToken.GoogleAccountAccessTokenRequests(
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
		controller.CustomLogger.Error("GoogleAccountAccessTokenRequest Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *GoogleAccountAuthKeyController,
) request(
	absoluteUrl string,
) {
	var authKeyRequestResponse *apiOutputFormatter.GoogleAccountAuthKeySDC

	authKeyRequestResponse = controller.GoogleAccountAuthRequest(
		controller.UserInfo,
		absoluteUrl,
	)

	var accessTokenRequestResponse *apiOutputFormatter.GoogleAccountAccessTokenSDC

	accessTokenRequestResponse = controller.GoogleAccountAccessTokenRequest(
		controller.UserInfo,
		authKeyRequestResponse,
	)

	conf := config.NewConf()

	if accessTokenRequestResponse.Message.GoogleAccountAccessToken != nil {
		if len(*accessTokenRequestResponse.Message.GoogleAccountAccessToken) > 0 {
			token := (*accessTokenRequestResponse.Message.GoogleAccountAccessToken)[0].AccessToken

			redirectTo := conf.GoogleAccountAuth.CreateGoogleAuthAfterRedirectURL(
				token,
			)

			controller.Ctx.Redirect(301, redirectTo)
		}
	}
}
