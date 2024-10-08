package apiOutputFormatter

type QRCodeSDC struct {
	ConnectionKey       string              `json:"connection_key"`
	RedisKey            string              `json:"redis_key"`
	Filepath            string              `json:"filepath"`
	APIStatusCode       int                 `json:"api_status_code"`
	RuntimeSessionID    string              `json:"runtime_session_id"`
	BusinessPartnerID   *int                `json:"business_partner"`
	ServiceLabel        string              `json:"service_label"`
	APIType             string              `json:"api_type"`
	BusinessPartner     *BusinessPartnerDoc `json:"BusinessPartner"`
	Event               *EventDoc           `json:"Event"`
	Article             *ArticleDoc         `json:"Article"`
	Shop                *ShopDoc            `json:"Shop"`
	Site                *SiteDoc            `json:"Site"`
	Station             *StationDoc         `json:"Station"`
	BusStop             *BusStopDoc         `json:"BusStop"`
	Participation       *ParticipationDoc   `json:"Participation"`
	APISchema           string              `json:"api_schema"`
	Accepter            []string            `json:"accepter"`
	Deleted             bool                `json:"deleted"`
	SQLUpdateResult     *bool               `json:"sql_update_result"`
	SQLUpdateError      string              `json:"sql_update_error"`
	SubfuncResult       *bool               `json:"subfunc_result"`
	SubfuncError        string              `json:"subfunc_error"`
	ExconfResult        *bool               `json:"exconf_result"`
	ExconfError         string              `json:"exconf_error"`
	APIProcessingResult *bool               `json:"api_processing_result"`
	APIProcessingError  string              `json:"api_processing_error"`
	DocData             string              `json:"doc_data"`
}

type BusinessPartnerDoc struct {
	BusinessPartner int        `json:"BusinessPartner"`
	GeneralDoc      *HeaderDoc `json:"GeneralDoc"`
}

type EventDoc struct {
	Event     int        `json:"Event"`
	HeaderDoc *HeaderDoc `json:"HeaderDoc"`
}

type ArticleDoc struct {
	Article   int        `json:"Article"`
	HeaderDoc *HeaderDoc `json:"HeaderDoc"`
}

type ShopDoc struct {
	Shop      int        `json:"Shop"`
	HeaderDoc *HeaderDoc `json:"HeaderDoc"`
}

type SiteDoc struct {
	Site      int        `json:"Site"`
	HeaderDoc *HeaderDoc `json:"HeaderDoc"`
}

type StationDoc struct {
	Station   int        `json:"Station"`
	HeaderDoc *HeaderDoc `json:"HeaderDoc"`
}

type BusStopDoc struct {
	BusStop   int        `json:"BusStop"`
	HeaderDoc *HeaderDoc `json:"HeaderDoc"`
}

type ParticipationDoc struct {
	Participation int        `json:"Participation"`
	HeaderDoc     *HeaderDoc `json:"HeaderDoc"`
}

type HeaderDoc struct {
	APIServiceName            string  `json:"APIServiceName"`
	ServiceLabel              string  `json:"ServiceLabel"`
	APIType                   string  `json:"APIType"`
	URLFormatBeforeConversion string  `json:"URLFormatBeforeConversion"`
	URLFormatForRead          string  `json:"URLFormatForRead"`
	MountPath                 string  `json:"EventMountPath"`
	BusinessPartner           *int    `json:"BusinessPartner"`
	Event                     *int    `json:"Event"`
	Article                   *int    `json:"Article"`
	Shop                      *int    `json:"Shop"`
	Site                      *int    `json:"Site"`
	Station                   *int    `json:"Station"`
	BusStop                   *int    `json:"BusStop"`
	Participation             *int    `json:"Participation"`
	DocType                   string  `json:"DocType"`
	DocVersionID              int     `json:"DocVersionID"`
	DocID                     *string `json:"DocID"`
	FileExtension             string  `json:"FileExtension"`
	FileName                  string  `json:"FileName"`
	DocIssuerBusinessPartner  int     `json:"DocIssuerBusinessPartner"`
}
