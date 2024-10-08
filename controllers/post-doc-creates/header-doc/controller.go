package controllersPostDocCreatesHeaderDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsPostDoc "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/post-doc"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/base64"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
	"io/ioutil"
)

type PostDocCreatesHeaderDocController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *PostDocCreatesHeaderDocController) Post() {
	//inputParameter := types.PostDocInputRead(&controller.Controller)
	docData, err := controller.getImagesAsBase64()
	if err != nil {
		services.HandleError(
			&controller.Controller,
			err,
			nil,
		)
		controller.CustomLogger.Error("PostDocCreatesHeaderDoc getImagesAsBase64 error")
	}

	requestBody := types.PostDocSDC{}
	input := controller.Ctx.Request.MultipartForm.Value["headerInfo"]

	if len(input) > 0 {
		jsonInput := input[0]
		err := json.Unmarshal([]byte(jsonInput), &requestBody)
		if err != nil {
			services.HandleError(
				&controller.Controller,
				err,
				nil,
			)
			controller.CustomLogger.Error("PostDocCreatesHeaderDoc Unmarshal error")
			return
		}
	}
	requestBody.DocData = *docData
	controller.request(requestBody)
}

func (
	controller *PostDocCreatesHeaderDocController,
) PostDocCreatesHeaderDocController(
	requestPram *types.Request,
	input types.PostDocSDC,
) *apiOutputFormatter.PostDocSDC {
	responseJsonData := apiOutputFormatter.PostDocSDC{}

	responseBody := apiModuleRuntimesRequestsPostDoc.PostDocCreatesRequestHeaderDoc(
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
		controller.CustomLogger.Error("PostDocCreatesHeaderDoc Unmarshal error")
	}

	return &responseJsonData
}

func (controller *PostDocCreatesHeaderDocController) getImagesAsBase64() (*string, error) {
	var base64Images []string

	images := controller.Ctx.Request.MultipartForm.File["images"]

	//for _, fileHeader := range files {
	//	file, err := fileHeader.Open()
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	fileBytes, err := ioutil.ReadAll(file)
	//	if err != nil {
	//		file.Close()
	//		return nil, err
	//	}
	//
	//	base64Encoded := base64.StdEncoding.EncodeToString(fileBytes)
	//	base64Images = append(base64Images, base64Encoded)
	//
	//	file.Close()
	//}

	if len(images) > 0 {
		fileHeader := images[0]
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		base64Encoded := base64.StdEncoding.EncodeToString(fileBytes)
		base64Images = append(base64Images, base64Encoded)
	} else {
		return nil, xerrors.New("No images found")
	}

	return &base64Images[0], nil
}

func (
	controller *PostDocCreatesHeaderDocController,
) request(
	input types.PostDocSDC,
) {
	var response *apiOutputFormatter.PostDocSDC

	response = controller.PostDocCreatesHeaderDocController(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
