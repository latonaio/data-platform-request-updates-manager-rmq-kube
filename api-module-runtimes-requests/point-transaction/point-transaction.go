package apiModuleRuntimesRequestsPointTransaction

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/formatter"
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type PointTransactionReq struct {
	Header   Header   `json:"PointTransaction"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	PointTransaction                      *int    `json:"PointTransaction"`
	PointTransactionType                  string  `json:"PointTransactionType"`
	PointTransactionDate                  string  `json:"PointTransactionDate"`
	PointTransactionTime                  string  `json:"PointTransactionTime"`
	Sender                                int     `json:"Sender"`
	Receiver                              int     `json:"Receiver"`
	PointSymbol                           string  `json:"PointSymbol"`
	PlusMinus                             string  `json:"PlusMinus"`
	PointTransactionAmount                float32 `json:"PointTransactionAmount"`
	PointTransactionObjectType            string  `json:"PointTransactionObjectType"`
	PointTransactionObject                int     `json:"PointTransactionObject"`
	SenderPointBalanceBeforeTransaction   float32 `json:"SenderPointBalanceBeforeTransaction"`
	SenderPointBalanceAfterTransaction    float32 `json:"SenderPointBalanceAfterTransaction"`
	ReceiverPointBalanceBeforeTransaction float32 `json:"ReceiverPointBalanceBeforeTransaction"`
	ReceiverPointBalanceAfterTransaction  float32 `json:"ReceiverPointBalanceAfterTransaction"`
	Attendance                            *int    `json:"Attendance"`
	Participation                         *int    `json:"Participation"`
	CreationDate                          string  `json:"CreationDate"`
	CreationTime                          string  `json:"CreationTime"`
	IsCancelled                           *bool   `json:"IsCancelled"`
}

func FuncPointTransactionCreatesRequestHeader(
	requestPram *types.Request,
	input PointTransactionReq,
) PointTransactionReq {
	req := PointTransactionReq{
		Header:  input.Header,
		APIType: "creates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func PointTransactionCreatesRequestHeader(
	requestPram *types.Request,
	inputSDC types.PointTransactionSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_TRANSACTION_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request PointTransactionReq

	isCancelled := false

	if inputSDC.Accepter[0] == "Event" {
		input, err := formatter.ConvertToPointTransactionCreatesHeaderFromEvent(inputSDC)
		if err != nil {
			services.HandleError(
				controller,
				err,
				nil,
			)
		}
		request = FuncPointTransactionCreatesRequestHeader(
			requestPram,
			PointTransactionReq{
				Header: Header{
					PointTransaction:       nil,
					PointTransactionType:   input.Header.PointTransactionType,
					PointTransactionDate:   formattedDate,
					PointTransactionTime:   formattedTime,
					Sender:                 input.Header.Sender,
					Receiver:               input.Header.Receiver,
					PointSymbol:            "POYPO",
					PlusMinus:              input.Header.PlusMinus,
					PointTransactionAmount: input.Header.PointTransactionAmount,
					//PointTransactionAmount:     				3000,
					PointTransactionObjectType: input.Header.PointTransactionObjectType,
					PointTransactionObject:     input.Header.PointTransactionObject,
					//SenderPointBalanceBeforeTransaction:  	input.Header.SenderPointBalanceBeforeTransaction,
					SenderPointBalanceBeforeTransaction: 100000,
					//SenderPointBalanceAfterTransaction:		input.Header.SenderPointBalanceAfterTransaction,
					SenderPointBalanceAfterTransaction: 97000,
					//ReceiverPointBalanceBeforeTransaction: 	input.Header.ReceiverPointBalanceBeforeTransaction,
					ReceiverPointBalanceBeforeTransaction: 2000,
					//ReceiverPointBalanceAfterTransaction:		input.Header.ReceiverPointBalanceAfterTransaction,
					ReceiverPointBalanceAfterTransaction: 5000,
					Attendance:                           input.Header.Attendance,
					Participation:                        input.Header.Participation,
					CreationDate:                         formattedDate,
					CreationTime:                         formattedTime,
					IsCancelled:                          &isCancelled,
				},
			},
		)
	} else {
		input := inputSDC
		request = FuncPointTransactionCreatesRequestHeader(
			requestPram,
			PointTransactionReq{
				Header: Header{
					PointTransaction:       nil,
					PointTransactionType:   input.HeaderFromEvent.PointTransactionType,
					PointTransactionDate:   formattedDate,
					PointTransactionTime:   formattedTime,
					Sender:                 input.HeaderFromEvent.Sender,
					Receiver:               input.HeaderFromEvent.Receiver,
					PointSymbol:            "POYPO",
					PlusMinus:              input.HeaderFromEvent.PlusMinus,
					PointTransactionAmount: input.HeaderFromEvent.PointTransactionAmount,
					//PointTransactionAmount:     				3000,
					PointTransactionObjectType: input.HeaderFromEvent.PointTransactionObjectType,
					PointTransactionObject:     input.HeaderFromEvent.PointTransactionObject,
					//SenderPointBalanceBeforeTransaction:		input.Header.SenderPointBalanceBeforeTransaction,
					SenderPointBalanceBeforeTransaction: 100000,
					//SenderPointBalanceAfterTransaction:		input.Header.SenderPointBalanceAfterTransaction,
					SenderPointBalanceAfterTransaction: 97000,
					//ReceiverPointBalanceBeforeTransaction:	input.Header.ReceiverPointBalanceBeforeTransaction,
					ReceiverPointBalanceBeforeTransaction: 2000,
					//ReceiverPointBalanceAfterTransaction:		input.Header.ReceiverPointBalanceAfterTransaction,
					ReceiverPointBalanceAfterTransaction: 5000,
					Attendance:                           input.HeaderFromEvent.Attendance,
					Participation:                        input.HeaderFromEvent.Participation,
					CreationDate:                         formattedDate,
					CreationTime:                         formattedTime,
					IsCancelled:                          &isCancelled,
				},
			},
		)
	}

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
