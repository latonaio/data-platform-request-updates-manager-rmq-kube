package apiModuleRuntimesRequestsAttendance

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type AttendanceReq struct {
	Header   Header   `json:"Attendance"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Attendance           *int   `json:"Attendance"`
	AttendanceDate       string `json:"AttendanceDate"`
	AttendanceTime       string `json:"AttendanceTime"`
	Attender             int    `json:"Attender"`
	AttendanceObjectType string `json:"AttendanceObjectType"`
	AttendanceObject     int    `json:"AttendanceObject"`
	Participation        *int   `json:"Participation"`
	CreationDate         string `json:"CreationDate"`
	CreationTime         string `json:"CreationTime"`
	IsCancelled          *bool  `json:"IsCancelled"`
}

func FuncAttendanceCreatesRequestAll(
	requestPram *types.Request,
	input AttendanceReq,
) AttendanceReq {
	req := AttendanceReq{
		Header:  input.Header,
		APIType: "creates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func AttendanceCreatesRequestAll(
	requestPram *types.Request,
	input types.AttendanceSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ATTENDANCE_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request AttendanceReq

	isCancelled := false

	request = FuncAttendanceCreatesRequestAll(
		requestPram,
		AttendanceReq{
			Header: Header{
				Attendance:           nil,
				AttendanceDate:       formattedDate,
				AttendanceTime:       formattedTime,
				Attender:             input.Header.Attender,
				AttendanceObjectType: "EVENT",
				AttendanceObject:     input.Header.AttendanceObject,
				Participation:        nil,
				CreationDate:         formattedDate,
				CreationTime:         formattedTime,
				IsCancelled:          &isCancelled,
			},
		},
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}

func FuncAttendanceCancelsRequestAll(
	requestPram *types.Request,
	input AttendanceReq,
) AttendanceReq {
	req := AttendanceReq{
		Header:  input.Header,
		APIType: "cancels",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func AttendanceCancelsRequestAll(
	requestPram *types.Request,
	input types.AttendanceSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ATTENDANCE_SRV"
	aPIType := "cancels"

	//	currentDateTime := time.Now()
	//	formattedDate := currentDateTime.Format("2006-01-02")
	//	formattedTime := currentDateTime.Format("15:04:05")

	var request AttendanceReq

	isCancelled := true

	request = FuncAttendanceCancelsRequestAll(
		requestPram,
		AttendanceReq{
			Header: Header{
				Attendance:  input.Header.Attendance,
				IsCancelled: &isCancelled,
			},
		},
	)

	marshaledRequest, err := json.Marshal(request)
	if err != nil {
		services.HandleError(
			controller,
			err,
			nil,
		)
	}

	responseBody := services.Request(
		aPIServiceName,
		aPIType,
		ioutil.NopCloser(strings.NewReader(string(marshaledRequest))),
		controller,
	)

	return responseBody
}
