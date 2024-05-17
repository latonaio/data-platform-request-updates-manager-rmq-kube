package apiModuleRuntimesRequestsAuthenticator

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type AuthenticatorReq struct {
	User     User     `json:"Authenticator"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type User struct {
	UserID              string  `json:"UserID"`
	BusinessPartner     int     `json:"BusinessPartner"`
	Password            string  `json:"Password"`
	Qos                 string  `json:"Qos"`
	IsEncrypt           bool    `json:"IsEncrypt"`
	Language            string  `json:"Language"`
	LastLoginDate       *string `json:"LastLoginDate"`
	LastLoginTime       *string `json:"LastLoginTime"`
	CreationDate        string  `json:"CreationDate"`
	CreationTime        string  `json:"CreationTime"`
	LastChangeDate      string  `json:"LastChangeDate"`
	LastChangeTime      string  `json:"LastChangeTime"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func FuncAuthenticatorCreatesRequestUser(
	requestPram *types.Request,
	input AuthenticatorReq,
) AuthenticatorReq {
	req := AuthenticatorReq{
		User:    input.User,
		APIType: "creates",
		Accepter: []string{
			"User",
		},
	}
	return req
}

func AuthenticatorCreatesRequestUser(
	requestPram *types.Request,
	input types.AuthenticatorSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_AUTHENTICATOR_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request AuthenticatorReq

	isEncrypt := true
	isMarkedForDeletion := false
	businessPartner := 401

	request = FuncAuthenticatorCreatesRequestUser(
		requestPram,
		AuthenticatorReq{
			User: User{
				UserID:              input.User.UserID,
				BusinessPartner:     businessPartner,
				Password:            input.User.Password,
				Qos:                 input.User.Qos,
				IsEncrypt:           isEncrypt,
				Language:            input.User.Language,
				LastLoginDate:       input.User.LastLoginDate,
				LastLoginTime:       input.User.LastLoginTime,
				CreationDate:        formattedDate,
				CreationTime:        formattedTime,
				LastChangeDate:      formattedDate,
				LastChangeTime:      formattedTime,
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

func FuncAuthenticatorUpdatesRequestUser(
	requestPram *types.Request,
	input AuthenticatorReq,
) AuthenticatorReq {
	req := AuthenticatorReq{
		User:    input.User,
		APIType: "updates",
		Accepter: []string{
			"User",
		},
	}
	return req
}

func AuthenticatorUpdatesRequestUser(
	requestPram *types.Request,
	input types.AuthenticatorSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_AUTHENTICATOR_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request AuthenticatorReq

	businessPartner := 401
	isEncrypt := true

	request = FuncAuthenticatorUpdatesRequestUser(
		requestPram,
		AuthenticatorReq{
			User: User{
				UserID:          input.User.UserID,
				BusinessPartner: businessPartner,
				Password:        input.User.Password,
				Qos:             input.User.Qos,
				IsEncrypt:       isEncrypt,
				LastLoginDate:   input.User.LastLoginDate,
				LastLoginTime:   input.User.LastLoginTime,
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
