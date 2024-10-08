package apiModuleRuntimesRequestsQRCode

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type QRCodeReq struct {
	QRCode   QRCode   `json:"QRCode"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
}

type QRCodeGlobal struct {
	QRCode QRCode `json:"QRCode"`
}

type QRCode struct {
	APIServiceName            string `json:"APIServiceName"`
	ServiceLabel              string `json:"ServiceLabel"`
	APIType                   string `json:"APIType"`
	URLFormatBeforeConversion string `json:"URLFormatBeforeConversion"`
	BusinessPartner           *int   `json:"BusinessPartner"`
	Event                     *int   `json:"Event"`
	Article                   *int   `json:"Article"`
	Shop                      *int   `json:"Shop"`
	Site                      *int   `json:"Site"`
	Station                   *int   `json:"Station"`
	BusStop                   *int   `json:"BusStop"`
	Participation             *int   `json:"Participation"`
}

func FuncQRCodeGeneratesRequestBusinessPartner(
	requestPram *types.Request,
	input QRCodeReq,
) QRCodeReq {
	req := QRCodeReq{
		QRCode:   input.QRCode,
		APIType:  "generates",
		Accepter: []string{"QRCodeBusinessPartner"},
	}
	return req
}

func QRCodeGeneratesRequestBusinessPartner(
	requestPram *types.Request,
	input types.QRCodeSDC,
	controller *beego.Controller,
) []byte {

	aPIServiceName := "DPFM_API_QRCODE_SRV"
	aPIType := "generates"

	var request QRCodeReq

	request = FuncQRCodeGeneratesRequestBusinessPartner(
		requestPram,
		QRCodeReq{
			QRCode: QRCode{
				APIServiceName:            input.QRCodeGlobal.APIServiceName,
				ServiceLabel:              input.QRCodeGlobal.ServiceLabel,
				APIType:                   input.QRCodeGlobal.APIType,
				URLFormatBeforeConversion: input.QRCodeGlobal.URLFormatBeforeConversion,
				BusinessPartner:		   input.QRCodeGlobal.BusinessPartner,
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

func FuncQRCodeGeneratesRequestEvent(
	requestPram *types.Request,
	input QRCodeReq,
) QRCodeReq {
	req := QRCodeReq{
		QRCode:   input.QRCode,
		APIType:  "generates",
		Accepter: []string{"QRCodeEvent"},
	}
	return req
}

func QRCodeGeneratesRequestEvent(
	requestPram *types.Request,
	input types.QRCodeSDC,
	controller *beego.Controller,
) []byte {

	aPIServiceName := "DPFM_API_QRCODE_SRV"
	aPIType := "generates"

	var request QRCodeReq

	request = FuncQRCodeGeneratesRequestEvent(
		requestPram,
		QRCodeReq{
			QRCode: QRCode{
				APIServiceName:            input.QRCodeGlobal.APIServiceName,
				ServiceLabel:              input.QRCodeGlobal.ServiceLabel,
				APIType:                   input.QRCodeGlobal.APIType,
				URLFormatBeforeConversion: input.QRCodeGlobal.URLFormatBeforeConversion,
//				BusinessPartner:           input.QRCodeGlobal.BusinessPartner,
				Event:                     input.QRCodeGlobal.Event,
//				Article:                   input.QRCodeGlobal.Article,
//				Shop:                      input.QRCodeGlobal.Shop,
//				Site:                      input.QRCodeGlobal.Site,
//				Station:                   input.QRCodeGlobal.Station,
//				BusStop:                   input.QRCodeGlobal.BusStop,
//				Participation:             input.QRCodeGlobal.Participation,
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

func FuncQRCodeGeneratesRequestSite(
	requestPram *types.Request,
	input QRCodeReq,
) QRCodeReq {
	req := QRCodeReq{
		QRCode:   input.QRCode,
		APIType:  "generates",
		Accepter: []string{"QRCodeSite"},
	}
	return req
}

func QRCodeGeneratesRequestSite(
	requestPram *types.Request,
	input types.QRCodeSDC,
	controller *beego.Controller,
) []byte {

	aPIServiceName := "DPFM_API_QRCODE_SRV"
	aPIType := "generates"

	var request QRCodeReq

	request = FuncQRCodeGeneratesRequestSite(
		requestPram,
		QRCodeReq{
			QRCode: QRCode{
				APIServiceName:            input.QRCodeGlobal.APIServiceName,
				ServiceLabel:              input.QRCodeGlobal.ServiceLabel,
				APIType:                   input.QRCodeGlobal.APIType,
				URLFormatBeforeConversion: input.QRCodeGlobal.URLFormatBeforeConversion,
				Site:					   input.QRCodeGlobal.Site,
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

func FuncQRCodeGeneratesRequestParticipation(
	requestPram *types.Request,
	input QRCodeReq,
) QRCodeReq {
	req := QRCodeReq{
		QRCode:   input.QRCode,
		APIType:  "generates",
		Accepter: []string{"QRCodeParticipation"},
	}
	return req
}

func QRCodeGeneratesRequestParticipation(
	requestPram *types.Request,
	input types.QRCodeSDC,
	controller *beego.Controller,
) []byte {

	aPIServiceName := "DPFM_API_QRCODE_SRV"
	aPIType := "generates"

	var request QRCodeReq

	request = FuncQRCodeGeneratesRequestParticipation(
		requestPram,
		QRCodeReq{
			QRCode: QRCode{
				APIServiceName:            input.QRCodeGlobal.APIServiceName,
				ServiceLabel:              input.QRCodeGlobal.ServiceLabel,
				APIType:                   input.QRCodeGlobal.APIType,
				URLFormatBeforeConversion: input.QRCodeGlobal.URLFormatBeforeConversion,
				Participation:			   input.QRCodeGlobal.Participation,
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
