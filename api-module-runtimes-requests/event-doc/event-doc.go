package apiModuleRuntimesRequestsEventDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type EventDocReq struct {
	EventDoc    EventDoc `json:"Event"`
	EventQRCode EventDoc `json:"QRCode"`
	APIType     string   `json:"api_type"`
	Accepter    []string `json:"accepter"`
	DocData     string   `json:"doc_data"`
}

type EventDoc struct {
	Event     int       `json:"Event"`
	HeaderDoc HeaderDoc `json:"HeaderDoc"`
}

type HeaderDoc struct {
	Event                    int     `json:"Event"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}

func FuncEventDocCreatesRequestHeaderDoc(
	requestPram *types.Request,
	input EventDocReq,
) EventDocReq {
	req := EventDocReq{
		EventDoc:    input.EventDoc,
		EventQRCode: input.EventQRCode,
		APIType:     "creates",
		Accepter:    input.Accepter,
		DocData:     input.DocData,
	}
	return req
}

func EventDocCreatesRequestHeaderDoc(
	requestPram *types.Request,
	input types.EventDocSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_EVENT_DOC_SRV"
	aPIType := "creates"

	var request EventDocReq

	fileName := "EventImage"

	request = FuncEventDocCreatesRequestHeaderDoc(
		requestPram,
		EventDocReq{
			EventDoc: EventDoc{
				Event: input.Event.Event,
				HeaderDoc: HeaderDoc{
					Event:                    input.Event.HeaderDoc.Event,
					DocType:                  input.Event.HeaderDoc.DocType,
					DocVersionID:             1,
					DocID:                    input.Event.HeaderDoc.DocID,
					FileExtension:            input.Event.HeaderDoc.FileExtension,
					FileName:                 fileName,
					FilePath:                 input.Event.HeaderDoc.FilePath,
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
