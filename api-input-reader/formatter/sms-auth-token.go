package formatter

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"data-platform-request-updates-manager-rmq-kube/services"
	"encoding/json"
	"golang.org/x/xerrors"
)

type SMSAuthTokenSDC struct {
	ConnectionKey    string       `json:"connection_key"`
	Result           bool         `json:"result"`
	RedisKey         string       `json:"redis_key"`
	Filepath         string       `json:"filepath"`
	APIStatusCode    int          `json:"api_status_code"`
	RuntimeSessionID string       `json:"runtime_session_id"`
	BusinessPartner  *int         `json:"business_partner"`
	ServiceLabel     string       `json:"service_label"`
	APIType          string       `json:"api_type"`
	SMSAuthToken     SMSAuthToken `json:"SMSAuthToken"`
	APISchema        string       `json:"api_schema"`
	Accepter         []string     `json:"accepter"`
	Deleted          bool         `json:"deleted"`
}

type SMSAuthToken struct {
	UserID            string `json:"UserID"`
	MobilePhoneNumber string `json:"MobilePhoneNumber"`
}

func ConvertToSMSAuthToken(
	input types.SMSAuthTokenSDC,
) (*SMSAuthTokenSDC, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToSMSAuthToken data marshal error :%#v", err.Error())
	}

	convertData := SMSAuthTokenSDC{}
	err = json.Unmarshal(raw, &convertData)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToSMSAuthToken input data marshal error :%#v", err.Error())
	}

	internationalPhoneNumber, err := services.ConvertToInternationalPhoneNumber(convertData.SMSAuthToken.MobilePhoneNumber)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToSMSAuthToken international phone number convert error :%#v", err.Error())
	}
	convertData.SMSAuthToken.MobilePhoneNumber = internationalPhoneNumber

	return &convertData, nil
}
