package apiModuleRuntimesResponsesGoogleAccountInitialAuth

type GoogleAccountInitialAuthRes struct {
	Message GoogleAccountInitialAuth `json:"message,omitempty"`
}

//type GoogleAccountInitialAuthGlobal struct {
//	URL *[]GoogleAccountInitialAuth `json:"URL,omitempty"`
//}

type GoogleAccountInitialAuth struct {
	URL string `json:"URL"`
}
