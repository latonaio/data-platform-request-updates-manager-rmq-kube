package apiModuleRuntimesRequestsPost

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type PostReq struct {
	Header   Header   `json:"Post"`
	Headers  []Header `json:"Posts"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type Header struct {
	Post                         *int             `json:"Post"`
	PostType                     string           `json:"PostType"`
	PostOwner                    int              `json:"PostOwner"`
	PostOwnerBusinessPartnerRole string           `json:"PostOwnerBusinessPartnerRole"`
	Description                  *string          `json:"Description"`
	LongText                     string           `json:"LongText"`
	Site                         *int             `json:"Site"`
	Tag1                         *string          `json:"Tag1"`
	Tag2                         *string          `json:"Tag2"`
	Tag3                         *string          `json:"Tag3"`
	Tag4                         *string          `json:"Tag4"`
	IsPublished                  bool             `json:"IsPublished"`
	CreationDate                 string           `json:"CreationDate"`
	CreationTime                 string           `json:"CreationTime"`
	LastChangeDate               string           `json:"LastChangeDate"`
	LastChangeTime               string           `json:"LastChangeTime"`
	IsMarkedForDeletion          *bool            `json:"IsMarkedForDeletion"`
	InstagramMedia               []InstagramMedia `json:"InstagramMedia"`
	Friend                       []Friend         `json:"Friend"`
}

type InstagramMedia struct {
	Post                            int     `json:"Post"`
	InstagramMediaID                string  `json:"InstagramMediaID"`
	InstagramMediaType              string  `json:"InstagramMediaType"`
	InstagramMediaCaption           *string `json:"InstagramMediaCaption"`
	InstagramMediaPermaLink         string  `json:"InstagramMediaPermaLink"`
	InstagramMediaURL               string  `json:"InstagramMediaURL"`
	InstagramMediaVideoThumbnailURL *string `json:"InstagramMediaVideoThumbnailURL"`
	InstagramMediaTimeStamp         string  `json:"InstagramMediaTimeStamp"`
	InstagramMediaDate              string  `json:"InstagramMediaDate"`
	InstagramMediaTime              string  `json:"InstagramMediaTime"`
	InstagramUserName               string  `json:"InstagramUserName"`
	CreationDate                    string  `json:"CreationDate"`
	CreationTime                    string  `json:"CreationTime"`
	LastChangeDate                  string  `json:"LastChangeDate"`
	LastChangeTime                  string  `json:"LastChangeTime"`
	IsMarkedForDeletion             *bool   `json:"IsMarkedForDeletion"`
}

type Friend struct {
	Post                int    `json:"Post"`
	Friend              int    `json:"Friend"`
	CreationDate        string `json:"CreationDate"`
	CreationTime        string `json:"CreationTime"`
	LastChangeDate      string `json:"LastChangeDate"`
	LastChangeTime      string `json:"LastChangeTime"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}

func FuncPostCreatesRequestHeader(
	requestPram *types.Request,
	input PostReq,
) PostReq {
	req := PostReq{
		Header:  input.Header,
		APIType: "creates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func PostCreatesRequestHeader(
	requestPram *types.Request,
	input types.PostSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POST_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request PostReq
	var header Header

	isPublished := false
	isMarkedForDeletion := false

	site := 1

	header = Header{
		Post:                         nil,
		PostType:                     "100",
		PostOwner:                    input.Header.PostOwner,
		PostOwnerBusinessPartnerRole: input.Header.PostOwnerBusinessPartnerRole,
		Description:                  input.Header.Description,
		LongText:                     input.Header.LongText,
		Site:                         &site,
		Tag1:                         input.Header.Tag1,
		Tag2:                         input.Header.Tag2,
		Tag3:                         input.Header.Tag3,
		Tag4:                         input.Header.Tag4,
		IsPublished:                  isPublished,
		CreationDate:                 formattedDate,
		CreationTime:                 formattedTime,
		LastChangeDate:               formattedDate,
		LastChangeTime:               formattedTime,
		IsMarkedForDeletion:          &isMarkedForDeletion,
	}

	request = FuncPostCreatesRequestHeader(
		requestPram,
		PostReq{
			Header: header,
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

func FuncPostCreatesRequestHeadersWithInstagramMedia(
	requestPram *types.Request,
	input PostReq,
) PostReq {
	req := PostReq{
		Headers: input.Headers,
		APIType: "creates",
		Accepter: []string{
			"Headers",
			"InstagramMedia",
		},
	}
	return req
}

func PostCreatesRequestHeadersWithInstagramMedia(
	requestPram *types.Request,
	input types.PostSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POST_SRV"
	aPIType := "creates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request PostReq
	var headers []Header

	isPublished := false
	isMarkedForDeletion := false

	site := 1

	for _, v := range input.Headers {
		headers = append(headers, Header{
			Post:                         nil,
			PostType:                     "100",
			PostOwner:                    v.PostOwner,
			PostOwnerBusinessPartnerRole: v.PostOwnerBusinessPartnerRole,
			Description:                  v.Description,
			LongText:                     v.LongText,
			Site:                         &site,
			Tag1:                         input.Header.Tag1,
			Tag2:                         input.Header.Tag2,
			Tag3:                         input.Header.Tag3,
			Tag4:                         input.Header.Tag4,
			IsPublished:                  isPublished,
			CreationDate:                 formattedDate,
			CreationTime:                 formattedTime,
			LastChangeDate:               formattedDate,
			LastChangeTime:               formattedTime,
			IsMarkedForDeletion:          &isMarkedForDeletion,
			InstagramMedia: []InstagramMedia{
				{
					Post:                            0,
					InstagramMediaID:                v.PostInstagramMedia[0].InstagramMediaID,
					InstagramMediaType:              v.PostInstagramMedia[0].InstagramMediaType,
					InstagramMediaCaption:           v.PostInstagramMedia[0].InstagramMediaCaption,
					InstagramMediaPermaLink:         v.PostInstagramMedia[0].InstagramMediaPermaLink,
					InstagramMediaURL:               v.PostInstagramMedia[0].InstagramMediaURL,
					InstagramMediaVideoThumbnailURL: v.PostInstagramMedia[0].InstagramMediaVideoThumbnailURL,
					InstagramMediaTimeStamp:         v.PostInstagramMedia[0].InstagramMediaTimeStamp,
					InstagramMediaDate:              v.PostInstagramMedia[0].InstagramMediaDate,
					InstagramMediaTime:              v.PostInstagramMedia[0].InstagramMediaTime,
					InstagramUserName:               v.PostInstagramMedia[0].InstagramMediaUserName,
					CreationDate:                    formattedDate,
					CreationTime:                    formattedTime,
					LastChangeDate:                  formattedDate,
					LastChangeTime:                  formattedTime,
					IsMarkedForDeletion:             &isMarkedForDeletion,
				},
			},
		})
	}

	request = FuncPostCreatesRequestHeadersWithInstagramMedia(
		requestPram,
		PostReq{
			Headers: headers,
		},
	)

	//request = FuncPostCreatesRequestHeadersWithInstagramMedia(
	//	requestPram,
	//	PostReq{
	//		Headers: Header{
	//			Post:                         nil,
	//			PostType:                     "100",
	//			PostOwner:                    input.Header.PostOwner,
	//			PostOwnerBusinessPartnerRole: input.Header.PostOwnerBusinessPartnerRole,
	//			Description:                  input.Header.Description,
	//			LongText:                     input.Header.LongText,
	//			Site:                         input.Header.Site,
	//			Tag1:                         input.Header.Tag1,
	//			Tag2:                         input.Header.Tag2,
	//			Tag3:                         input.Header.Tag3,
	//			Tag4:                         input.Header.Tag4,
	//			IsPublished:                  isPublished,
	//			CreationDate:                 formattedDate,
	//			CreationTime:                 formattedTime,
	//			LastChangeDate:               formattedDate,
	//			LastChangeTime:               formattedTime,
	//			IsMarkedForDeletion:          &isMarkedForDeletion,
	//			InstagramMedia: []InstagramMedia{
	//				{
	//					InstagramMediaID:                input.Header.PostInstagramMedia[0].InstagramMediaID,
	//					InstagramMediaType:              input.Header.PostInstagramMedia[0].InstagramMediaType,
	//					InstagramMediaCaption:           input.Header.PostInstagramMedia[0].InstagramMediaCaption,
	//					InstagramMediaPermaLink:         input.Header.PostInstagramMedia[0].InstagramMediaPermaLink,
	//					InstagramMediaURL:               input.Header.PostInstagramMedia[0].InstagramMediaURL,
	//					InstagramMediaVideoThumbnailURL: input.Header.PostInstagramMedia[0].InstagramMediaVideoThumbnailURL,
	//					InstagramMediaTimeStamp:         input.Header.PostInstagramMedia[0].InstagramMediaTimeStamp,
	//					InstagramMediaDate:              input.Header.PostInstagramMedia[0].InstagramMediaDate,
	//					InstagramMediaTime:              input.Header.PostInstagramMedia[0].InstagramMediaTime,
	//					InstagramUserName:               input.Header.PostInstagramMedia[0].InstagramUserName,
	//					CreationDate:                    formattedDate,
	//					CreationTime:                    formattedTime,
	//					LastChangeDate:                  formattedDate,
	//					LastChangeTime:                  formattedTime,
	//					IsMarkedForDeletion:             &isMarkedForDeletion,
	//				},
	//			},
	//		},
	//	},
	//)

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

func FuncPostUpdatesRequestHeader(
	requestPram *types.Request,
	input PostReq,
) PostReq {
	req := PostReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Header",
		},
	}
	return req
}

func PostUpdatesRequestHeader(
	requestPram *types.Request,
	input types.PostSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POST_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request PostReq

	request = FuncPostUpdatesRequestHeader(
		requestPram,
		PostReq{
			Header: Header{
				Post:           input.Header.Post,
				Description:    input.Header.Description,
				LongText:       input.Header.LongText,
				Site:           input.Header.Site,
				Tag1:           input.Header.Tag1,
				Tag2:           input.Header.Tag2,
				Tag3:           input.Header.Tag3,
				Tag4:           input.Header.Tag4,
				IsPublished:    input.Header.IsPublished,
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

func FuncPostUpdatesRequestHeadersWithInstagramMedia(
	requestPram *types.Request,
	input PostReq,
) PostReq {
	req := PostReq{
		Header:  input.Header,
		APIType: "updates",
		Accepter: []string{
			"Header",
			"InstagramMedia",
		},
	}
	return req
}

func PostUpdatesRequestHeadersWithInstagramMedia(
	requestPram *types.Request,
	input types.PostSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POST_SRV"
	aPIType := "updates"

	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04:05")

	var request PostReq

	request = FuncPostUpdatesRequestHeadersWithInstagramMedia(
		requestPram,
		PostReq{
			Header: Header{
				Post:           nil,
				Description:    input.Header.Description,
				LongText:       input.Header.LongText,
				Site:           input.Header.Site,
				Tag1:           input.Header.Tag1,
				Tag2:           input.Header.Tag2,
				Tag3:           input.Header.Tag3,
				Tag4:           input.Header.Tag4,
				IsPublished:    input.Header.IsPublished,
				LastChangeDate: formattedDate,
				LastChangeTime: formattedTime,
				InstagramMedia: []InstagramMedia{
					{
						InstagramMediaID:                input.Header.PostInstagramMedia[0].InstagramMediaID,
						InstagramMediaCaption:           input.Header.PostInstagramMedia[0].InstagramMediaCaption,
						InstagramMediaPermaLink:         input.Header.PostInstagramMedia[0].InstagramMediaPermaLink,
						InstagramMediaURL:               input.Header.PostInstagramMedia[0].InstagramMediaURL,
						InstagramMediaVideoThumbnailURL: input.Header.PostInstagramMedia[0].InstagramMediaVideoThumbnailURL,
						InstagramMediaTimeStamp:         input.Header.PostInstagramMedia[0].InstagramMediaTimeStamp,
						InstagramMediaDate:              input.Header.PostInstagramMedia[0].InstagramMediaDate,
						InstagramMediaTime:              input.Header.PostInstagramMedia[0].InstagramMediaTime,
						InstagramUserName:               input.Header.PostInstagramMedia[0].InstagramMediaUserName,
						LastChangeDate:                  formattedDate,
						LastChangeTime:                  formattedTime,
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
