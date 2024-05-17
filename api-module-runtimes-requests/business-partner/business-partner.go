package apiModuleRuntimesRequestsBusinessPartner

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type BusinessPartnerReq struct {
	General  General  `json:"BusinessPartner"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type General struct {
	BusinessPartner               *int                      `json:"BusinessPartner"`
	BusinessPartnerType           string                    `json:"BusinessPartnerType"`
	BusinessPartnerFullName       *string                   `json:"BusinessPartnerFullName"`
	BusinessPartnerName           string                    `json:"BusinessPartnerName"`
	Industry                      *string                   `json:"Industry"`
	LegalEntityRegistration       *string                   `json:"LegalEntityRegistration"`
	Country                       string                    `json:"Country"`
	Language                      string                    `json:"Language"`
	Currency                      *string                   `json:"Currency"`
	Representative                *string                   `json:"Representative"`
	PhoneNumber                   *string                   `json:"PhoneNumber"`
	OrganizationBPName1           *string                   `json:"OrganizationBPName1"`
	OrganizationBPName2           *string                   `json:"OrganizationBPName2"`
	OrganizationBPName3           *string                   `json:"OrganizationBPName3"`
	OrganizationBPName4           *string                   `json:"OrganizationBPName4"`
	Tag1                          *string                   `json:"Tag1"`
	Tag2                          *string                   `json:"Tag2"`
	Tag3                          *string                   `json:"Tag3"`
	Tag4                          *string                   `json:"Tag4"`
	OrganizationFoundationDate    *string                   `json:"OrganizationFoundationDate"`
	OrganizationLiquidationDate   *string                   `json:"OrganizationLiquidationDate"`
	BusinessPartnerBirthplaceName *string                   `json:"BusinessPartnerBirthplaceName"`
	BusinessPartnerDeathDate      *string                   `json:"BusinessPartnerDeathDate"`
	GroupBusinessPartnerName1     *string                   `json:"GroupBusinessPartnerName1"`
	GroupBusinessPartnerName2     *string                   `json:"GroupBusinessPartnerName2"`
	AddressID                     *int                      `json:"AddressID"`
	BusinessPartnerIDByExtSystem  *string                   `json:"BusinessPartnerIDByExtSystem"`
	BusinessPartnerIsBlocked      *bool                     `json:"BusinessPartnerIsBlocked"`
	CertificateAuthorityChain     *string                   `json:"CertificateAuthorityChain"`
	UsageControlChain             *string                   `json:"UsageControlChain"`
	CreationDate                  string                    `json:"CreationDate"`
	LastChangeDate                string                    `json:"LastChangeDate"`
	IsMarkedForDeletion           *bool                     `json:"IsMarkedForDeletion"`
	Role                          []Role                    `json:"Role"`
	Person                        []Person                  `json:"Person"`
	Address                       []Address                 `json:"Address"`
	Rank                          []Rank                    `json:"Rank"`
	PersonMobilePhoneAuth         []PersonMobilePhoneAuth   `json:"PersonMobilePhoneAuth"`
	PersonGoogleAccountAuth       []PersonGoogleAccountAuth `json:"PersonGoogleAccountAuth"`
	PersonInstagramAuth           []PersonInstagramAuth     `json:"PersonInstagramAuth"`
}

type Role struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	BusinessPartnerRole string `json:"BusinessPartnerRole"`
	ValidityStartDate   string `json:"ValidityStartDate"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type Person struct {
	BusinessPartner          int     `json:"BusinessPartner"`
	BusinessPartnerType      string  `json:"BusinessPartnerType"`
	FirstName                *string `json:"FirstName"`
	LastName                 *string `json:"LastName"`
	FullName                 *string `json:"FullName"`
	MiddleName               *string `json:"MiddleName"`
	NickName                 string  `json:"NickName"`
	Gender                   string  `json:"Gender"`
	Language                 string  `json:"Language"`
	CorrespondenceLanguage   *string `json:"CorrespondenceLanguage"`
	BirthDate                *string `json:"BirthDate"`
	Nationality              string  `json:"Nationality"`
	EmailAddress             *string `json:"EmailAddress"`
	MobilePhoneNumber        *string `json:"MobilePhoneNumber"`
	ProfileComment           *string `json:"ProfileComment"`
	PreferableLocalSubRegion string  `json:"PreferableLocalSubRegion"`
	PreferableLocalRegion    string  `json:"PreferableLocalRegion"`
	PreferableCountry        string  `json:"PreferableCountry"`
	ActPurpose               string  `json:"ActPurpose"`
	CreationDate             string  `json:"CreationDate"`
	LastChangeDate           string  `json:"LastChangeDate"`
	IsMarkedForDeletion      *bool   `json:"IsMarkedForDeletion"`
}

type Address struct {
	BusinessPartner int      `json:"BusinessPartner"`
	AddressID       int      `json:"AddressID"`
	PostalCode      string   `json:"PostalCode"`
	LocalSubRegion  string   `json:"LocalSubRegion"`
	LocalRegion     string   `json:"LocalRegion"`
	Country         string   `json:"Country"`
	GlobalRegion    string   `json:"GlobalRegion"`
	TimeZone        string   `json:"TimeZone"`
	District        *string  `json:"District"`
	StreetName      *string  `json:"StreetName"`
	CityName        *string  `json:"CityName"`
	Building        *string  `json:"Building"`
	Floor           *int     `json:"Floor"`
	Room            *int     `json:"Room"`
	XCoordinate     *float32 `json:"XCoordinate"`
	YCoordinate     *float32 `json:"YCoordinate"`
	ZCoordinate     *float32 `json:"ZCoordinate"`
	Site            *int     `json:"Site"`
}

type Rank struct {
	BusinessPartner     int    `json:"BusinessPartner"`
	RankType            string `json:"RankType"`
	Rank                int    `json:"Rank"`
	ValidityStartDate   string `json:"ValidityStartDate"`
	ValidityEndDate     string `json:"ValidityEndDate"`
	CreationDate        string `json:"CreationDate"`
	LastChangeDate      string `json:"LastChangeDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

type PersonMobilePhoneAuth struct {
	BusinessPartner            int    `json:"BusinessPartner"`
	BusinessPartnerType        string `json:"BusinessPartnerType"`
	MobilePhoneNumber          string `json:"MobilePhoneNumber"`
	MobilePhoneIsAuthenticated bool   `json:"MobilePhoneIsAuthenticated"`
	CreationDate               string `json:"CreationDate"`
	LastChangeDate             string `json:"LastChangeDate"`
	IsMarkedForDeletion        *bool  `json:"IsMarkedForDeletion"`
}

type PersonGoogleAccountAuth struct {
	BusinessPartner              int    `json:"BusinessPartner"`
	BusinessPartnerType          string `json:"BusinessPartnerType"`
	EmailAddress                 string `json:"EmailAddress"`
	GoogleID                     string `json:"GoogleID"`
	GoogleAccountIsAuthenticated bool   `json:"GoogleAccountIsAuthenticated"`
	CreationDate                 string `json:"CreationDate"`
	LastChangeDate               string `json:"LastChangeDate"`
	IsMarkedForDeletion          *bool  `json:"IsMarkedForDeletion"`
}

type PersonInstagramAuth struct {
	BusinessPartner            int     `json:"BusinessPartner"`
	BusinessPartnerType        string  `json:"BusinessPartnerType"`
	InstagramID                string  `json:"InstagramID"`
	InstagramUserName          string  `json:"InstagramUserName"`
	InstagramHasProfilePicture bool    `json:"InstagramHasProfilePicture"`
	InstagramProfilePictureURI *string `json:"InstagramProfilePictureURI"`
	InstagramIsPrivate         bool    `json:"InstagramIsPrivate"`
	InstagramIsPublished       bool    `json:"InstagramIsPublished"`
	InstagramIsAuthenticated   bool    `json:"InstagramIsAuthenticated"`
	CreationDate               string  `json:"CreationDate"`
	LastChangeDate             string  `json:"LastChangeDate"`
	IsMarkedForDeletion        *bool   `json:"IsMarkedForDeletion"`
}

func FuncBusinessPartnerCreatesRequestAll(
	requestPram *types.Request,
	input BusinessPartnerReq,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		General: input.General,
		APIType: "creates",
		Accepter: []string{
			"General",
			"Role",
			"Person",
			"Address",
			"Rank",
		},
	}
	return req
}

func BusinessPartnerCreatesRequestAll(
	requestPram *types.Request,
	input types.BusinessPartnerSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")

	var request BusinessPartnerReq

	isMarkedForDeletion := false

	request = FuncBusinessPartnerCreatesRequestAll(
		requestPram,
		BusinessPartnerReq{
			General: General{
				BusinessPartner:     nil,
				BusinessPartnerType: "02",
				BusinessPartnerName: input.General.BusinessPartnerName,
				Country:             input.General.Country,
				Language:            input.General.Language,
				AddressID:           input.General.AddressID,
				CreationDate:        formattedDate,
				LastChangeDate:      formattedDate,
				IsMarkedForDeletion: &isMarkedForDeletion,
				Role: []Role{
					{
						BusinessPartnerRole: "CONSUMER",
						ValidityStartDate:   formattedDate,
						ValidityEndDate:     "9999-12-31",
						CreationDate:        formattedDate,
						LastChangeDate:      formattedDate,
						IsMarkedForDeletion: &isMarkedForDeletion,
					},
				},
				Person: []Person{
					{
						BusinessPartnerType:      "02",
						FirstName:                input.General.BusinessPartnerPerson[0].FirstName,
						LastName:                 input.General.BusinessPartnerPerson[0].LastName,
						FullName:                 input.General.BusinessPartnerPerson[0].FullName,
						NickName:                 input.General.BusinessPartnerPerson[0].NickName,
						Gender:                   input.General.BusinessPartnerPerson[0].Gender,
						Language:                 input.General.BusinessPartnerPerson[0].Language,
						Nationality:              input.General.BusinessPartnerPerson[0].Nationality,
						EmailAddress:             input.General.BusinessPartnerPerson[0].EmailAddress,
						MobilePhoneNumber:        input.General.BusinessPartnerPerson[0].MobilePhoneNumber,
						PreferableLocalSubRegion: input.General.BusinessPartnerPerson[0].PreferableLocalSubRegion,
						PreferableLocalRegion:    input.General.BusinessPartnerPerson[0].PreferableLocalRegion,
						PreferableCountry:        input.General.BusinessPartnerPerson[0].PreferableCountry,
						ActPurpose:               input.General.BusinessPartnerPerson[0].ActPurpose,
						CreationDate:             formattedDate,
						LastChangeDate:           formattedDate,
						IsMarkedForDeletion:      &isMarkedForDeletion,
					},
				},
				Address: []Address{
					{
						AddressID:      input.General.BusinessPartnerAddress[0].AddressID,
						LocalSubRegion: input.General.BusinessPartnerAddress[0].LocalSubRegion,
						LocalRegion:    input.General.BusinessPartnerAddress[0].LocalRegion,
						Country:        input.General.BusinessPartnerAddress[0].Country,
						GlobalRegion:   input.General.BusinessPartnerAddress[0].GlobalRegion,
						TimeZone:       input.General.BusinessPartnerAddress[0].TimeZone,
					},
				},
				Rank: []Rank{
					{
						RankType:            "0001",
						Rank:                1,
						ValidityStartDate:   formattedDate,
						ValidityEndDate:     "9999-12-31",
						CreationDate:        formattedDate,
						LastChangeDate:      formattedDate,
						IsMarkedForDeletion: &isMarkedForDeletion,
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

func FuncBusinessPartnerCreatesRequestRole(
	requestPram *types.Request,
	input BusinessPartnerReq,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		General: input.General,
		APIType: "creates",
		Accepter: []string{
			"Role",
		},
	}
	return req
}

func BusinessPartnerCreatesRequestRole(
	requestPram *types.Request,
	input types.BusinessPartnerSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")

	var request BusinessPartnerReq

	isMarkedForDeletion := false

	request = FuncBusinessPartnerCreatesRequestRole(
		requestPram,
		BusinessPartnerReq{
			General: General{
				BusinessPartner: input.General.BusinessPartner,
				Role: []Role{
					{
						BusinessPartnerRole: input.General.BusinessPartnerRole[0].BusinessPartnerRole,
						ValidityStartDate:   formattedDate,
						ValidityEndDate:     "9999-12-31",
						CreationDate:        formattedDate,
						LastChangeDate:      formattedDate,
						IsMarkedForDeletion: &isMarkedForDeletion,
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

func FuncBusinessPartnerCreatesRequestPersonMobilePhoneAuth(
	requestPram *types.Request,
	input BusinessPartnerReq,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		General: input.General,
		APIType: "creates",
		Accepter: []string{
			"PersonMobilePhoneAuth",
		},
	}
	return req
}

func BusinessPartnerCreatesRequestPersonMobilePhoneAuth(
	requestPram *types.Request,
	input types.BusinessPartnerSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")

	var request BusinessPartnerReq

	isMarkedForDeletion := false

	request = FuncBusinessPartnerCreatesRequestPersonMobilePhoneAuth(
		requestPram,
		BusinessPartnerReq{
			General: General{
				BusinessPartner: input.General.BusinessPartner,
				PersonMobilePhoneAuth: []PersonMobilePhoneAuth{
					{
						BusinessPartnerType:        "02",
						MobilePhoneNumber:          input.General.BusinessPartnerPersonMobilePhoneAuth[0].MobilePhoneNumber,
						MobilePhoneIsAuthenticated: input.General.BusinessPartnerPersonMobilePhoneAuth[0].MobilePhoneIsAuthenticated,
						CreationDate:               formattedDate,
						LastChangeDate:             formattedDate,
						IsMarkedForDeletion:        &isMarkedForDeletion,
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

func FuncBusinessPartnerCreatesRequestPersonGoogleAccountAuth(
	requestPram *types.Request,
	input BusinessPartnerReq,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		General: input.General,
		APIType: "creates",
		Accepter: []string{
			"PersonGoogleAccountAuth",
		},
	}
	return req
}

func BusinessPartnerCreatesRequestPersonGoogleAccountAuth(
	requestPram *types.Request,
	input types.BusinessPartnerSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")

	var request BusinessPartnerReq

	isMarkedForDeletion := false

	request = FuncBusinessPartnerCreatesRequestPersonGoogleAccountAuth(
		requestPram,
		BusinessPartnerReq{
			General: General{
				BusinessPartner: input.General.BusinessPartner,
				PersonGoogleAccountAuth: []PersonGoogleAccountAuth{
					{
						BusinessPartnerType:          "02",
						EmailAddress:                 input.General.BusinessPartnerPersonGoogleAccountAuth[0].EmailAddress,
						GoogleID:                     input.General.BusinessPartnerPersonGoogleAccountAuth[0].GoogleID,
						GoogleAccountIsAuthenticated: input.General.BusinessPartnerPersonGoogleAccountAuth[0].GoogleAccountIsAuthenticated,
						CreationDate:                 formattedDate,
						LastChangeDate:               formattedDate,
						IsMarkedForDeletion:          &isMarkedForDeletion,
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

func FuncBusinessPartnerCreatesRequestPersonInstagramAuth(
	requestPram *types.Request,
	input BusinessPartnerReq,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		General: input.General,
		APIType: "creates",
		Accepter: []string{
			"PersonInstagramAuth",
		},
	}
	return req
}

func BusinessPartnerCreatesRequestPersonInstagramAuth(
	requestPram *types.Request,
	input types.BusinessPartnerSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")

	var request BusinessPartnerReq

	isMarkedForDeletion := false

	request = FuncBusinessPartnerCreatesRequestPersonInstagramAuth(
		requestPram,
		BusinessPartnerReq{
			General: General{
				BusinessPartner: input.General.BusinessPartner,
				PersonInstagramAuth: []PersonInstagramAuth{
					{
						BusinessPartnerType:        "02",
						InstagramID:                input.General.BusinessPartnerPersonInstagramAuth[0].InstagramID,
						InstagramUserName:          input.General.BusinessPartnerPersonInstagramAuth[0].InstagramUserName,
						InstagramHasProfilePicture: input.General.BusinessPartnerPersonInstagramAuth[0].InstagramHasProfilePicture,
						InstagramProfilePictureURI: input.General.BusinessPartnerPersonInstagramAuth[0].InstagramProfilePictureURI,
						InstagramIsPrivate:         input.General.BusinessPartnerPersonInstagramAuth[0].InstagramIsPrivate,
						InstagramIsPublished:       input.General.BusinessPartnerPersonInstagramAuth[0].InstagramIsPublished,
						InstagramIsAuthenticated:   input.General.BusinessPartnerPersonInstagramAuth[0].InstagramIsAuthenticated,
						CreationDate:               formattedDate,
						LastChangeDate:             formattedDate,
						IsMarkedForDeletion:        &isMarkedForDeletion,
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

func FuncBusinessPartnerUpdatesRequestPerson(
	requestPram *types.Request,
	input BusinessPartnerReq,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		General: input.General,
		APIType: "updates",
		Accepter: []string{
			"Person",
		},
	}
	return req
}

func BusinessPartnerUpdatesRequestPerson(
	requestPram *types.Request,
	input types.BusinessPartnerSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")

	var request BusinessPartnerReq

	isMarkedForDeletion := false

	request = FuncBusinessPartnerUpdatesRequestPerson(
		requestPram,
		BusinessPartnerReq{
			General: General{
				BusinessPartner: input.General.BusinessPartner,
				Person: []Person{
					{
						FirstName:                input.General.BusinessPartnerPerson[0].FirstName,
						LastName:                 input.General.BusinessPartnerPerson[0].LastName,
						FullName:                 input.General.BusinessPartnerPerson[0].FullName,
						NickName:                 input.General.BusinessPartnerPerson[0].NickName,
						Gender:                   input.General.BusinessPartnerPerson[0].Gender,
						Language:                 input.General.BusinessPartnerPerson[0].Language,
						BirthDate:                input.General.BusinessPartnerPerson[0].BirthDate,
						Nationality:              input.General.BusinessPartnerPerson[0].Nationality,
						EmailAddress:             input.General.BusinessPartnerPerson[0].EmailAddress,
						MobilePhoneNumber:        input.General.BusinessPartnerPerson[0].MobilePhoneNumber,
						ProfileComment:           input.General.BusinessPartnerPerson[0].ProfileComment,
						PreferableLocalSubRegion: input.General.BusinessPartnerPerson[0].PreferableLocalSubRegion,
						PreferableLocalRegion:    input.General.BusinessPartnerPerson[0].PreferableLocalRegion,
						PreferableCountry:        input.General.BusinessPartnerPerson[0].PreferableCountry,
						ActPurpose:               input.General.BusinessPartnerPerson[0].ActPurpose,
						LastChangeDate:           formattedDate,
						IsMarkedForDeletion:      &isMarkedForDeletion,
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

func FuncBusinessPartnerUpdatesRequestAddress(
	requestPram *types.Request,
	input BusinessPartnerReq,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		General: input.General,
		APIType: "updates",
		Accepter: []string{
			"Address",
		},
	}
	return req
}

func BusinessPartnerUpdatesRequestAddress(
	requestPram *types.Request,
	input types.BusinessPartnerSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "updates"

	var request BusinessPartnerReq

	request = FuncBusinessPartnerUpdatesRequestAddress(
		requestPram,
		BusinessPartnerReq{
			General: General{
				BusinessPartner: input.General.BusinessPartner,
				Address: []Address{
					{
						AddressID:      input.General.BusinessPartnerAddress[0].AddressID,
						PostalCode:     input.General.BusinessPartnerAddress[0].PostalCode,
						LocalSubRegion: input.General.BusinessPartnerAddress[0].LocalSubRegion,
						LocalRegion:    input.General.BusinessPartnerAddress[0].LocalRegion,
						Country:        input.General.BusinessPartnerAddress[0].Country,
						GlobalRegion:   input.General.BusinessPartnerAddress[0].GlobalRegion,
						TimeZone:       input.General.BusinessPartnerAddress[0].TimeZone,
						District:       input.General.BusinessPartnerAddress[0].District,
						StreetName:     input.General.BusinessPartnerAddress[0].StreetName,
						CityName:       input.General.BusinessPartnerAddress[0].CityName,
						Building:       input.General.BusinessPartnerAddress[0].Building,
						Floor:          input.General.BusinessPartnerAddress[0].Floor,
						Room:           input.General.BusinessPartnerAddress[0].Room,
						XCoordinate:    input.General.BusinessPartnerAddress[0].XCoordinate,
						YCoordinate:    input.General.BusinessPartnerAddress[0].YCoordinate,
						ZCoordinate:    input.General.BusinessPartnerAddress[0].ZCoordinate,
						Site:           input.General.BusinessPartnerAddress[0].Site,
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

func FuncBusinessPartnerUpdatesRequestRank(
	requestPram *types.Request,
	input BusinessPartnerReq,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		General: input.General,
		APIType: "updates",
		Accepter: []string{
			"Rank",
		},
	}
	return req
}

func BusinessPartnerUpdatesRequestRank(
	requestPram *types.Request,
	input types.BusinessPartnerSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")

	var request BusinessPartnerReq

	isMarkedForDeletion := false

	request = FuncBusinessPartnerUpdatesRequestRank(
		requestPram,
		BusinessPartnerReq{
			General: General{
				BusinessPartner: input.General.BusinessPartner,
				Rank: []Rank{
					{
						RankType:            input.General.BusinessPartnerRank[0].RankType,
						Rank:                input.General.BusinessPartnerRank[0].Rank,
						ValidityStartDate:   input.General.BusinessPartnerRank[0].ValidityStartDate,
						ValidityEndDate:     input.General.BusinessPartnerRank[0].ValidityEndDate,
						LastChangeDate:      formattedDate,
						IsMarkedForDeletion: &isMarkedForDeletion,
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

func FuncBusinessPartnerUpdatesRequestPersonMobilePhoneAuth(
	requestPram *types.Request,
	input BusinessPartnerReq,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		General: input.General,
		APIType: "updates",
		Accepter: []string{
			"PersonMobilePhoneAuth",
		},
	}
	return req
}

func BusinessPartnerUpdatesRequestPersonMobilePhoneAuth(
	requestPram *types.Request,
	input types.BusinessPartnerSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")

	var request BusinessPartnerReq

	isMarkedForDeletion := false

	request = FuncBusinessPartnerUpdatesRequestPersonMobilePhoneAuth(
		requestPram,
		BusinessPartnerReq{
			General: General{
				BusinessPartner: input.General.BusinessPartner,
				PersonMobilePhoneAuth: []PersonMobilePhoneAuth{
					{
						MobilePhoneIsAuthenticated: input.General.BusinessPartnerPersonMobilePhoneAuth[0].MobilePhoneIsAuthenticated,
						LastChangeDate:             formattedDate,
						IsMarkedForDeletion:        &isMarkedForDeletion,
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

func FuncBusinessPartnerUpdatesRequestPersonGoogleAccountAuth(
	requestPram *types.Request,
	input BusinessPartnerReq,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		General: input.General,
		APIType: "updates",
		Accepter: []string{
			"PersonMobilePhoneAuth",
		},
	}
	return req
}

func FuncBusinessPartnerUpdatesRequestPersonInstagramAuth(
	requestPram *types.Request,
	input BusinessPartnerReq,
) BusinessPartnerReq {
	req := BusinessPartnerReq{
		General: input.General,
		APIType: "updates",
		Accepter: []string{
			"PersonMobilePhoneAuth",
		},
	}
	return req
}

func BusinessPartnerUpdatesRequestPersonGoogleAccountAuth(
	requestPram *types.Request,
	input types.BusinessPartnerSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")

	var request BusinessPartnerReq

	isMarkedForDeletion := false

	request = FuncBusinessPartnerUpdatesRequestPersonGoogleAccountAuth(
		requestPram,
		BusinessPartnerReq{
			General: General{
				BusinessPartner: input.General.BusinessPartner,
				PersonGoogleAccountAuth: []PersonGoogleAccountAuth{
					{
						GoogleAccountIsAuthenticated: input.General.BusinessPartnerPersonGoogleAccountAuth[0].GoogleAccountIsAuthenticated,
						LastChangeDate:               formattedDate,
						IsMarkedForDeletion:          &isMarkedForDeletion,
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

func BusinessPartnerUpdatesRequestPersonInstagramAuth(
	requestPram *types.Request,
	input types.BusinessPartnerSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")

	var request BusinessPartnerReq

	isMarkedForDeletion := false

	request = FuncBusinessPartnerUpdatesRequestPersonInstagramAuth(
		requestPram,
		BusinessPartnerReq{
			General: General{
				BusinessPartner: input.General.BusinessPartner,
				PersonInstagramAuth: []PersonInstagramAuth{
					{
						InstagramID:                input.General.BusinessPartnerPersonInstagramAuth[0].InstagramID,
						InstagramUserName:          input.General.BusinessPartnerPersonInstagramAuth[0].InstagramUserName,
						InstagramHasProfilePicture: input.General.BusinessPartnerPersonInstagramAuth[0].InstagramHasProfilePicture,
						InstagramProfilePictureURI: input.General.BusinessPartnerPersonInstagramAuth[0].InstagramProfilePictureURI,
						InstagramIsPrivate:         input.General.BusinessPartnerPersonInstagramAuth[0].InstagramIsPrivate,
						InstagramIsPublished:       input.General.BusinessPartnerPersonInstagramAuth[0].InstagramIsPublished,
						InstagramIsAuthenticated:   input.General.BusinessPartnerPersonInstagramAuth[0].InstagramIsAuthenticated,
						LastChangeDate:             formattedDate,
						IsMarkedForDeletion:        &isMarkedForDeletion,
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
