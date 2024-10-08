package apiModuleRuntimesRequestsArticle

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type ArticleReq struct {
	Header   Header   `json:"Article"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Article                         *int                    `json:"Article"`
	ArticleType                     string                  `json:"ArticleType"`
	ArticleOwner                    int                     `json:"ArticleOwner"`
	ArticleOwnerBusinessPartnerRole string                  `json:"ArticleOwnerBusinessPartnerRole"`
	PersonResponsible               string                  `json:"PersonResponsible"`
	ValidityStartDate               string                  `json:"ValidityStartDate"`
	ValidityStartTime               string                  `json:"ValidityStartTime"`
	ValidityEndDate                 string                  `json:"ValidityEndDate"`
	ValidityEndTime                 string                  `json:"ValidityEndTime"`
	Description                     string                  `json:"Description"`
	LongText                        string                  `json:"LongText"`
	Introduction                    *string                 `json:"Introduction"`
	Site                            int                     `json:"Site"`
	Shop                            *int                    `json:"Shop"`
	Project                         *int                    `json:"Project"`
	WBSElement                      *int                    `json:"WBSElement"`
	Tag1                            *string                 `json:"Tag1"`
	Tag2                            *string                 `json:"Tag2"`
	Tag3                            *string                 `json:"Tag3"`
	Tag4                            *string                 `json:"Tag4"`
	DistributionProfile             string                  `json:"DistributionProfile"`
	QuestionnaireType               *string                 `json:"QuestionnaireType"`
	QuestionnaireTemplate           *string                 `json:"QuestionnaireTemplate"`
	CreationDate                    string                  `json:"CreationDate"`
	CreationTime                    string                  `json:"CreationTime"`
	LastChangeDate                  string                  `json:"LastChangeDate"`
	LastChangeTime                  string                  `json:"LastChangeTime"`
	CreateUser                      int                     `json:"CreateUser"`
	LastChangeUser                  int                     `json:"LastChangeUser"`
	IsReleased                      *bool                   `json:"IsReleased"`
	IsMarkedForDeletion             *bool                   `json:"IsMarkedForDeletion"`
	Partner                         []Partner               `json:"Partner"`
	Address                         []Address               `json:"Address"`
	Counter						  	[]Counter				`json:"Counter"`
	Like						  	[]Like				  	`json:"Like"`

}

type Partner struct {
	Article                 int     `json:"Article"`
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
	Article        int      `json:"Article"`
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
	Site           *int     `json:"Site"`
}

type Counter struct {
	Article					int		`json:"Article"`
	NumberOfLikes			int		`json:"NumberOfLikes"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	LastChangeDate			string	`json:"LastChangeDate"`
	LastChangeTime			string	`json:"LastChangeTime"`
}

type Like struct {
	Article					int		`json:"Article"`
	BusinessPartner			int		`json:"BusinessPartner"`
	Like					*bool	`json:"Like"`
	CreationDate			string	`json:"CreationDate"`
	CreationTime			string	`json:"CreationTime"`
	LastChangeDate			string	`json:"LastChangeDate"`
	LastChangeTime			string	`json:"LastChangeTime"`
}

func FuncArticleCreatesRequestAll(
	requestPram *types.Request,
	input ArticleReq,
) ArticleReq {
	req := ArticleReq{
		Header:  input.Header,
		APIType: "creates",
		Accepter: []string{
			"Header",
			"Address",
			"Counter",
			"Like",
		},
	}
	return req
}

func ArticleCreatesRequestAll(
	requestPram *types.Request,
	input types.ArticleSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ARTICLE_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request ArticleReq

	isReleased := false
	isMarkedForDeletion := false

	introduction := "Test Introduction"

	like := false

	shop := 1

	request = FuncArticleCreatesRequestAll(
		requestPram,
		ArticleReq{
			Header: Header{
				Article:                         nil,
				ArticleType:                     input.Header.ArticleType,
				ArticleOwner:                    input.Header.ArticleOwner,
				ArticleOwnerBusinessPartnerRole: input.Header.ArticleOwnerBusinessPartnerRole,
				PersonResponsible:               input.Header.PersonResponsible,
				ValidityStartDate:               input.Header.ValidityStartDate,
				ValidityStartTime:               input.Header.ValidityStartTime,
				ValidityEndDate:                 input.Header.ValidityEndDate,
				ValidityEndTime:                 input.Header.ValidityEndTime,
				Description:                     input.Header.Description,
				LongText:                        input.Header.LongText,
				Introduction:                    &introduction,
				Site:                            1,
				Shop:                            &shop,
				Tag1:                            input.Header.Tag1,
				Tag2:                            input.Header.Tag2,
				Tag3:                            input.Header.Tag3,
				Tag4:                            input.Header.Tag4,
				DistributionProfile:             input.Header.DistributionProfile,
				QuestionnaireType:               input.Header.QuestionnaireType,
				QuestionnaireTemplate:           input.Header.QuestionnaireTemplate,
				CreationDate:                    formattedDate,
				CreationTime:                    formattedTime,
				LastChangeDate:                  formattedDate,
				LastChangeTime:                  formattedTime,
				CreateUser:                      input.Header.CreateUser,
				LastChangeUser:                  input.Header.LastChangeUser,
				IsReleased:                      &isReleased,
				IsMarkedForDeletion:             &isMarkedForDeletion,
				Address: []Address{
					{
						AddressID:      input.Header.ArticleAddress[0].AddressID,
						PostalCode:     input.Header.ArticleAddress[0].PostalCode,
						LocalSubRegion: input.Header.ArticleAddress[0].LocalSubRegion,
						LocalRegion:    input.Header.ArticleAddress[0].LocalRegion,
						Country:        input.Header.ArticleAddress[0].Country,
						GlobalRegion:   input.Header.ArticleAddress[0].GlobalRegion,
						TimeZone:       input.Header.ArticleAddress[0].TimeZone,
						District:       input.Header.ArticleAddress[0].District,
						StreetName:     input.Header.ArticleAddress[0].StreetName,
						CityName:       input.Header.ArticleAddress[0].CityName,
						Building:       input.Header.ArticleAddress[0].Building,
						XCoordinate:    input.Header.ArticleAddress[0].XCoordinate,
						YCoordinate:    input.Header.ArticleAddress[0].YCoordinate,
						ZCoordinate:    input.Header.ArticleAddress[0].ZCoordinate,
					},
				},
				Counter: []Counter{
					{
						NumberOfLikes:				0,
						CreationDate:				formattedDate,
						CreationTime:				formattedTime,
						LastChangeDate:				formattedDate,
						LastChangeTime:				formattedTime,
					},
				},
				Like: []Like{
					{
						BusinessPartner:			input.Header.ArticleLike[0].BusinessPartner,
						Like:						&like,
						CreationDate:				formattedDate,
						CreationTime:				formattedTime,
						LastChangeDate:				formattedDate,
						LastChangeTime:				formattedTime,
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

func FuncArticleUpdatesRequestHeader(
	requestPram *types.Request,
	input ArticleReq,
) ArticleReq {
	req := ArticleReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func ArticleUpdatesRequestHeader(
	requestPram *types.Request,
	input types.ArticleSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ARTICLE_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request ArticleReq

	request = FuncArticleUpdatesRequestHeader(
		requestPram,
		ArticleReq{
			Header: Header{
				Article:                         input.Header.Article,
				ArticleType:                     input.Header.ArticleType,
				ArticleOwner:                    input.Header.ArticleOwner,
				ArticleOwnerBusinessPartnerRole: input.Header.ArticleOwnerBusinessPartnerRole,
				PersonResponsible:               input.Header.PersonResponsible,
				ValidityStartDate:               input.Header.ValidityStartDate,
				ValidityStartTime:               input.Header.ValidityStartTime,
				ValidityEndDate:                 input.Header.ValidityEndDate,
				ValidityEndTime:                 input.Header.ValidityEndTime,
				Description:                     input.Header.Description,
				LongText:                        input.Header.LongText,
				Introduction:                    input.Header.Introduction,
				Site:                            input.Header.Site,
				Shop:                            input.Header.Shop,
				Tag1:                            input.Header.Tag1,
				Tag2:                            input.Header.Tag2,
				Tag3:                            input.Header.Tag3,
				Tag4:                            input.Header.Tag4,
				DistributionProfile:             input.Header.DistributionProfile,
				QuestionnaireType:               input.Header.QuestionnaireType,
				QuestionnaireTemplate:           input.Header.QuestionnaireTemplate,
				LastChangeDate:                  formattedDate,
				LastChangeTime:                  formattedTime,
				LastChangeUser:                  input.Header.LastChangeUser,
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

func FuncArticleUpdatesRequestCounter(
	requestPram *types.Request,
	input ArticleReq,
) ArticleReq {
	req := ArticleReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Counter",
		},
	}
	return req
}

func ArticleUpdatesRequestCounter(
	requestPram *types.Request,
	input types.ArticleSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request ArticleReq

	request = FuncArticleUpdatesRequestCounter(
		requestPram,
		ArticleReq{
			Header: Header{
				Article:                         input.Header.Article,
				Counter: []Counter{
					{
						NumberOfLikes:				input.Header.ArticleCounter[0].NumberOfLikes,
						LastChangeDate:				formattedDate,
						LastChangeTime:				formattedTime,
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

func FuncArticleUpdatesRequestLike(
	requestPram *types.Request,
	input ArticleReq,
) ArticleReq {
	req := ArticleReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Like",
		},
	}
	return req
}

func ArticleUpdatesRequestLike(
	requestPram *types.Request,
	input types.ArticleSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request ArticleReq

	request = FuncArticleUpdatesRequestLike(
		requestPram,
		ArticleReq{
			Header: Header{
				Article:                         input.Header.Article,
				Like: []Like{
					{
						BusinessPartner:			input.Header.ArticleLike[0].BusinessPartner,
						Like:						input.Header.ArticleLike[0].Like,
						LastChangeDate:				formattedDate,
						LastChangeTime:				formattedTime,
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
