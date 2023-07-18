package apiOutputFormatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header *Header `json:"Header"`
	Item   *[]Item `json:"Item"`
}

type Header struct {
	BillOfMaterial                          int     `json:"BillOfMaterial"`
	ProductStandardQuantityInBaseUnit       float32 `json:"ProductStandardQuantityInBaseUnit"`
	ProductStandardQuantityInDeliveryUnit   float32 `json:"ProductStandardQuantityInDeliveryUnit"`
	ProductStandardQuantityInProductionUnit float32 `json:"ProductStandardQuantityInProductionUnit"`
	BillOfMaterialHeaderText                *string `json:"BillOfMaterialHeaderText"`
	ValidityStartDate                       *string `json:"ValidityStartDate"`
	ValidityEndDate                         *string `json:"ValidityEndDate"`
	Item                                    []Item  `json:"Item"`
}

type Item struct {
	BillOfMaterial                                 int      `json:"BillOfMaterial"`
	BillOfMaterialItem                             int      `json:"BillOfMaterialItem"`
	ComponentProductStandardQuantityInBaseUnit     float32  `json:"ComponentProductStandardQuantityInBaseUnit"`
	ComponentProductStandardQuantityInDeliveryUnit float32  `json:"ComponentProductStandardQuantityInDeliveryUnit"`
	ComponentProductStandardScrapInPercent         *float32 `json:"ComponentProductStandardScrapInPercent"`
	IsMarkedForBackflush                           *bool    `json:"IsMarkedForBackflush"`
	BillOfMaterialItemText                         *string  `json:"BillOfMaterialItemText"`
	ValidityStartDate                              *string  `json:"ValidityStartDate"`
	ValidityEndDate                                *string  `json:"ValidityEndDate"`
}
