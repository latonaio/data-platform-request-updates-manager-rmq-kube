package apiModuleRuntimesRequestsAddress

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type AddressReq struct {
	Address    Address    `json:"Address"`
	APIType    string     `json:"api_type"`
	Accepter   []string   `json:"accepter"`
}

type Address struct {
	AddressID			*int    `json:"AddressID"`
	ValidityStartDate	string  `json:"ValidityStartDate"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	PostalCode			*string	`json:"PostalCode"`
	LocalSubRegion		string	`json:"LocalSubRegion"`
	LocalRegion			string	`json:"LocalRegion"`
	Country				string	`json:"Country"`
	GlobalRegion		string	`json:"GlobalRegion"`
	TimeZone			string	`json:"TimeZone"`
	District			*string	`json:"District"`
	StreetName			*string	`json:"StreetName"`
	CityName			*string	`json:"CityName"`
	Building			*string	`json:"Building"`
	Floor				*int    `json:"Floor"`
	Room				*int    `json:"Room"`
	XCoordinate 		*float32 `json:"XCoordinate"`
	YCoordinate 		*float32 `json:"YCoordinate"`
	ZCoordinate 		*float32 `json:"ZCoordinate"`
	Site				*int	 `json:"Site"`
	CreationDate        string  `json:"CreationDate"`
	CreationTime        string  `json:"CreationTime"`
	LastChangeDate      string  `json:"LastChangeDate"`
	LastChangeTime      string  `json:"LastChangeTime"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}

func FuncAddressCreatesRequestAddress(
	requestPram *types.Request,
	input AddressReq,
) AddressReq {
	req := AddressReq{
		Address: input.Address,
		APIType: "creates",
		Accepter: []string{
			"Address",
		},
	}
	return req
}

func AddressCreatesRequestAddress(
	requestPram *types.Request,
	input types.AddressSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ADDRESS_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request AddressReq

	request = FuncAddressCreatesRequestAddress(
		requestPram,
		AddressReq{
			Address: Address{
				BusinessPartner:        input.Address.BusinessPartner,
			    AddressID:				input.Address.AddressID,
			    ValidityStartDate:		input.Address.ValidityStartDate,
			    ValidityEndDate:		input.Address.ValidityEndDate,
			    LocalSubRegion:			input.Address.LocalSubRegion,
			    LocalRegion:			input.Address.LocalRegion,
			    Country:				input.Address.Country,
			    GlobalRegion:			input.Address.GlobalRegion,
			    TimeZone:				input.Address.TimeZone,
			    CreationDate:			input.Address.CreationDate,
			    CreationTime:			input.Address.CreationTime,
			    LastChangeDate:			input.Address.LastChangeDate,
			    LastChangeTime:			input.Address.LastChangeTime,
			    IsMarkedForDeletion:	input.Address.IsMarkedForDeletion,
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

func FuncAddressUpdatesRequestAddress(
	requestPram *types.Request,
	input AddressReq,
) AddressReq {
	req := AddressReq{
		Address: input.Address,
		APIType: "updates",
		Accepter: []string{
			"Address",
		},
	}
	return req
}

func AddressUpdatesRequestAddress(
	requestPram *types.Request,
	input types.AddressSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ADDRESS_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request AddressReq

	request = FuncAddressUpdatesRequestAddress(
		requestPram,
		AddressReq{
			Address: Address{
				BusinessPartner:        input.Address.BusinessPartner,
			    AddressID:				input.Address.AddressID,
			    ValidityStartDate:		input.Address.ValidityStartDate,
			    ValidityEndDate:		input.Address.ValidityEndDate,
			    PostalCode:				input.Address.PostalCode,
			    LocalSubRegion:			input.Address.LocalSubRegion,
			    LocalRegion:			input.Address.LocalRegion,
			    Country:				input.Address.Country,
			    GlobalRegion:			input.Address.GlobalRegion,
			    TimeZone:				input.Address.TimeZone,
			    District:				input.Address.District,
			    StreetName:				input.Address.StreetName,
			    CityName:				input.Address.CityName,
			    Building:				input.Address.Building,
			    Floor:					input.Address.Floor,
			    Room:					input.Address.Room,
			    XCoordinate: 			input.Address.XCoordinate,
			    YCoordinate: 			input.Address.YCoordinate,
			    ZCoordinate: 			input.Address.ZCoordinate,
			    Site:		 			input.Address.Site,
			    LastChangeDate:			input.Address.LastChangeDate,
			    LastChangeTime:			input.Address.LastChangeTime,
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
