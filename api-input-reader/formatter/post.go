package formatter

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	apiModuleRuntimesResponsesInstagramUserMedia "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-responses/instagram-user-media"
	apiModuleRuntimesResponsesPost "data-platform-request-updates-manager-rmq-kube/api-module-runtimes-responses/post"
	apiOutputFormatter "data-platform-request-updates-manager-rmq-kube/api-output-formatter"
	"encoding/json"
	"golang.org/x/xerrors"
	"time"
)

type PostSDC struct {
	ConnectionKey    string       `json:"connection_key"`
	Result           bool         `json:"result"`
	RedisKey         string       `json:"redis_key"`
	Filepath         string       `json:"filepath"`
	APIStatusCode    int          `json:"api_status_code"`
	RuntimeSessionID string       `json:"runtime_session_id"`
	BusinessPartner  *int         `json:"business_partner"`
	ServiceLabel     string       `json:"service_label"`
	APIType          string       `json:"api_type"`
	PostHeaders      []PostHeader `json:"Posts"`
	APISchema        string       `json:"api_schema"`
	Accepter         []string     `json:"accepter"`
	Deleted          bool         `json:"deleted"`
}

type PostHeader struct {
	Post                         *int                 `json:"Post"`
	PostType                     string               `json:"PostType"`
	PostOwner                    int                  `json:"PostOwner"`
	PostOwnerBusinessPartnerRole string               `json:"PostOwnerBusinessPartnerRole"`
	Description                  *string              `json:"Description"`
	LongText                     string               `json:"LongText"`
	Site                         *int                 `json:"Site"`
	Tag1                         *string              `json:"Tag1"`
	Tag2                         *string              `json:"Tag2"`
	Tag3                         *string              `json:"Tag3"`
	Tag4                         *string              `json:"Tag4"`
	IsPublished                  bool                 `json:"IsPublished"`
	CreationDate                 string               `json:"CreationDate"`
	CreationTime                 string               `json:"CreationTime"`
	LastChangeDate               string               `json:"LastChangeDate"`
	LastChangeTime               string               `json:"LastChangeTime"`
	IsMarkedForDeletion          *bool                `json:"IsMarkedForDeletion"`
	PostHeaderDoc                []PostHeaderDoc      `json:"HeaderDoc"`
	PostInstagramMedia           []PostInstagramMedia `json:"InstagramMedia"`
}

type PostHeaderDoc struct {
	Post                     int     `json:"Post"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}

type PostInstagramMedia struct {
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

type PostDocSDC struct {
	ConnectionKey     string    `json:"connection_key"`
	Result            bool      `json:"result"`
	RedisKey          string    `json:"redis_key"`
	Filepath          string    `json:"filepath"`
	APIStatusCode     int       `json:"api_status_code"`
	RuntimeSessionID  string    `json:"runtime_session_id"`
	BusinessPartnerID *int      `json:"business_partner"`
	ServiceLabel      string    `json:"service_label"`
	Posts             []PostDoc `json:"Posts"`
	APISchema         string    `json:"api_schema"`
	Accepter          []string  `json:"accepter"`
	Deleted           bool      `json:"deleted"`
}

type PostDoc struct {
	Post      int           `json:"Post"`
	HeaderDoc PostHeaderDoc `json:"HeaderDoc"`
}

func ConvertToPostCreatesHeadersWithInstagramMedia(
	requestPram *types.Request,
	input *apiModuleRuntimesResponsesInstagramUserMedia.InstagramUserMediaRes,
) (*types.PostSDC, error) {
	var postHeaders []types.PostHeader

	for _, v := range input.Message.InstagramUserMedia {
		var instagramMediaCaption string
		if v.InstagramMediaCaption == nil {
			instagramMediaCaption = ""
		} else {
			instagramMediaCaption = *v.InstagramMediaCaption
		}

		instagramMediaTimeStamp := v.InstagramMediaTimeStamp

		utcTime, err := time.Parse("2006-01-02T15:04:05-0700", instagramMediaTimeStamp)
		if err != nil {
			return nil, xerrors.Errorf("Error parsing time: %w", err)
		}

		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			return nil, xerrors.Errorf("Error loading JST timezone: %w", err)
		}

		jstTime := utcTime.In(jst)

		formattedInstagramMediaTimeStamp := jstTime.Format("2006-01-02 15:04:05")

		instagramMediaDate := jstTime.Format("2006-01-02")
		instagramMediaTime := jstTime.Format("15:04:05")

		postHeaders = append(postHeaders, types.PostHeader{
			Post:                         nil,
			PostOwner:                    *requestPram.BusinessPartner,
			PostOwnerBusinessPartnerRole: *requestPram.BusinessPartnerRole,
			LongText:                     instagramMediaCaption,
			PostInstagramMedia: []types.PostInstagramMedia{
				{
					InstagramMediaID:                v.InstagramID,
					InstagramMediaType:              v.InstagramMediaType,
					InstagramMediaCaption:           v.InstagramMediaCaption,
					InstagramMediaPermaLink:         v.InstagramMediaPermaLink,
					InstagramMediaURL:               v.InstagramMediaURL,
					InstagramMediaVideoThumbnailURL: v.InstagramMediaVideoThumbnailURL,
					InstagramMediaTimeStamp:         formattedInstagramMediaTimeStamp,
					InstagramMediaDate:              instagramMediaDate,
					InstagramMediaTime:              instagramMediaTime,
					InstagramMediaUserName:          v.InstagramMediaUserName,
				},
			},
		})
	}

	var postSdc types.PostSDC

	postSdc.Headers = postHeaders

	return &postSdc, nil
}

func ConvertToPostDocCreatesHeaderDoc(
	requestPram *types.Request,
	//input *apiModuleRuntimesResponsesPost.PostRes,
	postResponse *apiOutputFormatter.PostSDC,
) (*types.PostDocSDC, error) {
	var input *apiModuleRuntimesResponsesPost.PostRes

	messageJSON, err := json.Marshal(postResponse)
	if err != nil {
		return nil, xerrors.Errorf("Error marshaling postResponse.Message: %w", err)
	}

	err = json.Unmarshal(messageJSON, &input)
	if err != nil {
		return nil, xerrors.Errorf("cache data unmarshal error: %w", err)
	}

	var postDoc []types.PostDoc

	if input.Message.InstagramMedia == nil {
		return nil, xerrors.Errorf("No Post Response headers in input")
	}

	for _, v := range *input.Message.InstagramMedia {
		instagramMediaType := v.InstagramMediaType
		if v.InstagramMediaType == "VIDEO" {
			instagramMediaType = "MP4"
		}

		fileExtension := func(mediaType string) string {
			switch mediaType {
			case "IMAGE":
				return "png"
			case "VIDEO":
				return "mp4"
			default:
				return ""
			}
		}(v.InstagramMediaType)

		postDoc = append(postDoc, types.PostDoc{
			Post: v.Post,
			HeaderDoc: types.PostHeaderDoc{
				Post:                     v.Post,
				DocType:                  instagramMediaType,
				DocVersionID:             1,
				DocID:                    &v.InstagramMediaID,
				FileExtension:            fileExtension,
				FileName:                 v.InstagramMediaID,
				FilePath:                 " ",
				DocIssuerBusinessPartner: 1001,
			},
		})
	}

	var postDocSdc types.PostDocSDC

	postDocSdc.Posts = postDoc

	return &postDocSdc, nil
}
