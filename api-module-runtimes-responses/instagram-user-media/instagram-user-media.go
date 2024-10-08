package apiModuleRuntimesResponsesInstagramUserMedia

type InstagramUserMediaRes struct {
	Message Message `json:"message,omitempty"`
}

type Message struct {
	InstagramUserMedia []InstagramUserMedia `json:"InstagramUserMediaResponse"`
}

type InstagramUserMedia struct {
	InstagramID                     string  `json:"InstagramID"`
	InstagramMediaType              string  `json:"InstagramMediaType"`
	InstagramMediaCaption           *string `json:"InstagramMediaCaption"`
	InstagramMediaPermaLink         string  `json:"InstagramMediaPermaLink"`
	InstagramMediaURL               string  `json:"InstagramMediaURL"`
	InstagramMediaVideoThumbnailURL *string `json:"InstagramMediaVideoThumbnailURL"`
	InstagramMediaTimeStamp         string  `json:"InstagramMediaTimeStamp"`
	InstagramMediaUserName          string  `json:"InstagramMediaUserName"`
}
