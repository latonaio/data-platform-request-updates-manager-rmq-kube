package apiModuleRuntimesRequestsShop

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type ShopReq struct {
	Header          Header          `json:"Shop"`
	APIType         string          `json:"api_type"`
	Accepter        []string        `json:"accepter"`
}

type Header struct {
	Shop							*int	`json:"Shop"`
	ShopType						string	`json:"ShopType"`
	ShopOwner						int		`json:"ShopOwner"`
	ShopOwnerBusinessPartnerRole	string	`json:"ShopOwnerBusinessPartnerRole"`
	Brand							*int	`json:"Brand"`
	PersonResponsible				string	`json:"PersonResponsible"`
	ValidityStartDate				string	`json:"ValidityStartDate"`
	ValidityStartTime				string	`json:"ValidityStartTime"`
	ValidityEndDate					string	`json:"ValidityEndDate"`
	ValidityEndTime					string	`json:"ValidityEndTime"`
	DailyOperationStartTime			string	`json:"DailyOperationStartTime"`
	DailyOperationEndTime			string	`json:"DailyOperationEndTime"`
	Description						string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Introduction					*string	`json:"Introduction"`
	OperationRemarks				*string	`json:"OperationRemarks"`
	PhoneNumber						*string	`json:"PhoneNumber"`
	AvailabilityOfParking			*bool	`json:"AvailabilityOfParking"`
	NumberOfParkingSpaces			*int	`json:"NumberOfParkingSpaces"`
	Site							int		`json:"Site"`
	Project							*int	`json:"Project"`
	WBSElement						*int	`json:"WBSElement"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	PointConsumptionType      		string  `json:"PointConsumptionType"`
	CreationDate					string	`json:"CreationDate"`
	CreationTime					string	`json:"CreationTime"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	CreateUser						int		`json:"CreateUser"`
	LastChangeUser					int		`json:"LastChangeUser"`
	IsReleased						*bool	`json:"IsReleased"`
	IsMarkedForDeletion				*bool	`json:"IsMarkedForDeletion"`
	Partner             			[]Partner `json:"Partner"`
	Address             			[]Address `json:"Address"`
}

type Partner struct {
	Shop                 	int     `json:"Shop"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
}

type Address struct {
	Shop     		int     	`json:"Shop"`
	AddressID   	int     	`json:"AddressID"`
	PostalCode  	string 		`json:"PostalCode"`
	LocalSubRegion 	string 		`json:"LocalSubRegion"`
	LocalRegion 	string 		`json:"LocalRegion"`
	Country     	string 		`json:"Country"`
	GlobalRegion   	string 		`json:"GlobalRegion"`
	TimeZone   		string 		`json:"TimeZone"`
	District    	*string 	`json:"District"`
	StreetName  	*string 	`json:"StreetName"`
	CityName    	*string 	`json:"CityName"`
	Building    	*string 	`json:"Building"`
	Floor       	*int		`json:"Floor"`
	Room        	*int		`json:"Room"`
	XCoordinate 	*float32	`json:"XCoordinate"`
	YCoordinate 	*float32	`json:"YCoordinate"`
	ZCoordinate 	*float32	`json:"ZCoordinate"`
	Site			*int		`json:"Site"`
}

func FuncShopCreatesRequestAll(
	requestPram *types.Request,
	input ShopReq,
) ShopReq {
	req := ShopReq{
		Header: input.Header,
		APIType: "creates",
		Accepter: []string{
			"Header",
			"Address",
		},
	}
	return req
}

func ShopCreatesRequestAll(
	requestPram *types.Request,
	input types.ShopSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SHOP_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request ShopReq

	isReleased := false
	isMarkedForDeletion := false

	request = FuncShopCreatesRequestAll(
		requestPram,
		ShopReq{
			Header: Header{
				Shop:				    		nil,
				ShopType:			    		input.Header.ShopType,
				ShopOwner:			    		input.Header.ShopOwner,
				ShopOwnerBusinessPartnerRole:	input.Header.ShopOwnerBusinessPartnerRole,
				Brand:							input.Header.Brand,
				PersonResponsible:	    		input.Header.PersonResponsible,
				ValidityStartDate:	    		formattedDate,
				ValidityStartTime:	    		formattedTime,
				ValidityEndDate:	    		"9999-12-31",
				ValidityEndTime:	    		"23:59:59",
				DailyOperationStartTime:		input.Header.DailyOperationStartTime,
				DailyOperationEndTime:		    input.Header.DailyOperationEndTime,
				Description:		    		input.Header.Description,
				LongText:	            		input.Header.LongText,
				Introduction:		    		input.Header.Introduction,
				OperationRemarks:		    	input.Header.OperationRemarks,
				PhoneNumber:		    		input.Header.PhoneNumber,
				AvailabilityOfParking:		    input.Header.AvailabilityOfParking,
				NumberOfParkingSpaces:		    input.Header.NumberOfParkingSpaces,
				Site:        					input.Header.Site,
				Tag1:		            		input.Header.Tag1,
				Tag2:		            		input.Header.Tag2,
				Tag3:		            		input.Header.Tag3,
				Tag4:		            		input.Header.Tag4,
                PointConsumptionType:     		input.Header.PointConsumptionType,
				CreationDate:		    		formattedDate,
				CreationTime:		    		formattedTime,
                LastChangeDate:         		formattedDate,
                LastChangeTime:         		formattedTime,
				CreateUser:                	    input.Header.CreateUser,
				LastChangeUser:            	    input.Header.LastChangeUser,
                IsReleased:    					&isReleased,
                IsMarkedForDeletion:    		&isMarkedForDeletion,
				Address:	[]Address{
					{
						AddressID:			input.Header.ShopAddress[0].AddressID,
						PostalCode:			input.Header.ShopAddress[0].PostalCode,
						LocalSubRegion:		input.Header.ShopAddress[0].LocalSubRegion,
						LocalRegion:		input.Header.ShopAddress[0].LocalRegion,
						Country:			input.Header.ShopAddress[0].Country,
						GlobalRegion:		input.Header.ShopAddress[0].GlobalRegion,
						TimeZone:			input.Header.ShopAddress[0].TimeZone,
						District:			input.Header.ShopAddress[0].District,
						StreetName:			input.Header.ShopAddress[0].StreetName,
						CityName:			input.Header.ShopAddress[0].CityName,
						Building:			input.Header.ShopAddress[0].Building,
						XCoordinate:		input.Header.ShopAddress[0].XCoordinate,
						YCoordinate:		input.Header.ShopAddress[0].YCoordinate,
						ZCoordinate:		input.Header.ShopAddress[0].ZCoordinate,
						Site:        		input.Header.ShopAddress[0].Site,
					},
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

func FuncShopUpdatesRequestHeader(
	requestPram *types.Request,
	input ShopReq,
) ShopReq {
	req := ShopReq{
		Header: input.Header,
		APIType: "updates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func ShopUpdatesRequestHeader(
	requestPram *types.Request,
	input types.ShopSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SHOP_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request ShopReq

	request = FuncShopUpdatesRequestHeader(
		requestPram,
		ShopReq{
			Header: Header{
				Shop:				    		input.Header.Shop,
				ShopType:			    		input.Header.ShopType,
				ShopOwnerBusinessPartnerRole:	input.Header.ShopOwnerBusinessPartnerRole,
				Brand:							input.Header.Brand,
				PersonResponsible:	    		input.Header.PersonResponsible,
				ValidityStartDate:	    		input.Header.ValidityStartDate,
				ValidityStartTime:	    		input.Header.ValidityStartTime,
				ValidityEndDate:	    		input.Header.ValidityEndDate,
				ValidityEndTime:	    		input.Header.ValidityEndTime,
				DailyOperationStartTime:		input.Header.DailyOperationStartTime,
				DailyOperationEndTime:		    input.Header.DailyOperationEndTime,
				Description:		    		input.Header.Description,
				LongText:	            		input.Header.LongText,
				Introduction:		    		input.Header.Introduction,
				OperationRemarks:		    	input.Header.OperationRemarks,
				PhoneNumber:		    		input.Header.PhoneNumber,
				AvailabilityOfParking:		    input.Header.AvailabilityOfParking,
				NumberOfParkingSpaces:		    input.Header.NumberOfParkingSpaces,
				Site:        					input.Header.Site,
				Tag1:		            		input.Header.Tag1,
				Tag2:		            		input.Header.Tag2,
				Tag3:		            		input.Header.Tag3,
				Tag4:		            		input.Header.Tag4,
                PointConsumptionType:     		input.Header.PointConsumptionType,
                LastChangeDate:         		formattedDate,
                LastChangeTime:         		formattedTime,
				LastChangeUser:            	    input.Header.LastChangeUser,
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
