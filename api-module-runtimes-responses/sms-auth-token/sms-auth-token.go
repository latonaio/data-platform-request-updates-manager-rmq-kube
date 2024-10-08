package apiModuleRuntimesResponsesSMSAuthToken

type SMSAuthTokenRes struct {
	Message SMSAuthToken `json:"message,omitempty"`
}

type SMSAuthToken struct {
	UserID				string `json:"UserID"`
	MobilePhoneNumber	string `json:"MobilePhoneNumber"`
	AuthenticationCode	int	   `json:"AuthenticationCode"`
}
