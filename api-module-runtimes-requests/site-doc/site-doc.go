package apiModuleRuntimesRequestsSiteDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type SiteDocReq struct {
	SiteDoc  SiteDoc  `json:"Site"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
	DocData  string   `json:"doc_data"`
}

type SiteDoc struct {
	Site      int       `json:"Site"`
	HeaderDoc HeaderDoc `json:"HeaderDoc"`
}

type HeaderDoc struct {
	Site                     int     `json:"Site"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}

func FuncSiteDocCreatesRequestHeaderDoc(
	requestPram *types.Request,
	input SiteDocReq,
) SiteDocReq {
	req := SiteDocReq{
		SiteDoc:  input.SiteDoc,
		APIType:  "creates",
		Accepter: []string{},
		DocData:  input.DocData,
	}
	return req
}

func SiteDocCreatesRequestHeaderDoc(
	requestPram *types.Request,
	input types.SiteDocSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SITE_DOC_SRV"
	aPIType := "creates"

	var request SiteDocReq

	fileName := "SiteImage"

	request = FuncSiteDocCreatesRequestHeaderDoc(
		requestPram,
		SiteDocReq{
			SiteDoc: SiteDoc{
				Site: input.Site.Site,
				HeaderDoc: HeaderDoc{
					Site:                     input.Site.HeaderDoc.Site,
					DocType:                  input.Site.HeaderDoc.DocType,
					DocVersionID:             1,
					DocID:                    input.Site.HeaderDoc.DocID,
					FileExtension:            input.Site.HeaderDoc.FileExtension,
					FileName:                 fileName,
					FilePath:                 input.Site.HeaderDoc.FilePath,
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
