package apiModuleRuntimesResponsesSiteDoc

type SiteDocRes struct {
	Message SiteDoc `json:"message,omitempty"`
}

type SiteDoc struct {
	HeaderDoc HeaderDoc `json:"HeaderDoc"`
}

type HeaderDoc struct {
	Site                     int     `json:"Site"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}
