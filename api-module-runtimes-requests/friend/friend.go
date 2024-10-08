package apiModuleRuntimesRequestsFriend

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type FriendReq struct {
	General  General  `json:"Friend"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type General struct {
	BusinessPartner           int    `json:"BusinessPartner"`
	Friend                    int    `json:"Friend"`
	BPBusinessPartnerType     string `json:"BPBusinessPartnerType"`
	FriendBusinessPartnerType string `json:"FriendBusinessPartnerType"`
	RankType                  string `json:"RankType"`
	Rank                      int    `json:"Rank"`
	FriendIsBlocked           bool   `json:"FriendIsBlocked"`
	CreationDate              string `json:"CreationDate"`
	CreationTime              string `json:"CreationTime"`
	LastChangeDate            string `json:"LastChangeDate"`
	LastChangeTime            string `json:"LastChangeTime"`
	IsMarkedForDeletion       *bool  `json:"IsMarkedForDeletion"`
}

func FuncFriendCreatesRequestGeneralMe(
	requestPram *types.Request,
	input FriendReq,
) FriendReq {
	req := FriendReq{
		General: input.General,
		APIType: "creates",
		Accepter: []string{
			"GeneralMe",
		},
	}
	return req
}

func FriendCreatesRequestGeneralMe(
	requestPram *types.Request,
	input types.FriendSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_FRIEND_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request FriendReq

	friendIsBlocked := false
	isMarkedForDeletion := false

	request = FuncFriendCreatesRequestGeneralMe(
		requestPram,
		FriendReq{
			General: General{
				BusinessPartner:           input.General.BusinessPartner,
				Friend:                    input.General.Friend,
				BPBusinessPartnerType:     "02",
				FriendBusinessPartnerType: "02",
				RankType:                  "COMM",
				Rank:                      1,
				FriendIsBlocked:           friendIsBlocked,
				CreationDate:              formattedDate,
				CreationTime:              formattedTime,
				LastChangeDate:            formattedDate,
				LastChangeTime:            formattedTime,
				IsMarkedForDeletion:       &isMarkedForDeletion,
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

func FuncFriendCreatesRequestGeneralFriend(
	requestPram *types.Request,
	input FriendReq,
) FriendReq {
	req := FriendReq{
		General: input.General,
		APIType: "creates",
		Accepter: []string{
			"GeneralFriend",
		},
	}
	return req
}

func FriendCreatesRequestGeneralFriend(
	requestPram *types.Request,
	input types.FriendSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_FRIEND_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request FriendReq

	friendIsBlocked := false
	isMarkedForDeletion := false

	request = FuncFriendCreatesRequestGeneralFriend(
		requestPram,
		FriendReq{
			General: General{
				BusinessPartner:           input.General.Friend,
				Friend:                    input.General.BusinessPartner,
				BPBusinessPartnerType:     "02",
				FriendBusinessPartnerType: "02",
				RankType:                  "COMM",
				Rank:                      1,
				FriendIsBlocked:           friendIsBlocked,
				CreationDate:              formattedDate,
				CreationTime:              formattedTime,
				LastChangeDate:            formattedDate,
				LastChangeTime:            formattedTime,
				IsMarkedForDeletion:       &isMarkedForDeletion,
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

func FuncFriendUpdatesRequestGeneral(
	requestPram *types.Request,
	input FriendReq,
) FriendReq {
	req := FriendReq{
		General: input.General,
		APIType: "updates",
		Accepter: []string{
			"General",
		},
	}
	return req
}

func FriendUpdatesRequestGeneral(
	requestPram *types.Request,
	input types.FriendSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_FRIEND_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request FriendReq

	request = FuncFriendUpdatesRequestGeneral(
		requestPram,
		FriendReq{
			General: General{
				BusinessPartner: input.General.BusinessPartner,
				Friend:          input.General.Friend,
				Rank:            input.General.Rank,
				FriendIsBlocked: input.General.FriendIsBlocked,
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
