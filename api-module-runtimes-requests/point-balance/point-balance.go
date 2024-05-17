package apiModuleRuntimesRequestsPointBalance

import (
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
	input types.PointBalanceSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POINT_BALANCE_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request PointBalanceReq

	request = FuncPointBalanceUpdatesRequestPointBalance(
		requestPram,
		PointBalanceReq{
			PointBalance: PointBalance{
				BusinessPartner: input.PointBalance.BusinessPartner,
				PointSymbol:     input.PointBalance.PointSymbol,
				//CurrentBalance:  currentBalance + input.PointBalance.PointTransactionAmount, // 要確認
				LastChangeDate: formattedDate,
				LastChangeTime: formattedTime,
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
