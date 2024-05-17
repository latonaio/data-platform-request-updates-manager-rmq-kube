package apiModuleRuntimesResponsesGoogleAccountAccessToken

type GoogleAccountAccessTokenRes struct {
	Message GoogleAccountAccessToken `json:"message,omitempty"`
}

//type GoogleAccountAccessTokenGlobal struct {
//	URL *[]GoogleAccountAccessToken `json:"URL,omitempty"`
//}

type GoogleAccountAccessToken struct {
	AccessToken string `json:"AccessToken"`
}
