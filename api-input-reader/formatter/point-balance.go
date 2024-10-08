package formatter

import (
	"data-platform-request-updates-manager-rmq-kube/api-input-reader/types"
	"encoding/json"
	"golang.org/x/xerrors"
)

type PointBalanceSDC struct {
	ConnectionKey            string                   `json:"connection_key"`
	Result                   bool                     `json:"result"`
	RedisKey                 string                   `json:"redis_key"`
	Filepath                 string                   `json:"filepath"`
	APIStatusCode            int                      `json:"api_status_code"`
	RuntimeSessionID         string                   `json:"runtime_session_id"`
	BusinessPartner          *int                     `json:"business_partner"`
	ServiceLabel             string                   `json:"service_label"`
	APIType                  string                   `json:"api_type"`
	PointBalancePointBalance PointBalancePointBalance `json:"PointBalance"`
	APISchema                string                   `json:"api_schema"`
	Accepter                 []string                 `json:"accepter"`
	Deleted                  bool                     `json:"deleted"`
}

type PointBalancePointBalance struct {
	BusinessPartner                      int                  `json:"BusinessPartner"`
	PointSymbol                          string               `json:"PointSymbol"`
	CurrentBalance                       *float32             `json:"CurrentBalance"`
	LimitBalance                         *float32             `json:"LimitBalance"`
	CreationDate                         string               `json:"CreationDate"`
	CreationTime                         string               `json:"CreationTime"`
	LastChangeDate                       string               `json:"LastChangeDate"`
	LastChangeTime                       string               `json:"LastChangeTime"`
	SenderPointBalanceAfterTransaction   float32              `json:"SenderPointBalanceAfterTransaction"`
	ReceiverPointBalanceAfterTransaction float32              `json:"ReceiverPointBalanceAfterTransaction"`
	PointBalanceByShop                   []PointBalanceByShop `json:"ByShop"`
}

type PointBalanceByShop struct {
	BusinessPartner int      `json:"BusinessPartner"`
	PointSymbol     string   `json:"PointSymbol"`
	Shop            int      `json:"Shop"`
	CurrentBalance  float32  `json:"CurrentBalance"`
	LimitBalance    *float32 `json:"LimitBalance"`
	CreationDate    string   `json:"CreationDate"`
	CreationTime    string   `json:"CreationTime"`
	LastChangeDate  string   `json:"LastChangeDate"`
	LastChangeTime  string   `json:"LastChangeTime"`
}

func ConvertToPointBalanceUpdatesPointConsumptionSender(
	input types.PointBalanceSDC,
) (*PointBalanceSDC, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToPointBalanceUpdatesPointConsumptionSender data marshal error :%#v", err.Error())
	}

	convertData := PointBalanceSDC{}
	err = json.Unmarshal(raw, &convertData)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToPointBalanceUpdatesPointConsumptionSender input data marshal error :%#v", err.Error())
	}
	convertData.PointBalancePointBalance.CurrentBalance = &convertData.PointBalancePointBalance.SenderPointBalanceAfterTransaction

	return &convertData, nil
}

func ConvertToPointBalanceUpdatesPointConsumptionReceiver(
	input types.PointBalanceSDC,
) (*PointBalanceSDC, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToPointBalanceUpdatesPointConsumptionReceiver data marshal error :%#v", err.Error())
	}

	convertData := PointBalanceSDC{}
	err = json.Unmarshal(raw, &convertData)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToPointBalanceUpdatesPointConsumptionReceiver input data marshal error :%#v", err.Error())
	}

	receiverPointBalanceAfterTransaction := []PointBalanceByShop{
		{
			CurrentBalance: convertData.PointBalancePointBalance.ReceiverPointBalanceAfterTransaction,
			Shop:           convertData.PointBalancePointBalance.PointBalanceByShop[0].Shop,
		},
	}

	convertData.PointBalancePointBalance.PointBalanceByShop = receiverPointBalanceAfterTransaction

	return &convertData, nil
}

func ConvertToPointBalanceUpdatesPointAcquisitionReceiver(
	input types.PointBalanceSDC,
) (*PointBalanceSDC, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToPointBalanceUpdatesPointAcquisitionReceiver data marshal error :%#v", err.Error())
	}

	convertData := PointBalanceSDC{}
	err = json.Unmarshal(raw, &convertData)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToPointBalanceUpdatesPointAcquisitionReceiver input data marshal error :%#v", err.Error())
	}
	convertData.PointBalancePointBalance.CurrentBalance = &convertData.PointBalancePointBalance.ReceiverPointBalanceAfterTransaction

	return &convertData, nil
}

func ConvertToPointBalanceUpdatesPointAcquisitionSender(
	input types.PointBalanceSDC,
) (*PointBalanceSDC, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToPointBalanceUpdatesPointAcquisitionSender data marshal error :%#v", err.Error())
	}

	convertData := PointBalanceSDC{}
	err = json.Unmarshal(raw, &convertData)
	if err != nil {
		return nil, xerrors.Errorf("ConvertToPointBalanceUpdatesPointAcquisitionSender input data marshal error :%#v", err.Error())
	}
	convertData.PointBalancePointBalance.PointBalanceByShop[0].CurrentBalance = convertData.PointBalancePointBalance.SenderPointBalanceAfterTransaction

	return &convertData, nil
}
