package apiModuleRuntimesRequestsPointBalance

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

type PointBalanceReq struct {
	PointBalance PointBalance `json:"PointBalance"`
	APIType      string       `json:"api_type"`
	Accepter     []string     `json:"accepter"`
}

type PointBalance struct {
	BusinessPartner int      `json:"BusinessPartner"`
	PointSymbol     string   `json:"PointSymbol"`
	CurrentBalance  float32  `json:"CurrentBalance"`
	LimitBalance    *float32 `json:"LimitBalance"`
	CreationDate    string   `json:"CreationDate"`
	CreationTime    string   `json:"CreationTime"`
	LastChangeDate  string   `json:"LastChangeDate"`
	LastChangeTime  string   `json:"LastChangeTime"`
	ByShop          []ByShop `json:"ByShop"`
}

type ByShop struct {
	BusinessPartner int      `json:"BusinessPartner"`
	PointSymbol     string   `json:"PointSymbol"`
	Shop            int      `json:"Shop"`
	CurrentBalance  float32  `json:"CurrentBalance"`
	LimitBalance    *float32 `json:"LimitBalance"`
	CreationDate    string   `json:"CreationDate"`
	CreationTime    string   `json:"CreationTime"`
	LastChangeDate  string   `json:"LastChangeDate"`
	LastChangeTime  string   `json:"LastChangeTime"`
}

func FuncPointBalanceCreatesRequestPointBalance(
	requestPram *types.Request,
	input PointBalanceReq,
) PointBalanceReq {
	req := PointBalanceReq{
		PointBalance: input.PointBalance,
		APIType:      "creates",
		Accepter: []string{
			"PointBalance",
		},
	}
	return req
}

func PointBalanceCreatesRequestPointBalance(
	requestPram *types.Request,
	input types.PointBalanceSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_BALANCE_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request PointBalanceReq

	request = FuncPointBalanceCreatesRequestPointBalance(
		requestPram,
		PointBalanceReq{
			PointBalance: PointBalance{
				BusinessPartner: input.PointBalance.BusinessPartner,
				PointSymbol:     "POYPO",
				CurrentBalance:  0,
				CreationDate:    formattedDate,
				CreationTime:    formattedTime,
				LastChangeDate:  formattedDate,
				LastChangeTime:  formattedTime,
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

func FuncPointBalanceUpdatesRequestPointBalance(
	requestPram *types.Request,
	input PointBalanceReq,
) PointBalanceReq {
	req := PointBalanceReq{
		PointBalance: input.PointBalance,
		APIType:      "updates",
		Accepter: []string{
			"PointBalance",
		},
	}
	return req
}

func PointBalanceUpdatesRequestPointBalance(
	requestPram *types.Request,
	inputSDC types.PointBalanceSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_BALANCE_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request PointBalanceReq

	switch inputSDC.Accepter[0] {
	case "PointConsumptionSender":
		func() {
			input, err := formatter.ConvertToPointBalanceUpdatesPointConsumptionSender(inputSDC)
			if err != nil {
				services.HandleError(
					controller,
					err,
					nil,
				)
			}

			request = FuncPointBalanceUpdatesRequestPointBalance(
				requestPram,
				PointBalanceReq{
					PointBalance: PointBalance{
						BusinessPartner: input.PointBalancePointBalance.BusinessPartner,
						PointSymbol:     "POYPO",
						CurrentBalance:  *input.PointBalancePointBalance.CurrentBalance,
						LastChangeDate:  formattedDate,
						LastChangeTime:  formattedTime,
					},
				},
			)
		}()
	case "PointAcquisitionReceiver":
		func() {
			input, err := formatter.ConvertToPointBalanceUpdatesPointAcquisitionReceiver(inputSDC)
			if err != nil {
				services.HandleError(
					controller,
					err,
					nil,
				)
			}

			request = FuncPointBalanceUpdatesRequestPointBalance(
				requestPram,
				PointBalanceReq{
					PointBalance: PointBalance{
						BusinessPartner: input.PointBalancePointBalance.BusinessPartner,
						PointSymbol:     "POYPO",
						CurrentBalance:  *input.PointBalancePointBalance.CurrentBalance,
						LastChangeDate:  formattedDate,
						LastChangeTime:  formattedTime,
					},
				},
			)
		}()
	default:
		//input := inputSDC
		//request = FuncPointBalanceUpdatesRequestPointBalance(
		//	requestPram,
		//	PointBalanceReq{
		//		PointBalance: PointBalance{
		//			BusinessPartner: input.PointBalance.BusinessPartner,
		//			PointSymbol:     "POYPO",
		//			CurrentBalance:  *input.PointBalance.CurrentBalance,
		//			LastChangeDate:  formattedDate,
		//			LastChangeTime:  formattedTime,
		//		},
		//	},
		//)
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

func FuncPointBalanceUpdatesRequestByShop(
	requestPram *types.Request,
	input PointBalanceReq,
) PointBalanceReq {
	req := PointBalanceReq{
		PointBalance: input.PointBalance,
		APIType:      "updates",
		Accepter: []string{
			"ByShop",
		},
	}
	return req
}

func PointBalanceUpdatesRequestByShop(
	requestPram *types.Request,
	inputSDC types.PointBalanceSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_BALANCE_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request PointBalanceReq

	switch inputSDC.Accepter[0] {
	case "PointConsumptionReceiver":
		func() {
			input, err := formatter.ConvertToPointBalanceUpdatesPointConsumptionReceiver(inputSDC)
			if err != nil {
				services.HandleError(
					controller,
					err,
					nil,
				)
			}

			request = FuncPointBalanceUpdatesRequestByShop(
				requestPram,
				PointBalanceReq{
					PointBalance: PointBalance{
						BusinessPartner: input.PointBalancePointBalance.BusinessPartner,
						PointSymbol:     "POYPO",
						ByShop: []ByShop{
							{
								Shop:           input.PointBalancePointBalance.PointBalanceByShop[0].Shop,
								CurrentBalance: input.PointBalancePointBalance.PointBalanceByShop[0].CurrentBalance,
								LastChangeDate: formattedDate,
								LastChangeTime: formattedTime,
							},
						},
					},
				},
			)
		}()
	case "PointAcquisitionSender":
		func() {
			input, err := formatter.ConvertToPointBalanceUpdatesPointAcquisitionSender(inputSDC)
			if err != nil {
				services.HandleError(
					controller,
					err,
					nil,
				)
			}

			request = FuncPointBalanceUpdatesRequestByShop(
				requestPram,
				PointBalanceReq{
					PointBalance: PointBalance{
						BusinessPartner: input.PointBalancePointBalance.BusinessPartner,
						PointSymbol:     "POYPO",
						ByShop: []ByShop{
							{
								Shop:           input.PointBalancePointBalance.PointBalanceByShop[0].Shop,
								CurrentBalance: input.PointBalancePointBalance.PointBalanceByShop[0].CurrentBalance,
								LastChangeDate: formattedDate,
								LastChangeTime: formattedTime,
							},
						},
					},
				},
			)
		}()
	default:
		input := inputSDC
		request = FuncPointBalanceUpdatesRequestByShop(
			requestPram,
			PointBalanceReq{
				PointBalance: PointBalance{
					BusinessPartner: input.PointBalance.BusinessPartner,
					PointSymbol:     "POYPO",
					ByShop: []ByShop{
						{
							Shop:           input.PointBalance.PointBalanceByShop[0].Shop,
							CurrentBalance: input.PointBalance.PointBalanceByShop[0].CurrentBalance,
							LastChangeDate: formattedDate,
							LastChangeTime: formattedTime,
						},
					},
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
