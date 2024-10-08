package apiModuleRuntimesRequestsParticipation

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type ParticipationReq struct {
	Header          Header          `json:"Participation"`
	APIType         string          `json:"api_type"`
	Accepter        []string        `json:"accepter"`
}

type Header struct {
	Participation				*int	`json:"Participation"`
	ParticipationDate			string	`json:"ParticipationDate"`
	ParticipationTime			string	`json:"ParticipationTime"`
	Participator				int		`json:"Participator"`
	ParticipationObjectType		string	`json:"ParticipationObjectType"`
	ParticipationObject			int		`json:"ParticipationObject"`
	Attendance					*int	`json:"Attendance"`
	Invitation					*int	`json:"Invitation"`
	CreationDate				string	`json:"CreationDate"`
	CreationTime				string	`json:"CreationTime"`
	IsCancelled					*bool	`json:"IsCancelled"`
}

func FuncParticipationCreatesRequestAll(
	requestPram *types.Request,
	input ParticipationReq,
) ParticipationReq {
	req := ParticipationReq{
		Header: input.Header,
		APIType: "creates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func ParticipationCreatesRequestAll(
	requestPram *types.Request,
	input types.ParticipationSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PARTICIPATION_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request ParticipationReq

	isCancelled := false

	request = FuncParticipationCreatesRequestAll(
		requestPram,
		ParticipationReq{
			Header: Header{
				Participation:              nil,
		        ParticipationDate:			formattedDate,
		        ParticipationTime:			formattedTime,
		        Participator:				input.Header.Participator,
		        ParticipationObjectType:	"EVENT",
		        ParticipationObject:		input.Header.ParticipationObject,
		        Attendance:					nil,
		        Invitation:					nil,
		        CreationDate:				formattedDate,
		        CreationTime:				formattedTime,
		        IsCancelled:				&isCancelled,
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

func FuncParticipationCancelsRequestAll(
	requestPram *types.Request,
	input ParticipationReq,
) ParticipationReq {
	req := ParticipationReq{
		Header: input.Header,
		APIType: "cancels",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func ParticipationCancelsRequestAll(
	requestPram *types.Request,
	input types.ParticipationSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PARTICIPATION_SRV"
	aPIType := "cancels"

//	currentDateTime := time.Now()
//	formattedDate := currentDateTime.Format("2006-01-02")
//	formattedTime := currentDateTime.Format("15:04:05")

	var request ParticipationReq

	isCancelled := true

	request = FuncParticipationCancelsRequestAll(
		requestPram,
		ParticipationReq{
			Header: Header{
				Participation:				input.Header.Participation,
		        IsCancelled:				&isCancelled,
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
