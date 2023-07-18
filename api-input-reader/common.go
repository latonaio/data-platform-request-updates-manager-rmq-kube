package apiInputReader

type Request struct {
	BusinessPartner *int    `json:"BusinessPartner"`
	UserID          *string `json:"UserId"`
	User            *string `json:"User"`
	Language        *string `json:"Language"`
	UserType        *string `json:"UserType"`
}
