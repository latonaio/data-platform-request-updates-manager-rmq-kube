package formatter

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"encoding/json"
	"golang.org/x/xerrors"
)

type PointTransactionSDC struct {
	ConnectionKey    string                 `json:"connection_key"`
	Result           bool                   `json:"result"`
	RedisKey         string                 `json:"redis_key"`
	Filepath         string                 `json:"filepath"`
	APIStatusCode    int                    `json:"api_status_code"`
	RuntimeSessionID string                 `json:"runtime_session_id"`
	BusinessPartner  *int                   `json:"business_partner"`
	ServiceLabel     string                 `json:"service_label"`
	APIType          string                 `json:"api_type"`
	Header           PointTransactionHeader `json:"PointTransaction"`
	APISchema        string                 `json:"api_schema"`
	Accepter         []string               `json:"accepter"`
	Deleted          bool                   `json:"deleted"`
}

type PointTransactionHeader struct {
	PointTransaction                      *int    `json:"PointTransaction"`
	PointTransactionType                  string  `json:"PointTransactionType"`
	PointTransactionDate                  string  `json:"PointTransactionDate"`
	PointTransactionTime                  string  `json:"PointTransactionTime"`
	Sender                                int     `json:"Sender"`
	Receiver                              int     `json:"Receiver"`
	PointSymbol                           string  `json:"PointSymbol"`
	PlusMinus                             string  `json:"PlusMinus"`
	PointTransactionAmount                float32 `json:"PointTransactionAmount"`
	PointTransactionObjectType            string  `json:"PointTransactionObjectType"`
	PointTransactionObject                int     `json:"PointTransactionObject"`
	SenderPointBalanceBeforeTransaction   float32 `json:"SenderPointBalanceBeforeTransaction"`
	SenderPointBalanceAfterTransaction    float32 `json:"SenderPointBalanceAfterTransaction"`
	ReceiverPointBalanceBeforeTransaction float32 `json:"ReceiverPointBalanceBeforeTransaction"`
	ReceiverPointBalanceAfterTransaction  float32 `json:"ReceiverPointBalanceAfterTransaction"`
	Attendance							  *int	  `json:"Attendance"`
	Participation						  *int	  `json:"Participation"`
	CreationDate                          string  `json:"CreationDate"`
	CreationTime                          string  `json:"CreationTime"`
	IsCancelled                           *bool   `json:"IsCancelled"`
	BusinessPartner                       int     `json:"BusinessPartner"`
	Event                                 *int    `json:"Event"`
	EventOwner                            *int    `json:"EventOwner"`
	PointConditionRateValue               float32 `json:"PointConditionRateValue"`
}

func ConvertToPointTransactionCreatesHeaderFromEvent(
	input types.PointTransactionSDC,
) (*PointTransactionSDC, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToPointTransactionCreatesHeaderFromEvent data marshal error :%#v", err.Error())
	}

	convertData := PointTransactionSDC{}
	err = json.Unmarshal(raw, &convertData)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToPointTransactionCreatesHeaderFromEvent input data marshal error :%#v", err.Error())
	}

	convertData.Header.Sender = *convertData.Header.EventOwner
	convertData.Header.PointTransactionType = "100"
	convertData.Header.Receiver = convertData.Header.BusinessPartner
	convertData.Header.PlusMinus = "+"
	convertData.Header.PointTransactionObject = *convertData.Header.Event
	convertData.Header.PointTransactionObjectType = "Event"
	convertData.Header.PointTransactionAmount = convertData.Header.PointConditionRateValue

	return &convertData, nil
}
