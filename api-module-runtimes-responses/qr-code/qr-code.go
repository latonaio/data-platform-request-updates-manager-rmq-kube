package apiModuleRuntimesResponsesQRCode

type QRCodeRes struct {
	Message QRCodeGlobal `json:"message,omitempty"`
}

type QRCodeGlobal struct {
	QRCode QRCode `json:"QRCode"`
}

type QRCode struct {
	APIServiceName            string  `json:"APIServiceName"`
	ServiceLabel              string  `json:"ServiceLabel"`
	APIType                   string  `json:"APIType"`
	URLFormatBeforeConversion string  `json:"URLFormatBeforeConversion"`
	URLFormatForRead          string  `json:"URLFormatForRead"`
	MountPath                 string  `json:"MountPath"`
	BusinessPartner           *int   `json:"BusinessPartner"`
	Event                     *int   `json:"Event"`
	Article                   *int   `json:"Article"`
	Shop                      *int   `json:"Shop"`
	Site                      *int   `json:"Site"`
	Station                   *int   `json:"Station"`
	BusStop                   *int   `json:"BusStop"`
	Participation             *int   `json:"Participation"`
	DocType                   string  `json:"DocType"`
	DocVersionID              int     `json:"DocVersionID"`
	DocID                     *string `json:"DocID"`
	FileExtension             string  `json:"FileExtension"`
	FileName                  string  `json:"FileName"`
	DocIssuerBusinessPartner  int     `json:"DocIssuerBusinessPartner"`
}
