package apiModuleRuntimesResponsesGoogleAccountAuthKey

type GoogleAccountAuthKeyRes struct {
	Message GoogleAccountAuthKey `json:"message,omitempty"`
}

//type GoogleAccountAuthKeyGlobal struct {
//	URL *[]GoogleAccountAuthKey `json:"URL,omitempty"`
//}

type GoogleAccountAuthKey struct {
	URL string `json:"URL"`
}
