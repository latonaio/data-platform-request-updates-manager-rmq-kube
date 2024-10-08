package controllersAttendanceCancelsAll

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesRequestsAttendance "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-requests/attendance"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type AttendanceCancelsAllController struct {
	beego.Controller
	RedisKey     string
	UserInfo     *types.Request
	CustomLogger *logger.Logger
}

func (controller *AttendanceCancelsAllController) Post() {
	inputParameter := types.AttendanceInputRead(&controller.Controller)
	controller.request(inputParameter)
}

func (
	controller *AttendanceCancelsAllController,
) AttendanceCancelsRequestAll(
	requestPram *types.Request,
	input types.AttendanceSDC,
) *apiOutputFormatter.AttendanceSDC {
	responseJsonData := apiOutputFormatter.AttendanceSDC{}
	responseBody := apiModuleRuntimesRequestsAttendance.AttendanceCancelsRequestAll(
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
		controller.CustomLogger.Error("AttendanceCancelsRequestAll Unmarshal error")
	}

	return &responseJsonData
}

func (
	controller *AttendanceCancelsAllController,
) request(
	input types.AttendanceSDC,
) {
	var response *apiOutputFormatter.AttendanceSDC

	response = controller.AttendanceCancelsRequestAll(
		controller.UserInfo,
		input,
	)

	controller.Data["json"] = &response
	controller.ServeJSON()
}
