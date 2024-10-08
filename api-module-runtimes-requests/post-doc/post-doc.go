package apiModuleRuntimesRequestsPostDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type PostDocReq struct {
	PostDoc  PostDoc   `json:"Post"`
	PostDocs []PostDoc `json:"Posts"`
	APIType  string    `json:"api_type"`
	Accepter []string  `json:"accepter"`
	DocData  string    `json:"doc_data"`
}

type PostDoc struct {
	Post      int       `json:"Post"`
	HeaderDoc HeaderDoc `json:"HeaderDoc"`
}

type HeaderDoc struct {
	Post                     int     `json:"Post"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}

func FuncPostDocCreatesRequestHeaderDoc(
	requestPram *types.Request,
	input PostDocReq,
) PostDocReq {
	req := PostDocReq{
		PostDoc:  input.PostDoc,
		APIType:  "creates",
		Accepter: []string{},
		DocData:  input.DocData,
	}
	return req
}

func FuncPostDocCreatesRequestHeaderDocs(
	requestPram *types.Request,
	input PostDocReq,
) PostDocReq {
	req := PostDocReq{
		PostDocs: input.PostDocs,
		APIType:  "creates",
		Accepter: []string{},
	}
	return req
}

func PostDocCreatesRequestHeaderDoc(
	requestPram *types.Request,
	input types.PostDocSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POST_DOC_SRV"
	aPIType := "creates"

	var request PostDocReq

	request = FuncPostDocCreatesRequestHeaderDoc(
		requestPram,
		PostDocReq{
			PostDoc: PostDoc{
				Post: input.Post.Post,
				HeaderDoc: HeaderDoc{
					Post:                     input.Post.HeaderDoc.Post,
					DocType:                  input.Post.HeaderDoc.DocType,
					DocVersionID:             1,
					DocID:                    input.Post.HeaderDoc.DocID,
					FileExtension:            input.Post.HeaderDoc.FileExtension,
					FileName:                 input.Post.HeaderDoc.FileName,
					FilePath:                 "",
					DocIssuerBusinessPartner: 1001,
				},
			},
			DocData: input.DocData,
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

func PostDocCreatesRequestHeaderDocs(
	requestPram *types.Request,
	input types.PostDocSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_POST_DOC_SRV"
	aPIType := "creates"

	var request PostDocReq

	var posts []PostDoc

	for _, v := range input.Posts {
		posts = append(posts, PostDoc{
			Post: v.Post,
			HeaderDoc: HeaderDoc{
				Post:                     v.HeaderDoc.Post,
				DocType:                  v.HeaderDoc.DocType,
				DocVersionID:             v.HeaderDoc.DocVersionID,
				DocID:                    v.HeaderDoc.DocID,
				FileExtension:            v.HeaderDoc.FileExtension,
				FileName:                 v.HeaderDoc.FileName,
				FilePath:                 v.HeaderDoc.FilePath,
				DocIssuerBusinessPartner: v.HeaderDoc.DocIssuerBusinessPartner,
			},
		})
	}

	request = FuncPostDocCreatesRequestHeaderDocs(
		requestPram,
		PostDocReq{
			PostDocs: posts,
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
