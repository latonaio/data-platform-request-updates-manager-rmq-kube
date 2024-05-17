package apiModuleRuntimesResponsesGoogleAccountUserInfo

type GoogleAccountUserInfoRes struct {
	Message GoogleAccountUserInfo `json:"message,omitempty"`
}

//type GoogleAccountUserInfoGlobal struct {
//	URL *[]GoogleAccountUserInfo `json:"URL,omitempty"`
//}

type GoogleAccountUserInfo struct {
	GoogleID		string `json:"GoogleID"`
	EmailAddress	string `json:"EmailAddress"`
	VerifiedEmail	bool   `json:"VerifiedEmail"`
	FullName		string `json:"FullName"`
	FirstName		string `json:"FirstName"`
	LastName		string `json:"LastName"`
	Picture			string `json:"Picture"`
	Locale			string `json:"Locale"`
}
