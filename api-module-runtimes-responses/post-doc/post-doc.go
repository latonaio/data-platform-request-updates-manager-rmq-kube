package apiModuleRuntimesResponsesPostDoc

type PostDocRes struct {
	Message PostDoc `json:"message,omitempty"`
}

type PostDoc struct {
	HeaderDoc  HeaderDoc    `json:"HeaderDoc"`
	HeaderDocs *[]HeaderDoc `json:"HeaderDocs"`
}

type HeaderDoc struct {
	Post                     int     `json:"Post"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}
