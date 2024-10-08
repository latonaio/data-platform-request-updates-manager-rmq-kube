package apiModuleRuntimesResponsesBusinessPartnerDoc

type BusinessPartnerDocRes struct {
	Message BusinessPartnerDoc `json:"message,omitempty"`
}

type BusinessPartnerDoc struct {
	GeneralDoc GeneralDoc `json:"GeneralDoc"`
}

type GeneralDoc struct {
	BusinessPartner          int     `json:"BusinessPartner"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}
