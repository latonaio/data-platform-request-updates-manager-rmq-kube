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
	SenderObjectType                      string  `json:"SenderObjectType"`
	SenderObject                          int     `json:"SenderObject"`
	ReceiverObjectType                    string  `json:"ReceiverObjectType"`
	ReceiverObject                        int     `json:"ReceiverObject"`
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
	Invitation                            *int    `json:"Invitation"`
	ValidityStartDate                     string  `json:"ValidityStartDate"`
	ValidityEndDate                       string  `json:"ValidityEndDate"`
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

	switch inputSDC.Accepter[0] {
	case "Event":
		func() {
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
						PointTransaction:                      nil,
						PointTransactionType:                  "100",
						PointTransactionDate:                  formattedDate,
						PointTransactionTime:                  formattedTime,
						SenderObjectType:                      "SHOP",
						SenderObject:                          input.PointTransactionHeader.SenderObject,
						ReceiverObjectType:                    "BUSINESS_PARTNER",
						ReceiverObject:                        input.PointTransactionHeader.ReceiverObject,
						PointSymbol:                           "POYPO",
						PlusMinus:                             "+",
						PointTransactionAmount:                input.PointTransactionHeader.PointTransactionAmount,
						PointTransactionObjectType:            "EVENT",
						PointTransactionObject:                input.PointTransactionHeader.PointTransactionObject,
						SenderPointBalanceBeforeTransaction:   input.PointTransactionHeader.SenderPointBalanceBeforeTransaction,
						SenderPointBalanceAfterTransaction:    input.PointTransactionHeader.SenderPointBalanceAfterTransaction,
						ReceiverPointBalanceBeforeTransaction: input.PointTransactionHeader.ReceiverPointBalanceBeforeTransaction,
						ReceiverPointBalanceAfterTransaction:  input.PointTransactionHeader.ReceiverPointBalanceAfterTransaction,
						Attendance:                            nil,
						Participation:                         nil,
						Invitation:                            nil,
						ValidityStartDate:                     formattedDate,
						ValidityEndDate:                       "9999-12-31",
						CreationDate:                          formattedDate,
						CreationTime:                          formattedTime,
						IsCancelled:                           &isCancelled,
					},
				},
			)
		}()
	case "PointConsumption":
		input, err := formatter.ConvertToPointTransactionCreatesHeaderFromPointConsumption(inputSDC)
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
					PointTransaction:                      nil,
					PointTransactionType:                  "200",
					PointTransactionDate:                  formattedDate,
					PointTransactionTime:                  formattedTime,
					SenderObjectType:                      "BUSINESS_PARTNER",
					SenderObject:                          input.PointTransactionHeader.SenderObject,
					ReceiverObjectType:                    "SHOP",
					ReceiverObject:                        input.PointTransactionHeader.ReceiverObject,
					PointSymbol:                           "POYPO",
					PlusMinus:                             "-",
					PointTransactionAmount:                input.PointTransactionHeader.PointTransactionAmount,
					PointTransactionObjectType:            "SHOP",
					PointTransactionObject:                input.PointTransactionHeader.PointTransactionObject,
					SenderPointBalanceBeforeTransaction:   input.PointTransactionHeader.SenderPointBalanceBeforeTransaction,
					SenderPointBalanceAfterTransaction:    input.PointTransactionHeader.SenderPointBalanceAfterTransaction,
					ReceiverPointBalanceBeforeTransaction: input.PointTransactionHeader.ReceiverPointBalanceBeforeTransaction,
					ReceiverPointBalanceAfterTransaction:  input.PointTransactionHeader.ReceiverPointBalanceAfterTransaction,
					Attendance:                            nil,
					Participation:                         nil,
					Invitation:                            nil,
					ValidityStartDate:                     formattedDate,
					ValidityEndDate:                       "9999-12-31",
					CreationDate:                          formattedDate,
					CreationTime:                          formattedTime,
					IsCancelled:                           &isCancelled,
				},
			},
		)
	default:
		input := inputSDC
		request = FuncPointTransactionCreatesRequestHeader(
			requestPram,
			PointTransactionReq{
				Header: Header{
					PointTransaction:                      nil,
					PointTransactionType:                  input.Header.PointTransactionType,
					PointTransactionDate:                  formattedDate,
					PointTransactionTime:                  formattedTime,
					SenderObjectType:                      input.Header.SenderObjectType,
					SenderObject:                          input.Header.SenderObject,
					ReceiverObjectType:                    input.Header.ReceiverObjectType,
					ReceiverObject:                        input.Header.ReceiverObject,
					PointSymbol:                           "POYPO",
					PlusMinus:                             input.Header.PlusMinus,
					PointTransactionAmount:                input.Header.PointTransactionAmount,
					PointTransactionObjectType:            input.Header.PointTransactionObjectType,
					PointTransactionObject:                input.Header.PointTransactionObject,
					SenderPointBalanceBeforeTransaction:   input.Header.SenderPointBalanceBeforeTransaction,
					SenderPointBalanceAfterTransaction:    input.Header.SenderPointBalanceAfterTransaction,
					ReceiverPointBalanceBeforeTransaction: input.Header.ReceiverPointBalanceBeforeTransaction,
					ReceiverPointBalanceAfterTransaction:  input.Header.ReceiverPointBalanceAfterTransaction,
					Attendance:                            nil,
					Participation:                         nil,
					Invitation:                            nil,
					ValidityStartDate:                     formattedDate,
					ValidityEndDate:                       "9999-12-31",
					CreationDate:                          formattedDate,
					CreationTime:                          formattedTime,
					IsCancelled:                           &isCancelled,
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

func FuncPointTransactionCancelsRequestHeader(
	requestPram *types.Request,
	input PointTransactionReq,
) PointTransactionReq {
	req := PointTransactionReq{
		Header: input.Header,
		APIType: "cancels",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func PointTransactionCancelsRequestHeader(
	requestPram *types.Request,
	input types.PointTransactionSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_TRANSACTION_SRV"
	aPIType := "cancels"

//	currentDateTime := time.Now()
//	formattedDate := currentDateTime.Format("2006-01-02")
//	formattedTime := currentDateTime.Format("15:04:05")

	var request PointTransactionReq

	isCancelled := true

	request = FuncPointTransactionCancelsRequestHeader(
		requestPram,
		PointTransactionReq{
			Header: Header{
				PointTransaction:			input.Header.PointTransaction,
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
