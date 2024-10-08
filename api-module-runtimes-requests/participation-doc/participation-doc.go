package apiModuleRuntimesRequestsParticipationDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ParticipationDocReq struct {
	ParticipationDoc    ParticipationDoc `json:"Participation"`
	ParticipationQRCode ParticipationDoc `json:"QRCode"`
	APIType             string           `json:"api_type"`
	Accepter            []string         `json:"accepter"`
	DocData             string           `json:"doc_data"`
}

type ParticipationDoc struct {
	Participation     int       `json:"Participation"`
	HeaderDoc         HeaderDoc `json:"HeaderDoc"`
}

type HeaderDoc struct {
	Participation            int     `json:"Participation"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}

func FuncParticipationDocCreatesRequestHeaderDoc(
	requestPram *types.Request,
	input ParticipationDocReq,
) ParticipationDocReq {
	req := ParticipationDocReq{
		ParticipationDoc:    input.ParticipationDoc,
		ParticipationQRCode: input.ParticipationQRCode,
		APIType:             "creates",
		Accepter:            input.Accepter,
		DocData:             input.DocData,
	}
	return req
}

func ParticipationDocCreatesRequestHeaderDoc(
	requestPram *types.Request,
	input types.ParticipationDocSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_PARTICIPATION_DOC_SRV"
	aPIType := "creates"

	var request ParticipationDocReq

	fileName := "ParticipationImage"

	request = FuncParticipationDocCreatesRequestHeaderDoc(
		requestPram,
		ParticipationDocReq{
			ParticipationDoc: ParticipationDoc{
				Participation: input.Participation.Participation,
				HeaderDoc: HeaderDoc{
					Participation:            input.Participation.HeaderDoc.Participation,
					DocType:                  input.Participation.HeaderDoc.DocType,
					DocVersionID:             1,
					DocID:                    input.Participation.HeaderDoc.DocID,
					FileExtension:            input.Participation.HeaderDoc.FileExtension,
					FileName:                 fileName,
					FilePath:                 input.Participation.HeaderDoc.FilePath,
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
