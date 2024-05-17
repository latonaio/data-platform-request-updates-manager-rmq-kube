package apiModuleRuntimesRequestsPost

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type PostReq struct {
	Header          Header          `json:"Post"`
	APIType         string          `json:"api_type"`
	Accepter        []string        `json:"accepter"`
}

type Header struct {
	Post							*int	`json:"Post"`
	PostType						string	`json:"PostType"`
	PostOwner						int		`json:"PostOwner"`
	PostOwnerBusinessPartnerRole	string	`json:"PostOwnerBusinessPartnerRole"`
	Description						string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
}

func FuncPostCreatesRequestAll(
	requestPram *types.Request,
	input PostReq,
) PostReq {
	req := PostReq{
		Header: input.Header,
		APIType: "creates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func PostCreatesRequestAll(
	requestPram *types.Request,
	input types.PostSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POST_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request PostReq

	isMarkedForDeletion := false

	request = FuncPostCreatesRequestAll(
		requestPram,
		PostReq{
			Header: Header{
				Post:				    		nil,
		        PostType:						input.Header.PostType,
		        PostOwner:						input.Header.PostDate,
		        PostOwnerBusinessPartnerRole:	input.Header.PostTime,
		        Description:					input.Header.Description,
		        LongText						input.Header.LongText,
		        Tag1:                           input.Header.Tag1,
		        Tag2:                           input.Header.Tag2,
		        Tag3:                           input.Header.Tag3,
		        Tag4:                           input.Header.Tag4,
		        CreationDate:					formattedDate,
		        CreationTime:					formattedTime,
		        LastChangeDate:					formattedDate,
		        LastChangeTime:					formattedTime,
		        IsMarkedForDeletion:			&isMarkedForDeletion,
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

func FuncPostUpdatesRequestHeader(
	requestPram *types.Request,
	input PostReq,
) PostReq {
	req := PostReq{
		Header: input.Header,
		APIType: "updates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func PostUpdatesRequestHeader(
	requestPram *types.Request,
	input types.PostSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POST_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request PostReq

	request = FuncPostUpdatesRequestHeader(
		requestPram,
		PostReq{
			Header: Header{
				Post:				    		input.Header.Post,
		        Description:					input.Header.Description,
		        LongText						input.Header.LongText,
		        Tag1:                           input.Header.Tag1,
		        Tag2:                           input.Header.Tag2,
		        Tag3:                           input.Header.Tag3,
		        Tag4:                           input.Header.Tag4,
		        LastChangeDate:					formattedDate,
		        LastChangeTime:					formattedTime,
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
