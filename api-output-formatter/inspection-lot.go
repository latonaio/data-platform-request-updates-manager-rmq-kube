package apiOutputFormatter

type InspectionLotSDC struct {
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

type InspectionLotHeader struct {
	InspectionLot                  int     `json:"InspectionLot"`
	InspectionPlan                 int     `json:"InspectionPlan"`
	InspectionPlantBusinessPartner int     `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                string  `json:"InspectionPlant"`
	Product                        string  `json:"Product"`
	ProductSpecification           *string `json:"ProductSpecification"`
	InspectionSpecification        *string `json:"InspectionSpecification"`
	ProductionOrder                *int    `json:"ProductionOrder"`
	ProductionOrderItem            *int    `json:"ProductionOrderItem"`
	InspectionLotHeaderText        *string `json:"InspectionLotHeaderText"`
	InspectionLotInspection        []InspectionLotInspection  `json:"InspectionLotInspection"`
}

type InspectionLotInspection struct {
	InspectionLot	                            int	        `json:"InspectionLot"`
    Inspection	                                int	        `json:"Inspection"`
    InspectionType                            	string	    `json:"InspectionType"`
    InspectionTypeValueUnit	                    *string	    `json:"InspectionTypeValueUnit"`
    InspectionTypePlannedValue	                *float32	`json:"InspectionTypePlannedValue"`
    InspectionTypeCertificateType	            *string	    `json:"InspectionTypeCertificateType"`
    InspectionTypeCertificateValueInText	    *string	    `json:"InspectionTypeCertificateValueInText"`
    InspectionTypeCertificateValueInQuantity	*float32	`json:"InspectionTypeCertificateValueInQuantity"`
    InspectionLotInspectionText	                *string	    `json:"InspectionLotInspectionText"`
}
