package apiModuleRuntimesRequestsInstagramUserMedia

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
	"time"
)

type InstagramUserMediaReq struct {
	InstagramUserMedia InstagramUserMedia `json:"InstagramUserMedia"`
	APIType            string             `json:"api_type"`
	Accepter           []string           `json:"accepter"`
}

type InstagramUserMedia struct {
	InstagramID string  `json:"InstagramID"`
	AccessToken string  `json:"AccessToken"`
	Fields      string  `json:"Fields"`
	Since       *string `json:"Since"`
	Until       *string `json:"Until"`
}

func FuncInstagramUserMediaRequestsMedias(
	requestPram *types.Request,
	input InstagramUserMediaReq,
) InstagramUserMediaReq {
	fields := "permalink,media_url,thumbnail_url,caption,media_type,timestamp,id,username"

	req := InstagramUserMediaReq{
		InstagramUserMedia: InstagramUserMedia{
			InstagramID: input.InstagramUserMedia.InstagramID,
			AccessToken: input.InstagramUserMedia.AccessToken,
			Fields:      fields,
			Since:       input.InstagramUserMedia.Since,
			Until:       input.InstagramUserMedia.Until,
		},
		APIType: "requests",
		Accepter: []string{
			"InstagramUserMedias",
		},
	}
	return req
}

func InstagramUserMediaRequestsMedias(
	requestPram *types.Request,
	input types.InstagramUserMediaSDC,
	controller *beego.Controller,
) []byte {
	aPIServiceName := "DPFM_API_INSTAGRAM_USER_MEDIA_SRV"
	aPIType := "requests"

	var request InstagramUserMediaReq
	fields := "permalink,media_url,thumbnail_url,caption,media_type,timestamp,id,username"

	now := time.Now()

	since := now.AddDate(0, 0, -90).Format("2006-01-02 15:04:05")
	until := now.Format("2006-01-02 15:04:05")

	request = FuncInstagramUserMediaRequestsMedias(
		requestPram,
		InstagramUserMediaReq{
			InstagramUserMedia: InstagramUserMedia{
				InstagramID: input.InstagramUserMedia.InstagramID,
				AccessToken: input.InstagramUserMedia.AccessToken,
				Fields:      fields,
				Since:       &since,
				Until:       &until,
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
