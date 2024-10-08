package apiModuleRuntimesRequestsShopDoc

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type ShopDocReq struct {
	ShopDoc  ShopDoc  `json:"Shop"`
	APIType  string   `json:"api_type"`
	Accepter []string `json:"accepter"`
	DocData  string   `json:"doc_data"`
}

type ShopDoc struct {
	Shop      int       `json:"Shop"`
	HeaderDoc HeaderDoc `json:"HeaderDoc"`
}

type HeaderDoc struct {
	Shop                     int     `json:"Shop"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}

func FuncShopDocCreatesRequestHeaderDoc(
	requestPram *types.Request,
	input ShopDocReq,
) ShopDocReq {
	req := ShopDocReq{
		ShopDoc:  input.ShopDoc,
		APIType:  "creates",
		Accepter: []string{},
		DocData:  input.DocData,
	}
	return req
}

func ShopDocCreatesRequestHeaderDoc(
	requestPram *types.Request,
	input types.ShopDocSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_SHOP_DOC_SRV"
	aPIType := "creates"

	var request ShopDocReq

	fileName := "ShopImage"

	request = FuncShopDocCreatesRequestHeaderDoc(
		requestPram,
		ShopDocReq{
			ShopDoc: ShopDoc{
				Shop: input.Shop.Shop,
				HeaderDoc: HeaderDoc{
					Shop:                     input.Shop.HeaderDoc.Shop,
					DocType:                  input.Shop.HeaderDoc.DocType,
					DocVersionID:             1,
					DocID:                    input.Shop.HeaderDoc.DocID,
					FileExtension:            input.Shop.HeaderDoc.FileExtension,
					FileName:                 fileName,
					FilePath:                 input.Shop.HeaderDoc.FilePath,
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
