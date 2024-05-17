package apiModuleRuntimesRequestsSite

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type SiteReq struct {
	Header   Header   `json:"Site"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Site                         *int      `json:"Site"`
	SiteType                     string    `json:"SiteType"`
	SiteOwner                    *int      `json:"SiteOwner"`
	SiteOwnerBusinessPartnerRole *string   `json:"SiteOwnerBusinessPartnerRole"`
	Brand                        *int      `json:"Brand"`
	PersonResponsible            *string   `json:"PersonResponsible"`
	URL                          *string   `json:"URL"`
	ValidityStartDate            string    `json:"ValidityStartDate"`
	ValidityStartTime            string    `json:"ValidityStartTime"`
	ValidityEndDate              string    `json:"ValidityEndDate"`
	ValidityEndTime              string    `json:"ValidityEndTime"`
	DailyOperationStartTime      string    `json:"DailyOperationStartTime"`
	DailyOperationEndTime        string    `json:"DailyOperationEndTime"`
	Description                  string    `json:"Description"`
	LongText                     string    `json:"LongText"`
	Introduction                 *string   `json:"Introduction"`
	OperationRemarks             *string   `json:"OperationRemarks"`
	PhoneNumber                  *string   `json:"PhoneNumber"`
	AvailabilityOfParking        *bool     `json:"AvailabilityOfParking"`
	NumberOfParkingSpaces        *int      `json:"NumberOfParkingSpaces"`
	SuperiorSite                 *int      `json:"SuperiorSite"`
	Project                      *int      `json:"Project"`
	WBSElement                   *int      `json:"WBSElement"`
	Tag1                         *string   `json:"Tag1"`
	Tag2                         *string   `json:"Tag2"`
	Tag3                         *string   `json:"Tag3"`
	Tag4                         *string   `json:"Tag4"`
	CreationDate                 string    `json:"CreationDate"`
	CreationTime                 string    `json:"CreationTime"`
	LastChangeDate               string    `json:"LastChangeDate"`
	LastChangeTime               string    `json:"LastChangeTime"`
	CreateUser                   int       `json:"CreateUser"`
	LastChangeUser               int       `json:"LastChangeUser"`
	IsReleased                   *bool     `json:"IsReleased"`
	IsMarkedForDeletion          *bool     `json:"IsMarkedForDeletion"`
	Partner                      []Partner `json:"Partner"`
	Address                      []Address `json:"Address"`
}

type Partner struct {
	Site                    int     `json:"Site"`
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
	Site           int      `json:"Site"`
	AddressID      int      `json:"AddressID"`
	PostalCode     string   `json:"PostalCode"`
	LocalSubRegion string   `json:"LocalSubRegion"`
	LocalRegion    string   `json:"LocalRegion"`
	Country        string   `json:"Country"`
	GlobalRegion   string   `json:"GlobalRegion"`
	TimeZone       string   `json:"TimeZone"`
	District       *string  `json:"District"`
	StreetName     *string  `json:"StreetName"`
	CityName       *string  `json:"CityName"`
	Building       *string  `json:"Building"`
	Floor          *int     `json:"Floor"`
	Room           *int     `json:"Room"`
	XCoordinate    *float32 `json:"XCoordinate"`
	YCoordinate    *float32 `json:"YCoordinate"`
	ZCoordinate    *float32 `json:"ZCoordinate"`
}

func FuncSiteCreatesRequestAll(
	requestPram *types.Request,
	input SiteReq,
) SiteReq {
	req := SiteReq{
		Header:  input.Header,
		APIType: "creates",
		Accepter: []string{
			"Header",
			//"Address",
		},
	}
	return req
}

func SiteCreatesRequestAll(
	requestPram *types.Request,
	input types.SiteSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SITE_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request SiteReq

	isReleased := false
	isMarkedForDeletion := false

	introduction := "Test Introduction"

	request = FuncSiteCreatesRequestAll(
		requestPram,
		SiteReq{
			Header: Header{
				Site:                         nil,
				SiteType:                     input.Header.SiteType,
				SiteOwner:                    input.Header.SiteOwner,
				SiteOwnerBusinessPartnerRole: input.Header.SiteOwnerBusinessPartnerRole,
				Brand:                        input.Header.Brand,
				PersonResponsible:            input.Header.PersonResponsible,
				URL:                          input.Header.URL,
				ValidityStartDate:            formattedDate,
				ValidityStartTime:            formattedTime,
				ValidityEndDate:              "9999-12-31",
				ValidityEndTime:              "23:59:59",
				DailyOperationStartTime:      input.Header.DailyOperationStartTime,
				DailyOperationEndTime:        input.Header.DailyOperationEndTime,
				Description:                  input.Header.Description,
				LongText:                     input.Header.LongText,
				Introduction:                 &introduction,
				OperationRemarks:             input.Header.OperationRemarks,
				PhoneNumber:                  input.Header.PhoneNumber,
				AvailabilityOfParking:        input.Header.AvailabilityOfParking,
				NumberOfParkingSpaces:        input.Header.NumberOfParkingSpaces,
				SuperiorSite:                 input.Header.SuperiorSite,
				Tag1:                         input.Header.Tag1,
				Tag2:                         input.Header.Tag2,
				Tag3:                         input.Header.Tag3,
				Tag4:                         input.Header.Tag4,
				CreationDate:                 formattedDate,
				CreationTime:                 formattedTime,
				LastChangeDate:               formattedDate,
				LastChangeTime:               formattedTime,
				CreateUser:                   input.Header.CreateUser,
				LastChangeUser:               input.Header.LastChangeUser,
				IsReleased:                   &isReleased,
				IsMarkedForDeletion:          &isMarkedForDeletion,
				//Address:	[]Address{
				//	{
				//		AddressID:			input.Header.SiteAddress[0].AddressID,
				//		PostalCode:			input.Header.SiteAddress[0].PostalCode,
				//		LocalSubRegion:		input.Header.SiteAddress[0].LocalSubRegion,
				//		LocalRegion:		input.Header.SiteAddress[0].LocalRegion,
				//		Country:			input.Header.SiteAddress[0].Country,
				//		GlobalRegion:		input.Header.SiteAddress[0].GlobalRegion,
				//		TimeZone:			input.Header.SiteAddress[0].TimeZone,
				//		District:			input.Header.SiteAddress[0].District,
				//		StreetName:			input.Header.SiteAddress[0].StreetName,
				//		CityName:			input.Header.SiteAddress[0].CityName,
				//		Building:			input.Header.SiteAddress[0].Building,
				//		XCoordinate:		input.Header.SiteAddress[0].XCoordinate,
				//		YCoordinate:		input.Header.SiteAddress[0].YCoordinate,
				//		ZCoordinate:		input.Header.SiteAddress[0].ZCoordinate,
				//	},
				//},
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

func FuncSiteUpdatesRequestHeader(
	requestPram *types.Request,
	input SiteReq,
) SiteReq {
	req := SiteReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func SiteUpdatesRequestHeader(
	requestPram *types.Request,
	input types.SiteSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SITE_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request SiteReq

	request = FuncSiteUpdatesRequestHeader(
		requestPram,
		SiteReq{
			Header: Header{
				Site:                         input.Header.Site,
				SiteType:                     input.Header.SiteType,
				SiteOwnerBusinessPartnerRole: input.Header.SiteOwnerBusinessPartnerRole,
				Brand:                        input.Header.Brand,
				PersonResponsible:            input.Header.PersonResponsible,
				URL:                          input.Header.URL,
				ValidityStartDate:            input.Header.ValidityStartDate,
				ValidityStartTime:            input.Header.ValidityStartTime,
				ValidityEndDate:              input.Header.ValidityEndDate,
				ValidityEndTime:              input.Header.ValidityEndTime,
				DailyOperationStartTime:      input.Header.DailyOperationStartTime,
				DailyOperationEndTime:        input.Header.DailyOperationEndTime,
				Description:                  input.Header.Description,
				LongText:                     input.Header.LongText,
				Introduction:                 input.Header.Introduction,
				OperationRemarks:             input.Header.OperationRemarks,
				PhoneNumber:                  input.Header.PhoneNumber,
				AvailabilityOfParking:        input.Header.AvailabilityOfParking,
				NumberOfParkingSpaces:        input.Header.NumberOfParkingSpaces,
				SuperiorSite:                 input.Header.SuperiorSite,
				Tag1:                         input.Header.Tag1,
				Tag2:                         input.Header.Tag2,
				Tag3:                         input.Header.Tag3,
				Tag4:                         input.Header.Tag4,
				LastChangeDate:               formattedDate,
				LastChangeTime:               formattedTime,
				LastChangeUser:               input.Header.LastChangeUser,
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
