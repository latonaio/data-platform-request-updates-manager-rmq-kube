package apiModuleRuntimesRequestsBusinessPartnerDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type BusinessPartnerDocReq struct {
	BusinessPartnerDoc  BusinessPartnerDoc `json:"BusinessPartner"`
	APIType             string   `json:"api_type"`
	Accepter            []string `json:"accepter"`
	DocData             string   `json:"doc_data"`
}

type BusinessPartnerDoc struct {
	BusinessPartner      int        `json:"BusinessPartner"`
	GeneralDoc           GeneralDoc `json:"GeneralDoc"`
}

type GeneralDoc struct {
	BusinessPartner          int     `json:"BusinessPartner"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}

func FuncBusinessPartnerDocCreatesRequestGeneralDoc(
	requestPram *types.Request,
	input BusinessPartnerDocReq,
) BusinessPartnerDocReq {
	req := BusinessPartnerDocReq{
		BusinessPartnerDoc:  input.BusinessPartnerDoc,
		APIType:  "creates",
		Accepter: []string{},
		DocData:  input.DocData,
	}
	return req
}

func BusinessPartnerDocCreatesRequestGeneralDoc(
	requestPram *types.Request,
	input types.BusinessPartnerDocSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_BUSINESS_PARTNER_DOC_SRV"
	aPIType := "creates"

	var request BusinessPartnerDocReq

	fileName := "BusinessPartnerImage"

	request = FuncBusinessPartnerDocCreatesRequestGeneralDoc(
		requestPram,
		BusinessPartnerDocReq{
			BusinessPartnerDoc: BusinessPartnerDoc{
				BusinessPartner: input.BusinessPartner.BusinessPartner,
				GeneralDoc: GeneralDoc{
					BusinessPartner:          input.BusinessPartner.GeneralDoc.BusinessPartner,
					DocType:                  input.BusinessPartner.GeneralDoc.DocType,
					DocVersionID:             1,
					DocID:                    input.BusinessPartner.GeneralDoc.DocID,
					FileExtension:            input.BusinessPartner.GeneralDoc.FileExtension,
					FileName:                 fileName,
					FilePath:                 input.BusinessPartner.GeneralDoc.FilePath,
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
