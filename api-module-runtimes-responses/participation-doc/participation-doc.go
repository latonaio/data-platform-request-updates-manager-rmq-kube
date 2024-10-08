package apiModuleRuntimesResponsesParticipationDoc

type ParticipationDocRes struct {
	Message ParticipationDoc `json:"message,omitempty"`
}

type ParticipationDoc struct {
	HeaderDoc HeaderDoc `json:"HeaderDoc"`
}

type HeaderDoc struct {
	Participation            int     `json:"Participation"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 string  `json:"FileName"`
	FilePath                 string  `json:"FilePath"`
	DocIssuerBusinessPartner int     `json:"DocIssuerBusinessPartner"`
}
