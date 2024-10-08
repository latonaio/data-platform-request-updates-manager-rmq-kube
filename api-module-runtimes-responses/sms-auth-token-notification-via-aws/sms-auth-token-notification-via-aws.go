package apiModuleRuntimesResponsesSMSAuthTokenNotificationViaAWS

type SMSAuthTokenNotificationViaAWSRes struct {
	Message SMSAuthTokenNotificationViaAWS `json:"message,omitempty"`
}

type SMSAuthTokenNotificationViaAWS struct {
	MobilePhoneNumber  string `json:"MobilePhoneNumber"`
	AuthenticationCode int    `json:"AuthenticationCode"`
}
