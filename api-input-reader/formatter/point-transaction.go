package formatter

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"encoding/json"
	"golang.org/x/xerrors"
)

type PointTransactionSDC struct {
	ConnectionKey          string                 `json:"connection_key"`
	Result                 bool                   `json:"result"`
	RedisKey               string                 `json:"redis_key"`
	Filepath               string                 `json:"filepath"`
	APIStatusCode          int                    `json:"api_status_code"`
	RuntimeSessionID       string                 `json:"runtime_session_id"`
	BusinessPartner        *int                   `json:"business_partner"`
	ServiceLabel           string                 `json:"service_label"`
	APIType                string                 `json:"api_type"`
	PointTransactionHeader PointTransactionHeader `json:"PointTransaction"`
	APISchema              string                 `json:"api_schema"`
	Accepter               []string               `json:"accepter"`
	Deleted                bool                   `json:"deleted"`
}

type PointTransactionHeader struct {
	PointTransaction                      *int    `json:"PointTransaction"`
	PointTransactionType                  string  `json:"PointTransactionType"`
	PointTransactionDate                  string  `json:"PointTransactionDate"`
	PointTransactionTime                  string  `json:"PointTransactionTime"`
	SenderObjectType                      string  `json:"SenderObjectType"`
	SenderObject                          int     `json:"SenderObject"`
	ReceiverObjectType                    string  `json:"ReceiverObjectType"`
	ReceiverObject                        int     `json:"ReceiverObject"`
	PointSymbol                           string  `json:"PointSymbol"`
	PlusMinus                             string  `json:"PlusMinus"`
	PointTransactionAmount                float32 `json:"PointTransactionAmount"`
	PointTransactionObjectType            string  `json:"PointTransactionObjectType"`
	PointTransactionObject                int     `json:"PointTransactionObject"`
	SenderPointBalanceBeforeTransaction   float32 `json:"SenderPointBalanceBeforeTransaction"`
	SenderPointBalanceAfterTransaction    float32 `json:"SenderPointBalanceAfterTransaction"`
	ReceiverPointBalanceBeforeTransaction float32 `json:"ReceiverPointBalanceBeforeTransaction"`
	ReceiverPointBalanceAfterTransaction  float32 `json:"ReceiverPointBalanceAfterTransaction"`
	Attendance                            *int    `json:"Attendance"`
	Participation                         *int    `json:"Participation"`
	ValidityStartDate                     string  `json:"ValidityStartDate"`
	ValidityEndDate                       string  `json:"ValidityEndDate"`
	CreationDate                          string  `json:"CreationDate"`
	CreationTime                          string  `json:"CreationTime"`
	IsCancelled                           *bool   `json:"IsCancelled"`
	BusinessPartner                       int     `json:"BusinessPartner"`
	Event                                 *int    `json:"Event"`
	EventOwner                            *int    `json:"EventOwner"`
	Shop                                  *int    `json:"Shop"`
	ShopOwner                             *int    `json:"ShopOwner"`
	PointConditionRateValue               float32 `json:"PointConditionRateValue"`
	SenderCurrentBalance                  float32 `json:"SenderCurrentBalance"`
	ReceiverCurrentBalance                float32 `json:"ReceiverCurrentBalance"`
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

	senderCurrentBalance := convertData.PointTransactionHeader.SenderCurrentBalance
	receiverCurrentBalance := convertData.PointTransactionHeader.ReceiverCurrentBalance
	pointConditionRateValue := convertData.PointTransactionHeader.PointConditionRateValue

	senderPointBalanceAfterTransaction := senderCurrentBalance - pointConditionRateValue
	receiverPointBalanceAfterTransaction := receiverCurrentBalance + pointConditionRateValue

	convertData.PointTransactionHeader.SenderObject = *convertData.PointTransactionHeader.Shop
	convertData.PointTransactionHeader.ReceiverObject = convertData.PointTransactionHeader.BusinessPartner
	convertData.PointTransactionHeader.PointTransactionObject = *convertData.PointTransactionHeader.Event
	convertData.PointTransactionHeader.PointTransactionAmount = convertData.PointTransactionHeader.PointConditionRateValue
	convertData.PointTransactionHeader.SenderPointBalanceBeforeTransaction = senderCurrentBalance
	convertData.PointTransactionHeader.SenderPointBalanceAfterTransaction = senderPointBalanceAfterTransaction
	convertData.PointTransactionHeader.ReceiverPointBalanceBeforeTransaction = receiverCurrentBalance
	convertData.PointTransactionHeader.ReceiverPointBalanceAfterTransaction = receiverPointBalanceAfterTransaction

	return &convertData, nil
}

func ConvertToPointTransactionCreatesHeaderFromPointConsumption(
	input types.PointTransactionSDC,
) (*PointTransactionSDC, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToPointTransactionCreatesHeaderFromPointConsumption data marshal error :%#v", err.Error())
	}

	convertData := PointTransactionSDC{}
	err = json.Unmarshal(raw, &convertData)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToPointTransactionCreatesHeaderFromPointConsumption input data marshal error :%#v", err.Error())
	}

	senderCurrentBalance := convertData.PointTransactionHeader.SenderCurrentBalance
	receiverCurrentBalance := convertData.PointTransactionHeader.ReceiverCurrentBalance
	pointConditionRateValue := convertData.PointTransactionHeader.PointConditionRateValue

	senderPointBalanceAfterTransaction := senderCurrentBalance - pointConditionRateValue
	receiverPointBalanceAfterTransaction := receiverCurrentBalance + pointConditionRateValue

	convertData.PointTransactionHeader.SenderObject = convertData.PointTransactionHeader.BusinessPartner
	convertData.PointTransactionHeader.ReceiverObject = *convertData.PointTransactionHeader.Shop
	convertData.PointTransactionHeader.PointTransactionObject = *convertData.PointTransactionHeader.Shop
	convertData.PointTransactionHeader.PointTransactionAmount = convertData.PointTransactionHeader.PointConditionRateValue
	convertData.PointTransactionHeader.SenderPointBalanceBeforeTransaction = senderCurrentBalance
	convertData.PointTransactionHeader.SenderPointBalanceAfterTransaction = senderPointBalanceAfterTransaction
	convertData.PointTransactionHeader.ReceiverPointBalanceBeforeTransaction = receiverCurrentBalance
	convertData.PointTransactionHeader.ReceiverPointBalanceAfterTransaction = receiverPointBalanceAfterTransaction

	return &convertData, nil
}
