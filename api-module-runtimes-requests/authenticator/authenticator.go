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
	UserID                   string                   `json:"UserID"`
	BusinessPartner          int                      `json:"BusinessPartner"`
	Password                 string                   `json:"Password"`
	Qos                      string                   `json:"Qos"`
	IsEncrypt                bool                     `json:"IsEncrypt"`
	Language                 string                   `json:"Language"`
	LastLoginDate            *string                  `json:"LastLoginDate"`
	LastLoginTime            *string                  `json:"LastLoginTime"`
	CreationDate             string                   `json:"CreationDate"`
	CreationTime             string                   `json:"CreationTime"`
	LastChangeDate           string                   `json:"LastChangeDate"`
	LastChangeTime           string                   `json:"LastChangeTime"`
	IsMarkedForDeletion      *bool                    `json:"IsMarkedForDeletion"`
	InitialEmailAuth         InitialEmailAuth         `json:"InitialEmailAuth"`
	InitialSMSAuth           InitialSMSAuth           `json:"InitialSMSAuth"`
	SMSAuth                  SMSAuth                  `json:"SMSAuth"`
	InitialGoogleAccountAuth InitialGoogleAccountAuth `json:"InitialGoogleAccountAuth"`
	GoogleAccountAuth        GoogleAccountAuth        `json:"GoogleAccountAuth"`
	InstagramAuth            InstagramAuth            `json:"InstagramAuth"`
}

type InitialEmailAuth struct {
	EmailAddress        string `json:"EmailAddress"`
	CreationDate        string `json:"CreationDate"`
	CreationTime        string `json:"CreationTime"`
	LastChangeDate      string `json:"LastChangeDate"`
	LastChangeTime      string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type InitialSMSAuth struct {
	MobilePhoneNumber   string `json:"MobilePhoneNumber"`
	AuthenticationCode  int    `json:"AuthenticationCode"`
	CreationDate        string `json:"CreationDate"`
	CreationTime        string `json:"CreationTime"`
	LastChangeDate      string `json:"LastChangeDate"`
	LastChangeTime      string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type SMSAuth struct {
	UserID              string `json:"UserID"`
	MobilePhoneNumber   string `json:"MobilePhoneNumber"`
	AuthenticationCode  int    `json:"AuthenticationCode"`
	CreationDate        string `json:"CreationDate"`
	CreationTime        string `json:"CreationTime"`
	LastChangeDate      string `json:"LastChangeDate"`
	LastChangeTime      string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type InitialGoogleAccountAuth struct {
	EmailAddress        string `json:"EmailAddress"`
	GoogleID            string `json:"GoogleID"`
	AccessToken         string `json:"AccessToken"`
	CreationDate        string `json:"CreationDate"`
	CreationTime        string `json:"CreationTime"`
	LastChangeDate      string `json:"LastChangeDate"`
	LastChangeTime      string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type GoogleAccountAuth struct {
	UserID              string `json:"UserID"`
	EmailAddress        string `json:"EmailAddress"`
	GoogleID            string `json:"GoogleID"`
	AccessToken         string `json:"AccessToken"`
	CreationDate        string `json:"CreationDate"`
	CreationTime        string `json:"CreationTime"`
	LastChangeDate      string `json:"LastChangeDate"`
	LastChangeTime      string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type InstagramAuth struct {
	UserID              string `json:"UserID"`
	InstagramID         string `json:"InstagramID"`
	AccessToken         string `json:"AccessToken"`
	CreationDate        string `json:"CreationDate"`
	CreationTime        string `json:"CreationTime"`
	LastChangeDate      string `json:"LastChangeDate"`
	LastChangeTime      string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
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

	request = FuncAuthenticatorCreatesRequestUser(
		requestPram,
		AuthenticatorReq{
			User: User{
				UserID:              input.User.UserID,
				BusinessPartner:     input.User.BusinessPartner,
				Password:            input.User.Password,
				Qos:                 input.User.Qos,
				IsEncrypt:           isEncrypt,
				Language:            "JA",
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

func FuncAuthenticatorCreatesRequestInitialEmailAuth(
	requestPram *types.Request,
	input AuthenticatorReq,
) AuthenticatorReq {
	req := AuthenticatorReq{
		User:    input.User,
		APIType: "creates",
		Accepter: []string{
			"InitialEmailAuth",
		},
	}
	return req
}

func AuthenticatorCreatesRequestInitialEmailAuth(
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

	isMarkedForDeletion := false

	request = FuncAuthenticatorCreatesRequestInitialEmailAuth(
		requestPram,
		AuthenticatorReq{
			User: User{
				InitialEmailAuth: InitialEmailAuth{
					EmailAddress:        input.User.AuthenticatorInitialEmailAuth.EmailAddress,
					CreationDate:        formattedDate,
					CreationTime:        formattedTime,
					LastChangeDate:      formattedDate,
					LastChangeTime:      formattedTime,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
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

func FuncAuthenticatorCreatesRequestInitialSMSAuth(
	requestPram *types.Request,
	input AuthenticatorReq,
) AuthenticatorReq {
	req := AuthenticatorReq{
		User:    input.User,
		APIType: "creates",
		Accepter: []string{
			"InitialSMSAuth",
		},
	}
	return req
}

func AuthenticatorCreatesRequestInitialSMSAuth(
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

	isMarkedForDeletion := false

	request = FuncAuthenticatorCreatesRequestInitialSMSAuth(
		requestPram,
		AuthenticatorReq{
			User: User{
				InitialSMSAuth: InitialSMSAuth{
					MobilePhoneNumber:   input.User.AuthenticatorInitialSMSAuth.MobilePhoneNumber,
					AuthenticationCode:  input.User.AuthenticatorInitialSMSAuth.AuthenticationCode,
					CreationDate:        formattedDate,
					CreationTime:        formattedTime,
					LastChangeDate:      formattedDate,
					LastChangeTime:      formattedTime,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
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

func FuncAuthenticatorCreatesRequestInitialGoogleAccountAuth(
	requestPram *types.Request,
	input AuthenticatorReq,
) AuthenticatorReq {
	req := AuthenticatorReq{
		User:    input.User,
		APIType: "creates",
		Accepter: []string{
			"InitialGoogleAccountAuth",
		},
	}
	return req
}

func AuthenticatorCreatesRequestInitialGoogleAccountAuth(
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

	isMarkedForDeletion := false

	request = FuncAuthenticatorCreatesRequestInitialGoogleAccountAuth(
		requestPram,
		AuthenticatorReq{
			User: User{
				InitialGoogleAccountAuth: InitialGoogleAccountAuth{
					EmailAddress:        input.User.AuthenticatorInitialGoogleAccountAuth.EmailAddress,
					GoogleID:            input.User.AuthenticatorInitialGoogleAccountAuth.GoogleID,
					AccessToken:         input.User.AuthenticatorInitialGoogleAccountAuth.AccessToken,
					CreationDate:        formattedDate,
					CreationTime:        formattedTime,
					LastChangeDate:      formattedDate,
					LastChangeTime:      formattedTime,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
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

func FuncAuthenticatorCreatesRequestInstagramAuth(
	requestPram *types.Request,
	input AuthenticatorReq,
) AuthenticatorReq {
	req := AuthenticatorReq{
		User:    input.User,
		APIType: "creates",
		Accepter: []string{
			"InstagramAuth",
		},
	}
	return req
}

func AuthenticatorCreatesRequestInstagramAuth(
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

	isMarkedForDeletion := false

	request = FuncAuthenticatorCreatesRequestInstagramAuth(
		requestPram,
		AuthenticatorReq{
			User: User{
				InstagramAuth: InstagramAuth{
					UserID:              input.User.AuthenticatorInstagramAuth.UserID,
					InstagramID:         input.User.AuthenticatorInstagramAuth.InstagramID,
					AccessToken:         input.User.AuthenticatorInstagramAuth.AccessToken,
					CreationDate:        formattedDate,
					CreationTime:        formattedTime,
					LastChangeDate:      formattedDate,
					LastChangeTime:      formattedTime,
					IsMarkedForDeletion: &isMarkedForDeletion,
				},
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

	request = FuncAuthenticatorUpdatesRequestUser(
		requestPram,
		AuthenticatorReq{
			User: User{
				UserID:         input.User.UserID,
				Password:       input.User.Password,
				LastLoginDate:  input.User.LastLoginDate,
				LastLoginTime:  input.User.LastLoginTime,
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

func FuncAuthenticatorUpdatesRequestInstagramAuth(
	requestPram *types.Request,
	input AuthenticatorReq,
) AuthenticatorReq {
	req := AuthenticatorReq{
		User:    input.User,
		APIType: "updates",
		Accepter: []string{
			"InstagramAuth",
		},
	}
	return req
}

func AuthenticatorUpdatesRequestInstagramAuth(
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

	request = FuncAuthenticatorUpdatesRequestInstagramAuth(
		requestPram,
		AuthenticatorReq{
			User: User{
				InstagramAuth: InstagramAuth{
					UserID:         input.User.AuthenticatorInstagramAuth.UserID,
					InstagramID:    input.User.AuthenticatorInstagramAuth.InstagramID,
					AccessToken:    input.User.AuthenticatorInstagramAuth.AccessToken,
					LastChangeDate: formattedDate,
					LastChangeTime: formattedTime,
				},
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
