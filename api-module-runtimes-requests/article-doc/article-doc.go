package apiModuleRuntimesRequestsArticleDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ArticleDocReq struct {
	ArticleDoc    ArticleDoc `json:"Article"`
	ArticleQRCode ArticleDoc `json:"QRCode"`
	APIType     string   `json:"api_type"`
	Accepter    []string `json:"accepter"`
	DocData     string   `json:"doc_data"`
}

type ArticleDoc struct {
	Article     int       `json:"Article"`
	HeaderDoc   HeaderDoc `json:"HeaderDoc"`
}

type HeaderDoc struct {
	Article                  int     `json:"Article"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}

func FuncArticleDocCreatesRequestHeaderDoc(
	requestPram *types.Request,
	input ArticleDocReq,
) ArticleDocReq {
	req := ArticleDocReq{
		ArticleDoc:    input.ArticleDoc,
		ArticleQRCode: input.ArticleQRCode,
		APIType:       "creates",
		Accepter:      input.Accepter,
		DocData:       input.DocData,
	}
	return req
}

func ArticleDocCreatesRequestHeaderDoc(
	requestPram *types.Request,
	input types.ArticleDocSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_ARTICLE_DOC_SRV"
	aPIType := "creates"

	var request ArticleDocReq

	fileName := "ArticleImage"

	request = FuncArticleDocCreatesRequestHeaderDoc(
		requestPram,
		ArticleDocReq{
			ArticleDoc: ArticleDoc{
				Article: input.Article.Article,
				HeaderDoc: HeaderDoc{
					Article:                  input.Article.HeaderDoc.Article,
					DocType:                  input.Article.HeaderDoc.DocType,
					DocVersionID:             1,
					DocID:                    input.Article.HeaderDoc.DocID,
					FileExtension:            input.Article.HeaderDoc.FileExtension,
					FileName:                 fileName,
					FilePath:                 input.Article.HeaderDoc.FilePath,
					DocIssuerBusinessPartner: 1001,
				},
			},
			DocData:  input.DocData,
			Accepter: input.Accepter,
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
