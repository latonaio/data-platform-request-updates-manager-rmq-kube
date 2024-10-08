package apiModuleRuntimesRequestsMessage

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type MessageReq struct {
	Header   Header   `json:"Message"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Message             *int   `json:"Message"`
	MessageType         string `json:"MessageType"`
	Sender              int    `json:"Sender"`
	Receiver            int    `json:"Receiver"`
	Language            string `json:"Language"`
	Title               string `json:"Title"`
	LongText            string `json:"LongText"`
	MessageIsSent       bool   `json:"MessageIsSent"`
	MessageIsRead       bool   `json:"MessageIsRead"`
	CreationDate        string `json:"CreationDate"`
	CreationTime        string `json:"CreationTime"`
	LastChangeDate      string `json:"LastChangeDate"`
	LastChangeTime      string `json:"LastChangeTime"`
	IsCancelled         *bool  `json:"IsCancelled"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

func FuncMessageCreatesRequestAll(
	requestPram *types.Request,
	input MessageReq,
) MessageReq {
	req := MessageReq{
		Header:  input.Header,
		APIType: "creates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func MessageCreatesRequestAll(
	requestPram *types.Request,
	input types.MessageSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_MESSAGE_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request MessageReq

	isCancelled := false
	isMarkedForDeletion := false

	request = FuncMessageCreatesRequestAll(
		requestPram,
		MessageReq{
			Header: Header{
				Message:             nil,
				MessageType:         input.Header.MessageType,
				Sender:              input.Header.Sender,
				Receiver:            input.Header.Receiver,
				Language:            input.Header.Language,
				Title:               input.Header.Title,
				LongText:            input.Header.LongText,
				MessageIsSent:       true,
				MessageIsRead:       false,
				CreationDate:        formattedDate,
				CreationTime:        formattedTime,
				LastChangeDate:      formattedDate,
				LastChangeTime:      formattedTime,
				IsCancelled:         &isCancelled,
				IsMarkedForDeletion: &isMarkedForDeletion,
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
